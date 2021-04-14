# PRify

PRify is a tool that interacts with Git and Git Servers fo you, simplifying continous delivery workflows. It's primarily designed to work with 'GitOps' style repos. 

There are a number of pain points around using git in a pipeline it aims to simplify:

- Branching and pushing, code for this is often re-written for each pipeline or CI system
- Creation of PRs, which many pipelines don't do just because integrating with the git server adds more work
- Duplication of pipelines. In many cases, different applications will be stored in different folders in the same repo (i.e. in a GitOps repo), but may have duplicated pipelines to give visibility of what's changing. PRify provides an alternative way to get this visibility by automatically raising separate PRs for changes in different folders

In a nutshell, PRify pushes code and automatically creates separate PRs for subfolders.

# Installation

Download the latest binaries from the [releases page](https://github.com/benjvi/PRify/releases). 

A dockerfile is also provided at `benjvi/prify:latest` for use in docker-based CI systems.

# Getting Started 

Create a `prify.yml` file  in the directory you will make changes to in CI, using [this example](https://github.com/benjvi/PRify/blob/main/e2e/examples/mvp/prify.yml) as a reference for the fields available.

In your CI pipeline:

- Make changes to the directory
- Setup any git or github credentials
- Navigate to the directory containing the changes and execute `prify run`

Depending on the configuration you set, the changes should been pushed to the git server and PRs may have been created.

# Example

I have an example ["apps-gitops" repo](https://github.com/benjvi/apps-gitops) that [uses PRify](https://github.com/benjvi/apps-gitops/blob/main/nonprod-cluster/prify.yml) to automatically push changes to to non-prod from [angular](https://github.com/benjvi/angular-realworld-example-app/blob/buildpacks/Jenkinsfile) and [(TODO) spring](https://github.com/benjvi/minimal-spring-web-demo/blob/main/Jenkinsfile) CI pipelines. In [a workflow](https://github.com/benjvi/apps-gitops/blob/main/.github/workflows/promote-to-prod.yml) triggered by changes to non-prod, PRify is [used again](https://github.com/benjvi/apps-gitops/blob/main/prod-cluster/prify.yml) to create PRs for the changes. 

