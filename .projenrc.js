const { awscdk } = require("projen");
const project = new awscdk.AwsCdkConstructLibrary({
  author: "Markus Siebert",
  authorAddress: "dev@markussiebert.com",
  cdkVersion: "2.1.0",
  defaultReleaseBranch: "main",

  name: "cdk-sops-secrets",
  repositoryUrl: "https://github.com/markussiebert/cdk-sops-secrets.git",

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

project.buildWorkflow.preBuildSteps.unshift({
  name: "download go artifact",
  uses: "actions/download-artifact@v2",
  with: {
    name: "build-artifact-go",
    path: "dist/*",
  },
})

console.log(project.github.workflows.map((wr) => wr.name));

const fixme = project.github.workflows.filter((wf) =>
  ["build", "release"].includes(wf.name)
);
fixme.forEach((wf) => {
  Object.keys(wf.jobs).forEach((key) => {
    if(key === "build") {
      console.log(wf.jobs[key])
    } else {
      wf.jobs[key].steps.unshift({
        name: "download go artifact",
        uses: "actions/download-artifact@v2",
        with: {
          name: "build-artifact-go",
          path: "dist/*",
        },
      },)
    }
    wf.jobs[key] = {...wf.jobs[key], needs: 'goreleaser'}
  });
  wf.addJob("goreleaser", {
    name: "goreleaser",
    runsOn: "ubuntu-latest",
    on: {
      pull_request: null,
      push: null,
    },
    permissions: {
      contents: "write",
    },
    steps: [
      {
        name: "Checkout",
        uses: "actions/checkout@v2",
        with: {
          "fetch-depth": 0,
        },
      },
      {
        name: "Fetch all tags",
        run: "git fetch --force --tags",
      },
      {
        name: "Set up Go",
        uses: "actions/setup-go@v2",
        with: {
          "go-version": 1.18,
        },
      },
      {
        name: "Run GoReleaser",
        uses: "goreleaser/goreleaser-action@v2",
        with: {
          distribution: "goreleaser",
          version: "latest",
          args: "release --rm-dist",
        },
        env: {
          GITHUB_TOKEN: "${{ secrets.GITHUB_TOKEN }}",
        },
      },
      {
        name: "Upload artifact",
        uses: "actions/upload-artifact@v2.1.1",
        with: {
          name: "build-artifact-go",
          path: "dist/*",
        },
      },
    ],
  });
});

project.synth();
