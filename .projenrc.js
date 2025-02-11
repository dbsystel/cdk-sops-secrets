const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Markus Siebert',
  authorAddress: 'markus.siebert@deutschebahn.com',
  cdkVersion: '2.177.0',
  majorVersion: 2,
  stability: 'stable',
  homepage: 'https://constructs.dev/packages/cdk-sops-secrets',
  description:
    'CDK Constructs that syncs your sops secrets into AWS SecretsManager secrets.',
  keywords: [
    'mozilla/sops',
    'getsops/sops',
    'sops',
    'kms',
    'gitops',
    'secrets management',
    'secrets',
  ],
  githubOptions: {
    mergify: false,
  },
  defaultReleaseBranch: 'main',
  npmignoreEnabled: true,
  autoApproveUpgrades: true,
  autoApproveOptions: {
    allowedUsernames: [
      'markussiebert',
      'renovate-bot',
      'renovate',
      'renovate[bot]',
    ],
  },
  buildWorkflowOptions: {

    preBuildSteps: [
      {
        name: "Setup Go ${{ matrix.go-version }}",
        uses: "actions/setup-go@v5",
        with: {
          "go-version": "${{ matrix.go-version }}",
        },
      },
      {
        name: "Display Go version",
        run: "go version"
      },
      {
        name: 'Test',
        run: 'scripts/lambda-test.sh',
      },
      {
        name: 'Build',
        run: 'scripts/lambda-build.sh',
      },
      {
        name: 'Upload artifact',
        uses: 'actions/upload-artifact@v4',
      },
    ]
  },
  name: 'cdk-sops-secrets',
  repositoryUrl: 'https://github.com/dbsystel/cdk-sops-secrets.git',
  bundledDeps: ['yaml'],
  // deps: [], /* Runtime dependencies of this module. */,
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  devDeps: [
    'json-schema-to-typescript',
  ] /* Build dependencies for this module. */,
  integrationTestAutoDiscover: true,
  prettier: true,
  prettierOptions: {
    settings: {
      semi: true,
      singleQuote: true,
      trailingComma: 'all',
      tabWidth: 2,
      ignoreFile: true,
    },
  },
  eslint: true,
  eslintOptions: { prettier: true },
  // seems to overwrite the wohle repository ... maybe another repository needed?
  //publishToGo: {
  //  moduleName: 'github.com/dbsystel/cdk-sops-secrets',
  //  githubTokenSecret: 'PROJEN_GITHUB_TOKEN',
  //},
  publishToMaven: {
    javaPackage: 'de.db.systel.cdkSopsSecrets',
    mavenGroupId: 'de.db.systel',
    mavenArtifactId: 'cdk-sops-secrets',
    mavenRepositoryUrl:
      'https://maven.pkg.github.com/dbsystel/cdk-sops-secrets',
  },
  publishToNuget: {
    dotNetNamespace: 'Db.De.Systel',
    packageId: 'Db.De.Systel.CdkSopsSecrets',
    nugetServer: 'https://nuget.pkg.github.com/dbsystel/index.json',
    nugetApiKeySecret: 'PROJEN_GITHUB_TOKEN',
  },
  publishToPypi: {
    distName: 'cdk-sops-secrets',
    module: 'cdk_sops_secrets',
  },
});

project.prettier.addIgnorePattern('/test-secrets/');
project.prettier.addIgnorePattern('/test/*snapshot');
project.prettier.addIgnorePattern('API.md');
project.prettier.addIgnorePattern('package.json');
project.jest.addIgnorePattern('/lambda/');
project.gitignore.addPatterns('*.iml', '.idea', '/assets');
project.npmignore.addPatterns(
  '/lambda',
  '/dist-lambda',
  '/scripts',
  '!/assets',
  '/lambda/',
  '/img/',
  '/test-secrets/',
  '/renovate.json',
  '/.prettier*',
  '/.whitesource',
  '/.gitattributes',
);

goreleaserArtifactsNamespace = 'build-artifact-goreleaser';

project.synth();
