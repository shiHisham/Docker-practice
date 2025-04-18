# Docker Image Push Practice – Docker Hub & GitHub Packages

This small project is a practice for building and pushing a Docker image to both Docker Hub and GitHub Container Registry (GHCR).

The image is based on `scratch`, just to keep it minimal. The goal here is to make sure everything works end-to-end: building, tagging, authentication, and pushing to both registries.

---

## Steps I followed

### 1. Docker Hub

- Created a Docker Hub account
- Logged in using the terminal: `docker login`
- Built the image using a Makefile
- Tagged and pushed the image to Docker Hub

You can check the image on my Docker Hub profile:
[https://hub.docker.com/u/hichamshih](https://hub.docker.com/u/hichamshih)

---

### 2. GitHub Packages (GHCR)

- Created a personal access token (PAT) from GitHub with the following scopes:
  - `write:packages`
  - `read:packages`
- Logged in to GitHub Container Registry using:
  ```bash
  echo ghp_xxxxxxxxxxxxxxxxxxxxxx | docker login ghcr.io -u shihisham --password-stdin
  ```
- Tagged and pushed the image to GHCR using the Makefile

Once pushed, the image appears under the "Packages" tab on GitHub profile:
[https://github.com/shihisham?tab=packages](https://github.com/shihisham?tab=packages)

---

## Notes

- Makefile is used to automate build and push steps
- Tags include both `latest` (default) and a custom version (`v1`)
- This was mainly a practice for workflow testing, not an actual image for production use

---

## Author

**Hicham Shih**  
[GitHub](https://github.com/shiHisham)