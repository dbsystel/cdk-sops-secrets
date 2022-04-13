const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Markus Siebert',
  authorAddress: 'dev@markussiebert.com',
  cdkVersion: '2.1.0',
  description:
    'CDK Constructs that syncs your sops secrets into AWS SecretsManager secrets.',
  keywords: [
    'mozilla/sops',
    'sops',
    'kms',
    'gitops',
    'secrets management',
    'secrets',
  ],
  defaultReleaseBranch: 'main',
  npmignoreEnabled: true,
  autoApproveUpgrades: true,
  autoApproveOptions: {
    allowedUsernames: ['markussiebert'],
  },
  name: 'cdk-sops-secrets',
  repositoryUrl: 'https://github.com/markussiebert/cdk-sops-secrets.git',
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
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
});
project.prettier.addIgnorePattern('/test-secrets/');
project.prettier.addIgnorePattern('/test/*snapshot');
project.prettier.addIgnorePattern('API.md');
project.prettier.addIgnorePattern('package.json');
project.jest.addIgnorePattern('/lambda/');
project.gitignore.addPatterns('/assets');
project.npmignore.addPatterns(
  '/lambda',
  '/dist-lambda',
  '/scripts',
  '!/assets',
);

goreleaserArtifactsNamespace = 'build-artifact-goreleaser';

additionalActions = [
  {
    name: 'Download zipper artifacts',
    uses: 'actions/download-artifact@v2',
    with: {
      name: 'zipper',
      path: 'assets',
    },
  },
];



project.buildWorkflow.preBuildSteps.unshift(...additionalActions);

console.log(project.github.workflows.map((wr) => wr.name));

const fixme = project.github.workflows.filter((wf) =>
  ['build', 'release'].includes(wf.name),
);
fixme.forEach((wf) => {
  Object.keys(wf.jobs).forEach((key) => {
    if (key !== 'build') {
      wf.jobs[key].steps.splice(1, 0, ...additionalActions);
    }
    if (['build', 'release'].includes(key)) {
      wf.jobs[key] = {
        ...wf.jobs[key],
        container: { image: 'jsii/superchain:1-buster-slim-node16' },
      };
    }
    wf.jobs[key] = {
      ...wf.jobs[key],
      needs: [...(wf.jobs[key].needs || []), 'zipper'],
    };
  });
  wf.addJob('gobuild', {
    name: 'gobuild',
    runsOn: 'ubuntu-latest',
    container: {
      image: 'golang:1.18.0-buster',
    },
    on: {
      pull_request: null,
      push: null,
    },
    permissions: {
      contents: 'write',
    },
    steps: [
      {
        name: 'Checkout',
        uses: 'actions/checkout@v2',
        with: {
          'fetch-depth': 0,
        },
      },
      {
        name: 'env',
        run: 'env',
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
        uses: 'actions/upload-artifact@v2.1.1',
        with: {
          name: 'gobuild',
          path: 'lambda/cdk-sops-secrets',
        },
      },
    ],
  });
  wf.addJob('zipper', {
    name: 'zipper',
    needs: 'gobuild',
    runsOn: 'ubuntu-latest',
    container: {
      image: 'alpine:latest',
    },
    on: {
      pull_request: null,
      push: null,
    },
    permissions: {
      contents: 'write',
    },
    steps: [
      {
        name: 'Prepare',
        run: 'apk add zip git@2.34.1',
      },
      {
        name: 'Checkout',
        uses: 'actions/checkout@v2',
        with: {
          'fetch-depth': 1,
        },
      },
      {
        name: 'Download gobuild artifacts',
        uses: 'actions/download-artifact@v2',
        with: {
          name: 'gobuild',
          path: 'lambda',
        },
      },
      {
        name: 'Zip',
        run: 'scripts/lambda-zip.sh',
      },
      {
        name: 'Upload artifact',
        uses: 'actions/upload-artifact@v2.1.1',
        with: {
          name: 'zipper',
          path: 'assets/cdk-sops-lambda.zip',
        },
      },
    ],
  });
});

project.synth();
