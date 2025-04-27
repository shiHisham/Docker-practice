# Practice 11 - Docker Security & Image Scanning 🔒🔁🚪

---

## 📚 Overview

This practice focuses on securing Docker containers and scanning images for vulnerabilities using **Trivy**. It builds upon the previous multi-service project consisting of:

- 🔁 API in **Golang**
- 🔁 API in **Node.js**
- 🔁 Frontend in **React**
- 🔁 PostgreSQL database
- 🚪 Reverse proxy with **Nginx**

The objective was to detect, analyze, and reduce security vulnerabilities before production deployment.

---

## 🧩 Key Changes from Previous Practices

- A new directory `Practice11_Security_ImageScanning/` was created.
- Application code, Dockerfiles, and Compose files were copied from the previous project.
- Networks were separated into:
  - `MSSmall_app_backend`: For API services and database.
  - `MSSmall_app_frontend`: For frontend and proxy.

---

## 🛠️ Tools Used

- **Trivy**: Main tool for vulnerability scanning.
- **npm audit fix**: Used to automatically fix vulnerabilities in Node.js projects.

---

## 📊 Vulnerability Scan Results

### Before Fixes

- **api-golang**: 15 vulnerabilities (CRITICAL: 1, HIGH: 4, MEDIUM: 10)
- **api-node**: 21 vulnerabilities (CRITICAL: 1, HIGH: 13, MEDIUM: 3, LOW: 4)
- **client-react**: 85 vulnerabilities (CRITICAL: 4, HIGH: 16, MEDIUM: 57, LOW: 6)

### After Fixes

- **api-golang**: 0 vulnerabilities
- **api-node**: 0 vulnerabilities
- **client-react**: 0 vulnerabilities



## 🛡️ Fixing and Hardening Process

### Client-React

- Updated base image from `nginxinc/nginx-unprivileged:1.23-alpine` to `nginxinc/nginx-unprivileged:alpine3.21-perl`, reducing vulnerabilities.
- Updated Nginx configuration with:
  - `server_tokens off` (hide version info).
  - `client_max_body_size 1M` (limit upload size).
  - HTTP methods restriction (`GET`, `POST`, `HEAD`, `OPTIONS` only).
- Final scan: **0 vulnerabilities**.

### Node.js API

- Updated vulnerable libraries manually using:
  ```bash
  npm audit fix --force
  ```
- Updated Node base image to `node:lts-alpine3.21`.
- Final scan: **0 vulnerabilities**.

### Golang API

- Updated `go.mod` and `go.sum` packages (`gin-gonic`, `pgx`, `x/crypto`).
- Used multi-stage builds with a final **scratch** image.
- Set non-root user in production image.
- Final scan: **0 vulnerabilities**.

### Postgres

- Official `postgres:15.1-alpine` image was used without modification.
- Identified that `read_only: true` cannot be applied inside Postgres due to internal write operations.

---

## 🏗️ Infrastructure and Security Improvements

- 🧩 **Multi-stage builds** for all services.
- 🛡️ **Non-root users** for Golang, Node.js, and React in production.
- 🔒 **Minimal production images** (Distroless/Scratch/Nginx-unprivileged).
- 🔁 **Separated networks** for frontend and backend communication.
- 🔄 **Healthchecks** for all services.
- 📦 **Docker Secrets** used for sensitive data.
- 🗂️ **Volumes** and **tmpfs** usage optimized.

---

## ⚙️ Problems Encountered and Solutions

| Issue                                 | Solution                                                            |   |   |
| ------------------------------------- | ------------------------------------------------------------------- | - | - |
| Nginx `Permission denied` error       | Added `pid /tmp/nginx.pid` or used default image configs.           |   |   |
| Duplicate `pid` directive             | Ensured no manual overriding inside `/etc/nginx/conf.d/` files.     |   |   |
| Postgres error with `read_only: true` | Dropped `read_only` for database service to allow essential writes. |   |   |

---

## 📌 Key Learnings

- How to perform vulnerability scanning on Docker images.
- How to update dependencies manually and safely.
- How to optimize and harden Dockerfiles.
- How to configure secure Nginx setups in containers.
- How to handle and interpret vulnerability reports in real-world scenarios.

---

🧾 *README content written based on personal implementation. Formatted and structured with ChatGPT for clarity.*