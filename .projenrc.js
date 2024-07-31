const { awscdk } = require('projen');
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Markus Siebert',
  authorAddress: 'markus.siebert@deutschebahn.com',
  cdkVersion: '2.144.0',
  stability: 'stable',
  homepage: 'https://constructs.dev/packages/cdk-sops-secrets',
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
    allowedUsernames: [
      'markussiebert',
      'renovate-bot',
      'renovate',
      'renovate[bot]',
    ],
  },

  name: 'cdk-sops-secrets',
  repositoryUrl: 'https://github.com/dbsystel/cdk-sops-secrets.git',
  bundledDeps: ['yaml'],
  // deps: [], /* Runtime dependencies of this module. */,
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
  '/codecov.yaml',
  '/.prettier*',
  '/.whitesource',
  '/.gitattributes',
);

goreleaserArtifactsNamespace = 'build-artifact-goreleaser';

additionalActions = [
  {
    name: 'Download zipper artifacts',
    uses: 'actions/download-artifact@v4',
    with: {
      name: 'zipper',
      path: 'assets',
    },
  },
];

project.buildWorkflow.preBuildSteps.unshift(...additionalActions);
project.buildWorkflow.preBuildSteps.push({
  name: 'Update snapshots: secret-inline',
  run: 'yarn run projen integ:secret-inline:snapshot',
});
project.buildWorkflow.preBuildSteps.push({
  name: 'Update snapshots: secret-asset',
  run: 'yarn run projen integ:secret-asset:snapshot',
});
project.buildWorkflow.preBuildSteps.push({
  name: 'Update snapshots: secret-multikms',
  run: 'yarn run projen integ:secret-multikms:snapshot',
});
project.buildWorkflow.preBuildSteps.push({
  name: 'Update snapshots: secret-manual',
  run: 'yarn run projen integ:secret-manual:snapshot',
});
project.buildWorkflow.addPostBuildSteps({
  name: 'Upload coverage to Codecov',
  uses: 'codecov/codecov-action@v4',
  with: {
    flags: 'cdk',
    directory: 'coverage',
  },
});
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
    if (['release'].includes(key)) {
      wf.jobs[key].steps.splice(5, 0, {
        name: 'Upload coverage to Codecov',
        uses: 'codecov/codecov-action@v4',
        with: {
          flags: 'cdk',
          directory: 'coverage',
        },
      });
    }
  });

  wf.addJob('gobuild', {
    name: 'gobuild',
    runsOn: 'ubuntu-latest',
    container: {
      image: 'golang:1.22-bullseye',
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
        name: 'Temporary workaround Checkout Issue #760 ',
        run: 'git config --global --add safe.directory /__w/cdk-sops-secrets/cdk-sops-secrets',
      },
      {
        name: 'Checkout',
        uses: 'actions/checkout@v4',
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
        name: 'Upload coverage to Codecov',
        uses: 'codecov/codecov-action@v4',
        env: {
          CODECOV_TOKEN: '${{ secrets.CODECOV_TOKEN }}',
        },
        with: {
          files: './coverage/coverage.out',
          flags: 'go-lambda',
        },
      },
      {
        name: 'Build',
        run: 'scripts/lambda-build.sh',
      },
      {
        name: 'Upload artifact',
        uses: 'actions/upload-artifact@v4',
        with: {
          name: 'gobuild',
          path: 'lambda/bootstrap',
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
        name: 'Temporary workaround',
        run: 'git config --global --add safe.directory /__w/cdk-sops-secrets/cdk-sops-secrets',
      },
      {
        name: 'Checkout',
        uses: 'actions/checkout@v4',
        with: {
          'fetch-depth': 0,
        },
      },
      {
        name: 'Download gobuild artifacts',
        uses: 'actions/download-artifact@v4',
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
        uses: 'actions/upload-artifact@v4',
        with: {
          name: 'zipper',
          path: 'assets/cdk-sops-lambda.zip',
        },
      },
    ],
  });
});

project.synth();
