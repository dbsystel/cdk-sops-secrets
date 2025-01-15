# Contributing

Thanks for your interest in our project. Contributions are welcome. Feel free to [open an issue](issues) with questions or reporting ideas and bugs, or [open pull requests](pulls) to contribute code.

We are committed to fostering a welcoming, respectful, and harassment-free environment. Be kind!

## How to buidl/deploy local

Install all necessary tools with `yarn install` and others manually like `go`

Build the go Lambda code:
```
./scripts/build.sh
```
Build the package (for CDK development only the first `js` build has to complete):
```
yarn projen build
```
Link the package:
```
yarn link
```
Switch to the path/project where you would like to use cdk-sops-secrets. \
Link the package to your local build source:
```
yarn link "cdk-sops-secrets"
```