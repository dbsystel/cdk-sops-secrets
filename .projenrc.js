const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Markus Siebert',
  authorAddress: 'dev@markussiebert.com',
  cdkVersion: '1.0.0',
  cdkVersionPinning: false,
  keywords: [
    'mozilla/sops',
    'sops',
    'kms',
    'gitops',
    'secrets management',
    'secrets',
  ],
  cdkDependenciesAsDeps: false,
  defaultReleaseBranch: 'main',
  npmignoreEnabled: true,
  name: 'cdk-sops-secrets',
  repositoryUrl: 'https://github.com/markussiebert/cdk-sops-secrets.git',
  peerDeps: [
    '@aws-cdk/aws-secretsmanager@^1.0.0',
    '@aws-cdk/aws-iam@^1.0.0',
    '@aws-cdk/aws-lambda@^1.0.0',
    '@aws-cdk/aws-logs@^1.0.0',
    '@aws-cdk/aws-s3-assets@^1.0.0',
    '@aws-cdk/aws-kms@^1.0.0',
  ],
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  devDeps: [
    '@aws-cdk/aws-secretsmanager',
    '@aws-cdk/aws-iam',
    '@aws-cdk/aws-lambda',
    '@aws-cdk/aws-logs',
    '@aws-cdk/aws-s3-assets',
    '@aws-cdk/aws-kms',
    '@aws-cdk/assertions',
  ],
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
project.npmignore.addPatterns('/lambda', '/dist-lambda', '!/assets');

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
        name: 'Fetch all tags',
        run: 'git fetch --force --tags',
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
        run: 'apk add zip git',
      },
      {
        name: 'Checkout',
        uses: 'actions/checkout@v2',
        with: {
          'fetch-depth': 0,
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
