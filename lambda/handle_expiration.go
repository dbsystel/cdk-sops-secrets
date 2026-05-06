package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/event"
)

const defaultExpirationSuffix = "_expiration"
const defaultDaysBeforeExpiration = 14

// expirationMessage is the payload published to the SNS topic.
type expirationMessage struct {
	SecretArn            string `json:"secretArn"`
	KeyName              string `json:"keyName"`
	ExpirationDate       string `json:"expirationDate"`
	NotificationDate     string `json:"notificationDate"`
	DaysBeforeExpiration int    `json:"daysBeforeExpiration"`
}

// sanitizeScheduleName converts a key name into a valid EventBridge Scheduler
// schedule name: only [a-zA-Z0-9_-], max 64 characters.
func sanitizeScheduleName(name string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	sanitized := re.ReplaceAllString(name, "-")
	if len(sanitized) > 64 {
		sanitized = sanitized[:64]
	}
	return sanitized
}

// parseExpirationDate attempts to parse a date string in ISO 8601 / RFC 3339
// formats. Accepted formats: "2006-01-02", "2006-01-02T15:04:05Z07:00".
func parseExpirationDate(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	formats := []string{
		"2006-01-02",
		time.RFC3339,
	}
	for _, format := range formats {
		if t, err := time.Parse(format, value); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %q", value)
}

// scanExpirationKeys scans the flattened secret map for keys that end with the
// given suffix and returns a map of base key → expiration time.
func scanExpirationKeys(secretMap map[string]string, suffix string) (map[string]time.Time, error) {
	result := make(map[string]time.Time)

	keys := make([]string, 0, len(secretMap))
	for key := range secretMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := secretMap[key]
		if !strings.HasSuffix(key, suffix) {
			continue
		}
		baseKey := strings.TrimSuffix(key, suffix)
		t, err := parseExpirationDate(value)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to parse expiration date for key %s with value %q: %w",
				key,
				value,
				err,
			)
		}
		result[baseKey] = t
	}
	return result, nil
}

// handleExpirationUpsert creates or updates EventBridge Scheduler schedules for
// all expiration keys found in the decrypted secret.
func handleExpirationUpsert(
	props *event.SopsSyncResourceProperties,
	clients client.AwsClient,
	secretMap map[string]string,
	secretArn string,
) error {
	logger := slog.With("Package", "main", "Function", "handleExpirationUpsert")
	exp := props.Expiration

	suffix := defaultExpirationSuffix
	if exp.ExpirationSuffix != nil && *exp.ExpirationSuffix != "" {
		suffix = *exp.ExpirationSuffix
	}

	days := defaultDaysBeforeExpiration
	if exp.DaysBeforeExpiration != nil {
		days = *exp.DaysBeforeExpiration
	}

	expirations, err := scanExpirationKeys(secretMap, suffix)
	if err != nil {
		return err
	}
	if len(expirations) == 0 {
		logger.Info("No expiration keys found in secret", "Suffix", suffix)
		return nil
	}

	for baseKey, expiresAt := range expirations {
		notifyAt := expiresAt.UTC().AddDate(0, 0, -days)
		// Skip schedules whose notification time is already in the past.
		if notifyAt.Before(time.Now().UTC()) {
			logger.Warn("Notification time is in the past, skipping schedule",
				"baseKey", baseKey,
				"expiresAt", expiresAt.Format("2006-01-02"),
				"notifyAt", notifyAt.Format("2006-01-02"),
			)
			continue
		}

		scheduleName := sanitizeScheduleName(baseKey)
		// EventBridge Scheduler at() expression format: at(YYYY-MM-DDTHH:mm:ss)
		scheduleExpr := fmt.Sprintf("at(%s)", notifyAt.Format("2006-01-02T15:04:05"))

		msg := expirationMessage{
			SecretArn:            secretArn,
			KeyName:              baseKey,
			ExpirationDate:       expiresAt.UTC().Format("2006-01-02"),
			NotificationDate:     notifyAt.UTC().Format("2006-01-02"),
			DaysBeforeExpiration: days,
		}
		msgJSON, err := json.Marshal(msg)
		if err != nil {
			return fmt.Errorf("failed to marshal expiration message for key %s: %v", baseKey, err)
		}

		logger.Info("Upserting expiration schedule",
			"scheduleName", scheduleName,
			"group", exp.ScheduleGroupName,
			"expression", scheduleExpr,
		)

		if err := clients.SchedulerCreateOrUpdateSchedule(
			scheduleName,
			exp.ScheduleGroupName,
			scheduleExpr,
			exp.TopicArn,
			exp.SchedulerRoleArn,
			string(msgJSON),
		); err != nil {
			return fmt.Errorf("failed to upsert schedule for key %s: %v", baseKey, err)
		}
	}
	return nil
}

// handleExpirationDelete removes all expiration schedules for a secret by
// listing and deleting every schedule in the group.
func handleExpirationDelete(
	props *event.SopsSyncResourceProperties,
	clients client.AwsClient,
) error {
	logger := slog.With("Package", "main", "Function", "handleExpirationDelete")
	exp := props.Expiration

	names, err := clients.SchedulerListSchedules(exp.ScheduleGroupName)
	if err != nil {
		// If the group doesn't exist yet (e.g. stack was never fully deployed)
		// treat it as a no-op rather than a hard failure.
		logger.Warn("Could not list schedules, assuming group does not exist",
			"Group", exp.ScheduleGroupName, "Error", err)
		return nil
	}

	for _, name := range names {
		logger.Info("Deleting expiration schedule", "Name", name, "Group", exp.ScheduleGroupName)
		if err := clients.SchedulerDeleteSchedule(name, exp.ScheduleGroupName); err != nil {
			// Log and continue so we attempt to delete all schedules even if one fails.
			logger.Error("Failed to delete schedule", "Name", name, "Error", err)
		}
	}
	return nil
}
