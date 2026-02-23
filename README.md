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

```
```bash 
docker pull dhbwmalte/devpos-shop:v0.1.0
```
```
```

### Run the container 


```
```bash 
docker run --rm -p 8080:8080 dhbwmalte/devpos-shop:v0.1.0
```
```

The shop will be available at:
`http://localhost:8080`

### Build the Image Locally

If you prefere to build the image yourself:

```bash 
docker build -t devops-shop:local .
docker run --rm -p 8080:8080 devops-shop:local
```
```
```
```
```
