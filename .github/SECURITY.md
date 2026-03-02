# Security Policy

## Supported Versions

We release patches for the latest major.minor version. See our [releases page](https://github.com/dbsystel/cdk-sops-secrets/releases) for the current version.

## Reporting a Vulnerability

Please report security issues via [GitHub Security Advisories](https://github.com/dbsystel/cdk-sops-secrets/security/advisories/new) ("Report a vulnerability" in the repository [Security tab](https://github.com/dbsystel/cdk-sops-secrets/security)) or e-mail the [CODEOWNERS](https://github.com/dbsystel/cdk-sops-secrets/blob/main/.github/CODEOWNERS) directly.

**DO NOT open a public issue for security vulnerabilities.**

We will acknowledge receipt of your report and work with you to understand and address the issue.

## Dependency Management

We use [Renovate](https://github.com/renovatebot/renovate) to automatically keep dependencies up to date. Routine updates are merged only after a 3-day delay ("cool-down") to reduce supply-chain risk from freshly compromised releases. High or critical severity vulnerabilities may be upgraded immediately; feel free to [open an issue](https://github.com/dbsystel/cdk-sops-secrets/issues/new) if urgent remediation is needed or if an automatic PR has not appeared.

## Security Measures

- **Branch Protection**: Required reviews and status checks on main branch
- **Dependency Scanning**: Automated via Renovate
- **Code Review**: All changes require maintainer approval
- **Workflow Security**: Minimal permissions and dangerous pattern prevention

## Public Disclosure

We kindly request you avoid public disclosure until a fix is available. We will coordinate a [CVE](https://www.cve.org/) if appropriate.
