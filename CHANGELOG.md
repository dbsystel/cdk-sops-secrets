# Changelog

All notable changes to this project will be documented in this file. See [commit-and-tag-version](https://github.com/absolute-version/commit-and-tag-version) for commit guidelines.

## [2.5.0](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.4.7...v2.5.0) (2026-01-12)


### Features

* **renovate:** enable github-actions manager for workflow updates ([68e9211](https://github.com/dbsystel/cdk-sops-secrets/commit/68e9211d3f65af4494b585c4f89d73c74017148a))


### Bug Fixes

* correct renovate.json configuration - replace invalid gitHubActions with pinDigests ([#1198](https://github.com/dbsystel/cdk-sops-secrets/issues/1198)) ([538080b](https://github.com/dbsystel/cdk-sops-secrets/commit/538080bc0a28de55ea94effcebf02f20b1239b07))
* **deps:** update go-deps ([#1184](https://github.com/dbsystel/cdk-sops-secrets/issues/1184)) ([00bbb76](https://github.com/dbsystel/cdk-sops-secrets/commit/00bbb76ebd8ff1577b93a29574e5b0d4adffb0b3))
* remove premature GitHub release creation from create-release workflow ([#1220](https://github.com/dbsystel/cdk-sops-secrets/issues/1220)) ([cd9d6c6](https://github.com/dbsystel/cdk-sops-secrets/commit/cd9d6c6c24022942497ada1109b5f0bddb1efa74))
* **renovate:** migrate to config:recommended and move allowedPostUpgradeCommands to workflow ([ec10e40](https://github.com/dbsystel/cdk-sops-secrets/commit/ec10e40a9a4624c4c83297ab46ad7a9031ceca94))
* **renovate:** scope postUpgradeTasks to npm manager only ([9ee43b9](https://github.com/dbsystel/cdk-sops-secrets/commit/9ee43b9fddd9679dd9cb806afb6a3ba8b719966d))
* **renovate:** unknown option error ([4e5c9ae](https://github.com/dbsystel/cdk-sops-secrets/commit/4e5c9ae3ff8c4441578764dd6105e6957fbf35f0))
* **renovate:** use renovate without docker ([697c90e](https://github.com/dbsystel/cdk-sops-secrets/commit/697c90ea6df0da720e0b7a69222dc3bd23a9b486))
* **test:** fix breaking changes introduces by jest upgrade ([#1193](https://github.com/dbsystel/cdk-sops-secrets/issues/1193)) ([e168ed5](https://github.com/dbsystel/cdk-sops-secrets/commit/e168ed5705f0939ba14dcf01ea02dac841fc6711))
* update release workflow to respect branch protection ([#1218](https://github.com/dbsystel/cdk-sops-secrets/issues/1218)) ([1b956ad](https://github.com/dbsystel/cdk-sops-secrets/commit/1b956ad588df03d79f76b8ca8bde6536baa1c1bc))

## [2.4.7](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.4.6...v2.4.7) (2025-11-26)

## [2.4.6](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.4.5...v2.4.6) (2025-11-26)


### Bug Fixes

* **package-lock.json:** update packages ([f93b485](https://github.com/dbsystel/cdk-sops-secrets/commit/f93b485123f4ae6814a6ac6aadb0e3c4a4cc1957))
* **SopsStringParameter:** add parameterRef to match IStringParameter ([ea04275](https://github.com/dbsystel/cdk-sops-secrets/commit/ea04275a6ef082b7548fc484a213e85641d6c8e5))
