const { awscdk } = require('projen');
const { NodePackageManager } = require('projen/lib/javascript');

const actions_SetupGo = [
  {
    name: 'Setup Go 1.24.5',
    uses: 'actions/setup-go@v5',
    with: {
      'go-version': '1.24.5',
      'cache-dependency-path': 'lambda/go.sum',
    },
  },
  {
    name: 'Display Go version',
    run: 'go version',
  },
];

const actions_UpgradeGoDeps = [
  {
    name: 'Upgrade Go dependencies',
    run: 'cd lambda && go get -u && go mod tidy && cd ..',
  },
];

const actions_TestBuild = [
  {
    name: 'Test lambda code',
    run: 'scripts/lambda-test.sh',
  },
  {
    name: 'Build lambda code and create zip',
    run: 'scripts/build.sh',
  },
];

const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Markus Siebert',
  authorAddress: 'markus.siebert@deutschebahn.com',
  cdkVersion: '2.191.0',
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
    allowedUsernames: ['markussiebert'],
  },
  npmTrustedPublishing: true,
  packageManager: NodePackageManager.NPM,
  buildWorkflowOptions: {
    mutableBuild: true,
    preBuildSteps: [...actions_SetupGo, ...actions_TestBuild],
  },
  releaseWorkflowSetupSteps: [...actions_SetupGo, ...actions_TestBuild],
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
    javaPackage: 'de.db.systel.cdksopssecrets',
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
project.gitignore.addPatterns('*.iml', '.idea', '/assets', '**/.DS_Store');
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

// Find UpgradeJob
const upgradeWF = project.github.workflows.find(
  (wf) => wf.name == 'upgrade-main',
);
upgradeWF.events.schedule[0] = {
  cron: '0 6 * * 0', // every Sunday at 6:00
};

const upgradeJob = upgradeWF.getJob('upgrade').steps;

// Find the index of the upgrade step (npm packages)
const upgradeIndex = upgradeJob.findIndex(
  (step) => step.name === 'Upgrade dependencies',
);
upgradeJob.splice(
  upgradeIndex + 1, // After upgrading the npm deps
  0,
  ...actions_SetupGo,
  ...actions_UpgradeGoDeps,
  ...actions_TestBuild,
  {
    name: 'Create new Snapshots',
    run: 'npx projen "integ:snapshot-all"',
  },
);

const prJob = upgradeWF.getJob('pr').steps;

prJob.push({
  name: 'Enable Pull Request Automerge',
  run: 'gh pr merge ${{ steps.create-pr.outputs.pull-request-number }} --merge --auto --delete-branch',
  env: {
    GH_TOKEN: '${{ secrets.PROJEN_GITHUB_TOKEN }}',
  },
});

project.synth();
