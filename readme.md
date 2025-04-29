# 🐳 Docker Practice Journey – From Basics to Production

[![Open to Contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat-square)](https://github.com/shiHisham/Docker-practice/issues)
[![Docker Hub](https://img.shields.io/badge/DockerHub-hichamshih-blue)](https://hub.docker.com/u/hichamshih)
[![GHCR](https://img.shields.io/badge/GHCR-shihisham--images-purple)](https://github.com/shihisham?tab=packages)

---

Welcome! This repository is a hands-on learning journey through Docker — starting from "Hello World" and advancing to multi-service applications with CI/CD, security scanning, and reverse proxy architecture.

I’m currently learning containerization and DevOps practices as part of my **preparation for a career in Cloud & Data Engineering**. Mastering tools like Docker is essential to build scalable, reproducible, and production-ready data systems.

---

## 📦 What Are Containers and Why We Should Learn Them?

Containers are a lightweight way to **package and isolate applications**. Unlike traditional virtual machines, they don’t include a full OS — just the binaries, libraries, and code your app needs.

### ✅ Why Containers Matter:
- **Portability**: Run anywhere — local, staging, or cloud
- **Isolation**: Prevent dependency conflicts
- **Efficiency**: Share the host kernel (less overhead)
- **Consistency**: Works on my machine ≠ problem solved
- **Automation**: Core to modern CI/CD pipelines

### 🚀 Best Practices for Using Containers:
- Use **multi-stage builds** to reduce image size
- Run apps as **non-root users**
- Use **named volumes** for persistent data
- Separate **dev and prod environments**
- Tag images with **version info**, not just `latest`
- Scan images for vulnerabilities before deploying

---

## 🐳 Why Docker?

Docker is the most widely adopted container platform. It provides a powerful CLI and rich ecosystem with:
- Docker CLI & API
- Docker Compose
- Docker Hub & Docker Desktop
- Seamless integration with CI/CD tools

### 🔁 Other Options Include:
| Tool       | Description                              |
|------------|------------------------------------------|
| **Podman** | Docker-compatible CLI, daemonless        |
| **containerd** | Lightweight runtime used by Kubernetes |
| **Buildah** | Focused on building images only         |

### ✅ Why I Chose Docker:
- Ubiquity in the job market
- Huge ecosystem (Hub, Compose, Dev Environments)
- Rich documentation and tutorials
- Ideal for learning and production

---

## 📘 Table of Contents

Here’s a complete index of all 11 practices, categorized by level:

### 🟢 Beginner
- [Practice 01](./Practice1_hello_world): Hello Docker World (Python script)
- [Practice 02](./Practice2_web_server): Flask Web Server in Docker
- [Practice 03](./Practice3_docker_compose): Python + PostgreSQL via Docker Compose
- [Practice 04](./Practice4_dokcer_Volumes): PostgreSQL + Docker Volumes

### 🟡 Intermediate
- [Practice 05](./Practice5_Small_multilang_app): Partially Dockerized Multi-Language App
- [Practice 06](./Practice6_Small_multilang_app_Dockerized): Fully Dockerized Multi-Language App
- [Practice 08](./Practice8_Publishing_Dockerhub): Publishing Docker Images to Registries

### 🔴 Advanced
- [Practice 07](./Practice7_Small_multilang_app_Dockerized_Multi stage): Microservices, Secrets, Healthchecks, Makefile Automation
- [Practice 09](./Practice9_CICD_GitHubActions_FullStack): Full CI/CD with GitHub Actions + Docker Hub
- [Practice 10](./Practice10_ReverseProxy_TLS): Reverse Proxy + SSL with Nginx
- [Practice 11](./Practice11_Security_ImageScanning): Docker Image Scanning + Security Hardening

---

## 🛠️ Tech Stack & Topics

- Docker (CLI, Dockerfiles, Compose, Volumes, Secrets)
- Multi-language stack: React + Node.js + Golang + PostgreSQL
- Nginx as Reverse Proxy
- Trivy for image vulnerability scanning
- GitHub Actions for CI/CD
- Docker Hub & GitHub Container Registry (GHCR)
- Makefile automation
- Dev/Prod environment separation

---

## 💬 About Me

I'm a software engineer with over 5 years of experience, primarily in backend and full-stack development. Recently, I’ve been diving deep into **containerization, DevOps, and cloud-native tools** as part of my transition toward **Cloud & Data Engineering**. While my end goal is to build powerful, scalable data systems, I strongly believe in **mastering DevOps and containerization fundamentals** — because production-ready data pipelines must be well-architected, portable, and secure.

This repo is a big step on that journey 🚀

---

## 🤝 Contributions Welcome

I'm still learning — so some practices may **miss best practices**, or there may be **better ways** to structure or secure things.

If you spot something that could be improved or want to suggest a new **Docker challenge**, I’d love to learn from it. Feel free to open an issue or PR 💬

---

## 📦 Repositories & Packages

- **Docker Hub**: [https://hub.docker.com/u/hichamshih](https://hub.docker.com/u/hichamshih)
- **GitHub Packages**: [https://github.com/shihisham?tab=packages](https://github.com/shihisham?tab=packages)

---

## 📜 License

This project is licensed under the MIT License.

---

Thanks for visiting! 🌊 Happy containerizing!