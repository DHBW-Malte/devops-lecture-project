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

The second one is coming form google. It get triggered with a push on our main. The core-function of it is to check if there were changes in any of the services. If so, release-please is using the format of our commit to generate an release changelog. There it will list all `fix` and `feat` commits with their commit messages. Based on the commits please-release will decide what kind of release it will be. With only `fix` commits the release will be a **patch** release and the version number of the affected service will increased by one at the patch position in the version-number. With a `feat` commit the release would become a **minor** release and with a `!feat` it would be a major (breaking) release.

### Third Workflow - `publish.yml`

The third and last workflow is responsible for the docker image creation and publishing. It get triggered if the release-please changelog & release pr get pushed on the new created release tag, e.g. auth-service-1.2.3. The publish workflow makes also use of this release tag for the next steps, for that it extract the service and the version number out of the tag. Then after the login on Docker Hub, it build and push for every new released service an image. For the build part it is using the Dockerfile which is expecting the service as a build argument. With that the correct image get build and then published on Dockerhub with the new version number.

---

## Kubernetes-Setup with Minikube

We now also provide a local Kubernetes Cluster for the microservices of our devops-shop backend.

### Prerequisites

- [Minikube](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download)
- [kubectl](https://kubernetes.io/docs/reference/kubectl/)

### Manifest Structure

Each microservice has its own manifest file in the `kubernetes/` directory.

Every manifest contains two resources:

1. **Deployment:** Defines which container image to run and how many replicas (pods) to create. Kubernetes ensures the desired number of pods is always running.
2. **Service:** Provides a stable internal DNS name and load balances incoming traffic across all pods of the deployment.

We choose following numbers of replicas for the services:

- **auth-service:** 4 replicas with the image `dhbwmalte/auth-service:0.3.1`
- **products-service:** 6 replicas with the image `dhbwmalte/products-service:0.2.1`
- **checkout-service:** 8 replicas with the image `dhbwmalte/checkout-service:0.3.1`

### Getting Started

Start the Minikube cluster:

```bash
minikube start
```

Deploy all services:

```bash
cd kubernetes
kubectl apply -f .
```

Verify that deployments and services are running:

```bash
kubectl get deployments
kubestl get services
```

### Testing a Service Locally

The services use ClusterIP by default, which means they are only reachable from within the cluster. To access a service from your machine, use `kubectl port-forward`:

```bash
kubectl port-forward service/products-service 8080:8080
```

Then, in another terminal:

```bash
curl http://localhost:8080/products
```

# DevSecOps – Container Image Scanning

This project implements a **DevSecOps pipeline** following the Shift-Left Security principle. Security scanning is integrated directly into the CI/CD pipeline and acts as a gate before any image is published to Docker Hub.

## How it works

Every release tag push triggers the `Publish Docker Images` workflow. The workflow follows a strict **build → scan → publish** order, ensuring that no vulnerable image ever reaches the registry.

```bash
docker build (local) → Generate SBOM → Scan for CVEs → Push to Docker Hub
```

The push step is only reached if the vulnerability scan passes. If a critical CVE is found, the pipeline aborts and nothing is published.

## SBOM Generation

A **Software Bill of Materials (SBOM)** is generated from the locally built container image using [Syft](https://github.com/anchore/syft) via `anchore/sbom-action`. The SBOM lists every package and dependency inside the image — comparable to an ingredient list.

- Format: `CycloneDX JSON`
- The SBOM is uploaded as a workflow artifact and attached to the corresponding GitHub Release for auditing purposes.

## Vulnerability Scanning

The generated SBOM is scanned against known CVE databases using [Grype](https://github.com/anchore/grype) via `anchore/scan-action`.

| Setting | Value |
|---|---|
| Severity cutoff | `critical` |
| Fail build on finding | `true` |

Findings below `critical` are reported but do not block the pipeline. Adjust `severity-cutoff` to `high` in `.github/workflows/publish.yml` if stricter enforcement is needed.

## Required Repository Permissions

The workflow requires the following permissions to attach the SBOM to the GitHub Release:

```yaml
permissions:
  contents: write
  actions: write
```

These are explicitly scoped — no broader permissions are granted.

## Tools Used

| Tool | Purpose |
|---|---|
| [Syft (anchore/sbom-action)](https://github.com/anchore/sbom-action) | SBOM generation from container image |
| [Grype (anchore/scan-action)](https://github.com/anchore/scan-action) | CVE scanning of the SBOM |
| [docker/build-push-action](https://github.com/docker/build-push-action) | Building the image locally before scanning |
