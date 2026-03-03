# Contributing

Thanks for your interest in our project. Contributions are welcome. Feel free to [open an issue](https://github.com/dbsystel/cdk-sops-secrets/issues/new) with questions or reporting ideas and bugs, or [open pull requests](https://github.com/dbsystel/cdk-sops-secrets/pulls) to contribute code.

We are committed to fostering a welcoming, respectful, and harassment-free environment. Be kind!

**Security Issues:** Please report security vulnerabilities via our [Security Policy](https://github.com/dbsystel/cdk-sops-secrets/blob/main/.github/SECURITY.md) instead of opening public issues.

## Contributing Code

1. Fork the repository
2. Install [mise](https://mise.jdx.dev/) - it manages all project dependencies and tools
3. Create a feature branch (`git checkout -b feature/amazing-feature`)
4. Make your changes
5. Run `mise run build` locally to ensure everything builds successfully
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to your branch (`git push origin feature/amazing-feature`)
8. [Open a Pull Request](https://github.com/dbsystel/cdk-sops-secrets/compare)

## Testing Policy

All contributions that add or modify functionality must include appropriate tests. This ensures code quality and prevents regressions.

**Test Requirements:**

- **TypeScript/CDK changes:** Add or update unit tests using Jest
- **Go Lambda changes:** Add or update unit tests using Go's testing package
- **New features:** Include both unit tests and integration tests where applicable
- **Bug fixes:** Add tests that verify the fix and prevent regression

Tests must pass in CI before pull requests can be merged. See the testing section below for how to run tests locally.

## How to build/deploy local

### Prerequisites

This project uses [mise](https://mise.jdx.dev/) for task management and tool versioning.

**Install mise:**

On macOS using Homebrew:

```bash
brew install mise
```

For other platforms, follow the [installation guide](https://mise.jdx.dev/getting-started.html).

Tool versions (Node.js, Go, etc.) are automatically managed via `mise.toml` ([tools] section).

### Dev Container Setup (Recommended)

For the simplest development environment setup, use the included dev container configuration. This provides a pre-configured environment with all necessary tools:

1. Open the project in VS Code (or any IDE that supports dev containers)
2. When prompted, select "Reopen in Container" (or use the Command Palette: "Dev Containers: Reopen in Container")
3. The dev container will automatically set up the environment with mise managing all required tools and dependencies

This approach ensures a consistent development environment across all contributors without manual tool installation.

### Install Dependencies

```bash
mise run install
```

### Building

Build the complete project (includes lambda, TypeScript compilation, and tests):

```bash
mise run build
```

Build only the Go lambda code:

```bash
mise run lambda:build
```

### Testing

#### Integration (CDK) Tests

Use generic tasks (optionally set TEST to limit to one):

Deploy all:

```bash
mise run integ:deploy
```

Deploy single:

```bash
TEST=SECRET mise run integ:deploy
```

Assert:

```bash
mise run integ:assert
# or single
TEST=SECRET mise run integ:assert
```

Update snapshots (failed only):

```bash
mise run integ:snapshot
```

Destroy stacks:

```bash
mise run integ:destroy
# or single
TEST=SECRET mise run integ:destroy
```

Combined workflow (deploy, assert, snapshot):

```bash
mise run integ:all
```

Run TypeScript tests:

```bash
mise run test
```

Run Go lambda unit tests (default set):

```bash
mise run lambda:test
```

Run Go lambda integration tests (explicit only):

```bash
mise run lambda:test:integration
```

### Packaging

Package for JavaScript/npm (default in CI):

```bash
mise run package:js
```

Package for all targets (Java, Python, .NET, JavaScript):

```bash
mise run package:all
```

### Local Development

You can still use `npm link` if desired, but the recommended workflow is to rely on `mise` tasks directly for building and testing.

To build and watch during local development:

```bash
mise run build
mise run watch
```

If you need to link the package (optional):

```bash
npm link
npm link "cdk-sops-secrets"
```

### Other Useful Tasks

Format code:

```bash
mise run format
```

Lint code:

```bash
mise run lint
```

See all available tasks:

```bash
mise tasks
```
