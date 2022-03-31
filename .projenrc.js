const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Markus Siebert',
  authorAddress: 'dev@markussiebert.com',
  cdkVersion: '1.0.0',
  cdkVersionPinning: false,
  defaultReleaseBranch: 'main',
  npmignoreEnabled: true,
  name: 'cdk-sops-secrets',
  repositoryUrl: 'https://github.com/markussiebert/cdk-sops-secrets.git',
  peerDeps: [
    '@aws-cdk/aws-secretsmanager@^1.0.0',
    '@aws-cdk/aws-iam@^1.0.0',
    '@aws-cdk/aws-lambda@^1.0.0',
    '@aws-cdk/aws-logs@^1.0.0',
    '@aws-cdk/custom-resources@^1.0.0',
    '@aws-cdk/aws-s3-assets@^1.0.0',
    '@aws-cdk/aws-kms@^1.0.0',
  ],
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  devDeps: [
    '@aws-cdk/aws-secretsmanager@^1.0.0',
    '@aws-cdk/aws-iam@^1.0.0',
    '@aws-cdk/aws-lambda@^1.0.0',
    '@aws-cdk/aws-logs@^1.0.0',
    '@aws-cdk/custom-resources@^1.0.0',
    '@aws-cdk/aws-s3-assets@^1.0.0',
    '@aws-cdk/aws-kms@^1.0.0',
  ],
  integrationTestAutoDiscover: true,
  // packageName: undefined,  /* The "name" in package.json. */
});

project.gitignore.addPatterns('/assets');
project.npmignore.addPatterns('/lambda', '/dist-lambda', '!/assets');

goreleaserArtifactsNamespace = 'build-artifact-goreleaser';

additionalActions = [
  {
    name: 'Download goreleaser artifacts',
    uses: 'actions/download-artifact@v2',
    with: {
      name: goreleaserArtifactsNamespace,
      path: 'assets',
    },
  },
  {
    name: 'Debug goreleaser artifacts',
    run: 'shasum assets/cdk-sops-lambda.zip && ls -la assets/cdk-sops-lambda.zip',
  },
];


project.buildWorkflow.preBuildSteps.unshift( ...additionalActions);

console.log(project.github.workflows.map((wr) => wr.name));

const fixme = project.github.workflows.filter((wf) =>
  ['build','release'].includes(wf.name),
);
fixme.forEach((wf) => {
  
  Object.keys(wf.jobs).forEach((key) => {
    if (key !== 'build') {
      wf.jobs[key].steps.splice(1, 0, ...additionalActions);
    }
    if (['build','release'].includes(key)) {
      wf.jobs[key] = { ...wf.jobs[key], container: { image: 'jsii/superchain:1-buster-slim-node16'} };
    }
    wf.jobs[key] = { ...wf.jobs[key], needs: 'goreleaser' };
  });
  wf.addJob('goreleaser', {
    name: 'goreleaser',
    runsOn: 'ubuntu-latest',
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
        name: 'Set up Go',
        uses: 'actions/setup-go@v2',
        with: {
          'go-version': 1.18,
        },
      },
      {
        name: 'Run GoReleaser',
        uses: 'goreleaser/goreleaser-action@v2',
        with: {
          distribution: 'goreleaser',
          version: 'latest',
          args: 'release --rm-dist',
        },
        env: {
          GITHUB_TOKEN: '${{ secrets.GITHUB_TOKEN }}',
        },
      },
      {
        name: 'Upload artifact',
        uses: 'actions/upload-artifact@v2.1.1',
        with: {
          name: goreleaserArtifactsNamespace,
          path: 'assets/cdk-sops-secrets.zip',
        },
      },
    ],
  });
});

project.synth();
