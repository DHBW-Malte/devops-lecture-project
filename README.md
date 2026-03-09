# devops-lecture-project-2025
---
Casimir Bomans & Malte Opderbeck
---

## Branch & Commit Guide

### Branch

This repository follows the **GitHub Flow**, which means:

* The **main** branch is always deployable
* Branch naming is using prefixes:
  * **Feature Branches:** `feature/<short description>`
  * **Fixing Branches:** `fix/<bug>`
  * **Maintenance Branches:** `chore/maintenance`
  * **Documentation Branches:** `docs/<topic>`

### Commits

Commit messages in this repository are in the **Conventional Commits** format, which means:

* The general format of a commit message is `type(scope): short description`
  * E.g. feat(api): add product search endpoint, fix(auth): handle expired jwt
* Commit types:
  * **Feature:** `feat`
  * **Bugfix:** `fix`
  * **Documentation:** `docs`
  * **Refactoring:** `refactor`
  * **Testing:** `test`
  * **Building, CI & Tooling:** `chore`
* General rule: **Commit small and often**

### PR-Requests

The **description** of a pull request should contain at least:
 * What has been changed?
 * Why?

---

## Dockerization

The devops-shop is now available as a container image on Docker Hub.

### Pull the image

```bash 
docker pull dhbwmalte/devpos-shop:v0.1.0
```


### Run the container 


```bash 
docker run --rm -p 8080:8080 dhbwmalte/devpos-shop:v0.1.0
```


The shop will be available at:
`http://localhost:8080`

### Build the Image Locally

If you prefer to build the image yourself:

```bash 
docker build -t devops-shop:local .
docker run --rm -p 8080:8080 devops-shop:local
```
---

## CI/CD-Pipeline with Docker Image Release

With the implementation of three GitHub workflows we created a CI/CD-Pipeline which ends in a release of three images on Docker Hub, one per service of our devops-shop (auth, products & checkout).

### First Workflow - `go.yml`

The first workflow get triggered with each push or pull request on our main. With this workflow we **build** and **test** each shop-server-service.

### Second Workflow - `release-please.yml`

The second one is coming form google. It get triggered with a push on our main. The core-function of it is to check if there were changes in any of the services. If so, release-please is using the format of our commit to generate an release changelog. There it will list all `fix` and `feat` commits with their commit messages. Based on the commits please-release will decide what kind of release it will be. With only `fix` commits the release will be a **patch** release and the version number of the affected service will increased by one at the patch position in the version-number. With a `feat` commit the release would become a **minor** release and with a `feat!` it would be a major (breaking) release.

### Third Workflow - `publish.yml`

The third and last workflow is responsible for the docker image creation and publishing. It get triggered if the release-please changelog & release pr get pushed on the new created release tag, e.g. auth-service-1.2.3. The publish workflow makes also use of this release tag for the next steps, for that it extract the service and the version number out of the tag. Then after the login on Docker Hub, it build and push for every new released service an image. For the build part it is using the Dockerfile which is expecting the service as a build argument. With that the correct image get build and then published on Dockerhub with the new version number.

---
