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
  * **Refactoring Branches:** `refactor/<short description>`

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
