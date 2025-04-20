# 🐳 Practice 8 – Docker Image Publishing to Docker Hub & GitHub Packages

---

## 📚 Overview

This small project demonstrates how to build a Docker image using the `scratch` base and publish it to both **Docker Hub** and **GitHub Container Registry (GHCR)**.

Although the image itself is intentionally minimal (`FROM scratch`), the focus is on mastering the full **build, tag, login, and push workflow** using Docker CLI and Makefile automation.

This is an essential DevOps practice that helps you get comfortable working with image registries and release pipelines.

---

## 🏗️ What’s in This Practice

### 🔧 Technologies Used
- **Docker**
- **Makefile**
- **Docker Hub**
- **GitHub Container Registry (GHCR)**

### 🐚 Base Image
```Dockerfile
FROM scratch
```

> A zero-layer image used here for simplicity and size.

---

## 📦 Makefile Commands

The included `Makefile` automates the process of building and pushing images:

```makefile
# Build a scratch-based image
make build-scratch-image
# Push to Docker Hub (latest + v1)
make push-dockerhub
# Push to GitHub Container Registry (GHCR)
make push-github-packages
```

> Make sure you’re logged in to Docker Hub and GitHub Packages before pushing. See instructions below.

---

## 🔑 Authentication Setup

### 1. Docker Hub

- Create a Docker Hub account (if not already done)
- Login via CLI:
  ```bash
  docker login
  ```

### 2. GitHub Container Registry (GHCR)

- Generate a **GitHub Personal Access Token (PAT)** with scopes:
  - `write:packages`
  - `read:packages`
- Login using:
  ```bash
  echo <YOUR_GITHUB_PAT> | docker login ghcr.io -u <yourusername> --password-stdin
  ```

---

## 🏷️ Tagging & Pushing

### Docker Hub
```bash
docker tag myfirst-scratch-image hichamshih/myfirst-scratch-image
docker push hichamshih/myfirst-scratch-image

docker tag myfirst-scratch-image hichamshih/myfirst-scratch-image:v1
docker push hichamshih/myfirst-scratch-image:v1
```

### GitHub Packages (GHCR)
```bash
docker tag myfirst-scratch-image ghcr.io/shihisham/myfirst-scratch-image
docker push ghcr.io/shihisham/myfirst-scratch-image

docker tag myfirst-scratch-image ghcr.io/shihisham/myfirst-scratch-image:v1
docker push ghcr.io/shihisham/myfirst-scratch-image:v1
```

---

## ✅ Summary of What This Practice Demonstrates

- 🔧 Building Docker images from scratch
- 🏷️ Tagging images with custom and versioned tags
- 🔐 Authenticating to Docker Hub and GHCR using CLI
- 📦 Publishing and managing container images in registries
- 📄 Automating the entire process via Makefile

---

## 🔍 Where to View the Images

- **Docker Hub Profile**:  
  [https://hub.docker.com/u/hichamshih](https://hub.docker.com/u/hichamshih)

- **GitHub Packages**:  
  [https://github.com/shihisham?tab=packages](https://github.com/shihisham?tab=packages)

---

## 🧠 Why This Matters

Pushing Docker images to a registry is a key part of any real CI/CD pipeline. This practice helps reinforce:

- Registry authentication
- Clean image tagging/versioning
- Pushing to both public and private repositories

---


---

🧾 *README content written based on my own implementation and learning. Formatting and structure assisted by ChatGPT for clarity and presentation.*
