# Changelog

All notable changes to this project will be documented in this file. See [commit-and-tag-version](https://github.com/absolute-version/commit-and-tag-version) for commit guidelines.

## [2.7.2](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.7.1...v2.7.2) (2026-05-19)


### Documentation

* add OpenSSF Scorecard badge to README ([f2651ef](https://github.com/dbsystel/cdk-sops-secrets/commit/f2651ef72f3a6976be778f99aea58b27deafd170))


### Miscellaneous Chores

* add build provenance attestation to release workflow ([ea5cb78](https://github.com/dbsystel/cdk-sops-secrets/commit/ea5cb7841d8cb0c47118cf23e3d2a648b281aa20))
* add OpenSSF Scorecard badge to README ([98fd265](https://github.com/dbsystel/cdk-sops-secrets/commit/98fd265e17c5bd550a4befa27e462bc3dc2d25f1))
* attach SLSA provenance to GitHub Release assets ([e7f527c](https://github.com/dbsystel/cdk-sops-secrets/commit/e7f527c1c047f39c692fa5598fa79453367fe6ed))
* **ci:** switch release process to release-please with GitHub App ([aae40e2](https://github.com/dbsystel/cdk-sops-secrets/commit/aae40e2a40565545f210db0777fef3bf57016e88))
* **deps:** update actions/attest-build-provenance digest to e8998f9 ([#1336](https://github.com/dbsystel/cdk-sops-secrets/issues/1336)) ([cc119d4](https://github.com/dbsystel/cdk-sops-secrets/commit/cc119d4019fce4a81afec77dff273567fa68e783))
* **deps:** update github-actions ([#1328](https://github.com/dbsystel/cdk-sops-secrets/issues/1328)) ([aed7c7c](https://github.com/dbsystel/cdk-sops-secrets/commit/aed7c7cb39a10995745675a48f0076dd0cbabf8a))
* **deps:** update npm-deps ([#1320](https://github.com/dbsystel/cdk-sops-secrets/issues/1320)) ([e3f518a](https://github.com/dbsystel/cdk-sops-secrets/commit/e3f518af8677889e1be623070d75a9e54cb9b5b9))
* **deps:** update npm-deps ([#1321](https://github.com/dbsystel/cdk-sops-secrets/issues/1321)) ([d2f0b51](https://github.com/dbsystel/cdk-sops-secrets/commit/d2f0b51dc9e1aa2ea6555c74476941fb5f805f66))
* **deps:** upgrade go depdendencies ([6a64b3b](https://github.com/dbsystel/cdk-sops-secrets/commit/6a64b3b07e79d651c53467423f23170334138895))
* **deps:** upgrade go dependencies ([1828b7c](https://github.com/dbsystel/cdk-sops-secrets/commit/1828b7c57d751f97264071ea63a8e23458ce8be7))
* enable rebase-behind-base-branch in renovate config ([ef36f79](https://github.com/dbsystel/cdk-sops-secrets/commit/ef36f797f97ff670f0edc3bc9b426b46f15c1746))
* enable rebase-behind-base-branch in renovate config ([911af85](https://github.com/dbsystel/cdk-sops-secrets/commit/911af85a5ee3ad315ec4a3cd9c4d7fd6ffa7db6e))
* make chore commits visible in release-please ([8ff8aa7](https://github.com/dbsystel/cdk-sops-secrets/commit/8ff8aa7ccb5f31f955ede86fee4b2009b2854011))
* make chore commits visible in release-please ([1f59e23](https://github.com/dbsystel/cdk-sops-secrets/commit/1f59e23b6029a941241afd50e7a4f53fc7d95f5d))
* pin golang to 1.26.2 in mise.toml ([0f5bff4](https://github.com/dbsystel/cdk-sops-secrets/commit/0f5bff49f1539cebb4bb193bb67cc7a1cdddf96c))
* pin golang to 1.26.3 in mise.toml ([c1bba89](https://github.com/dbsystel/cdk-sops-secrets/commit/c1bba89f61f1d283f4cfa1b317ddefc961333197))
* scope release-please permissions to job level ([b43ce11](https://github.com/dbsystel/cdk-sops-secrets/commit/b43ce1198b4bae4a53630329a0e1544af14fd324))
* scope release-please permissions to job level ([f5df6a7](https://github.com/dbsystel/cdk-sops-secrets/commit/f5df6a715af69ca46ceb7dd7698bcdec53978617))
* set minimun release age ([6292af2](https://github.com/dbsystel/cdk-sops-secrets/commit/6292af28ea022a03362f1dcfbdbf99e0a1f20fd0))
* set minimun release age ([8d7b716](https://github.com/dbsystel/cdk-sops-secrets/commit/8d7b7161f050dd6801459ea4f6c2cff8edb156f7))
* sync min-release-age to 1 day (renovate + npmrc) ([73dc285](https://github.com/dbsystel/cdk-sops-secrets/commit/73dc285cdf8841790111c38fbf8d6da472d08578))
* sync min-release-age to 1 days (renovate + npmrc) ([4e9976b](https://github.com/dbsystel/cdk-sops-secrets/commit/4e9976b17709b9126589e2b4297a9260d954d230))
* update node to 24.14.0, pin npm to 11.12.1 ([1d071da](https://github.com/dbsystel/cdk-sops-secrets/commit/1d071dadf3b459bac3ecf412df9a10b5ab0501a5))
* update node to 24.14.0, pin npm to 11.14.1 ([22b7407](https://github.com/dbsystel/cdk-sops-secrets/commit/22b7407bfe3f83a1114dcaf9cd042a832f2001ea))

## [2.7.1](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.7.0...v2.7.1) (2026-05-08)


### Bug Fixes

* assertion name ([ca35532](https://github.com/dbsystel/cdk-sops-secrets/commit/ca35532c1279fb4704e65c7745e7afeb4899e3a6))
* dont let the schedule name start with a slash ([1778a28](https://github.com/dbsystel/cdk-sops-secrets/commit/1778a28151c35a1089dd7fee1281cb9d3e22d48f))

## [2.7.0](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.7...v2.7.0) (2026-05-07)

### Features

* add secret expiration tracking ([eb6c9b7](https://github.com/dbsystel/cdk-sops-secrets/commit/eb6c9b76875b0bfbe3e0b95d92f41a6b697fd0c0))

## [2.6.7](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.6...v2.6.7) (2026-05-05)


### Bug Fixes

* introduce SOPS key as ssm parameter ([#1291](https://github.com/dbsystel/cdk-sops-secrets/issues/1291)) ([d18e6e2](https://github.com/dbsystel/cdk-sops-secrets/commit/d18e6e222c5b1fedece04be42d334fa420da4510))

## [2.6.6](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.5...v2.6.6) (2026-04-16)


### Bug Fixes

* bump Go indirect deps for CVE fixes and enable Renovate indirect updates ([c058243](https://github.com/dbsystel/cdk-sops-secrets/commit/c0582437b7f3d6f81eec0f2fcc8c9ca465aca646)), closes [#1305](https://github.com/dbsystel/cdk-sops-secrets/issues/1305)
* **renovate:** disable major updates for indirect Go deps ([c2ab489](https://github.com/dbsystel/cdk-sops-secrets/commit/c2ab4893d0b0d25afef55a7979eb7aa8bf0181ba))

## [2.6.5](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.4...v2.6.5) (2026-04-09)


### Bug Fixes

* move workflow permissions to job level ([1ea64e0](https://github.com/dbsystel/cdk-sops-secrets/commit/1ea64e00827ff1887f432cd89809c99fde9a66d3))
* prevent script injection in tag-on-merge workflow ([85fb85d](https://github.com/dbsystel/cdk-sops-secrets/commit/85fb85de35b99f54e2fc8a95d3d4e6a178574345))
* reduce pull-request-lint permissions to read-only ([8d53c56](https://github.com/dbsystel/cdk-sops-secrets/commit/8d53c5695257fdf7a3776ca8e40c4a44f8ecdd4f))
* remove dangerous workflow patterns ([1f20363](https://github.com/dbsystel/cdk-sops-secrets/commit/1f20363020d66e0a2d5e534ffcbf960559a682d3))
* remove unnecessary actions:write permission from create-release ([fd3401a](https://github.com/dbsystel/cdk-sops-secrets/commit/fd3401ae0ed1c36d925044fbabae3f69d94e47cf))

## [2.6.4](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.3...v2.6.4) (2026-03-01)

## [2.6.3](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.2...v2.6.3) (2026-02-20)

## [2.6.2](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.1...v2.6.2) (2026-02-16)

## [2.6.1](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.6.0...v2.6.1) (2026-02-07)

## [2.6.0](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.5.0...v2.6.0) (2026-02-07)


### Features

* make create-release workflow rerunnable ([#1225](https://github.com/dbsystel/cdk-sops-secrets/issues/1225)) ([cdf59b4](https://github.com/dbsystel/cdk-sops-secrets/commit/cdf59b43013520aa7b8c6cd45ba4f5a564fbbf37)), closes [#1215](https://github.com/dbsystel/cdk-sops-secrets/issues/1215)

## [2.5.0](https://github.com/dbsystel/cdk-sops-secrets/compare/v2.4.7...v2.5.0) (2026-01-12)


### Features

* Lambda runtime upgrade to AL2023 ([#1223](https://github.com/dbsystel/cdk-sops-secrets/issues/1223)) ([9d61d95](https://github.com/dbsystel/cdk-sops-secrets/commit/9d61d95554078578951419fda5296dbd00e9c76a)), closes [#1215](https://github.com/dbsystel/cdk-sops-secrets/issues/1215)
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
