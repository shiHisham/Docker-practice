# 🧪 Practice 7 - Full Stack Multi-Language Microservices App
### 🐋 Fully Dockerized · Multi-Stage Builds · Dev/Prod Separation · Secrets · Makefile Automation · Healthchecks + Auto-Restart

---

## 📙 Overview

This project builds on [Practice 6](../Practice6_Small_multilang_app_Dockerized), evolving the same multi-language, 3-tier application into a secure, production-ready stack.

- 🏗️ Multi-stage Docker builds  
- 🔐 Secure container images (Distroless, Scratch)  
- ⚙️ Dev and Prod separation via Docker Compose  
- 🔁 Hot reload in dev with tools like `air`, `nodemon`, and Vite  
- 🔑 Secrets management using `.env` and Docker secrets  
- 📦 Minimal image sizes, non-root containers, and custom networking  
- 🧾 Full Makefile automation  
- 🔍 Healthcheck integration for auto-restart and monitoring

This stack is ideal for both learning and production-like deployments.

### ✨ New in Practice 7:
- ✅ Production healthchecks for all services
- ✅ Auto-restart on failure using `restart: unless-stopped`
- ✅ Diagnostic testing of health behavior
- ✅ Better loopback handling inside containers (`127.0.0.1` instead of `localhost`)
- ✅ Watch command for container status in Makefile
- ✅ Refined permission handling in Dockerfiles for reliable testing

This setup reflects **how real-world Dockerized microservices are structured and monitored.**

---

## ⚙️ Stack Overview

| Service        | Language / Tool    | Description                      |
|----------------|--------------------|----------------------------------|
| `api-golang`   | Go (Gin)           | Backend API #1                   |
| `api-node`     | Node.js (Express)  | Backend API #2                   |
| `api-react`    | React + Vite       | Frontend SPA                     |
| `postgres`     | PostgreSQL         | Shared relational database       |

All services are fully containerized, networked, and managed with Docker Compose.

---

## 🪩 Architecture & Communication

- Each service runs in its own container.
- A custom Docker network `MSSmall_app` enables internal communication via service names.
- Persistent volumes are used for PostgreSQL.
- React communicates with the APIs through Vite proxy rewrites:

```js
proxy: {
  '/api-go': {
    target: 'http://api-golang:8080',
    rewrite: path => path.replace(/^\/api-go/, '')
  },
  '/api-node': {
    target: 'http://api-node:3000',
    rewrite: path => path.replace(/^\/api-node/, '')
  }
}
```

---

## 🧪 Development Setup

 Each service has its own multi-stage `Dockerfile` with a `dev` target.
- React uses Vite hot reload on port `1516`
- Golang uses [`air`](https://github.com/cosmtrek/air) for hot reload
- Node.js uses `nodemon`
- `.env` files are used to inject environment variables in development.

> All services communicate using their service names (e.g. `api-node`, `api-golang`) — never `localhost`.

---

## 🚀 Production Setup

- **Multi-stage builds** strip dev dependencies for small, production-only images.
- **Secrets** are injected via `/run/secrets/...` from Docker.
- **Images are minimal**:
  - `api-golang`: `scratch`
  - `api-node`: `distroless` replaced with `node:alpine` for CLI access no longer `distroless` in this new version
  - `api-react`: built static site served via `nginxinc/nginx-unprivileged`
- PostgreSQL runs the same image in both dev and prod, with a custom entrypoint script for secret injection.
- Healthcheck endpoints or binaries used to detect readiness
- `restart: unless-stopped` for resilience
---

## 📊 Docker Image Size Comparison

| Service        | Practice 6 Size | Practice 7 Size | Reduction          |
|----------------|------------------|------------------|---------------------|
| React Frontend | 290MB            | 41.3MB           | ✅ ~85% smaller      |
| Node.js API    | 1.14GB           | 170.68MB         | ✅ ~83% smaller      |
| Golang API     | 978MB            | 16.2MB           | ✅ ~98% smaller      |
| PostgreSQL     | 243MB            | 243MB            | ≈ Same               |

---

## 🔐 Secrets & Environment Handling

- **Dev**: `.env` + `env_file:` in Compose
- **Prod**: `Docker secrets` via mounted files and entrypoint scripts for PostgreSQL
- **Secrets** are injected into containers at runtime, not hardcoded in images.
- All backend apps support fallback logic for secret loading

---

## 📊 Healthcheck Integration (New)

### How It Works:
- Custom endpoints `/ping` created in Go, Node, and React (via Nginx)
- Healthchecks use `wget` or binary probes on `127.0.0.1` instead of `localhost`, because `localhost` resolves to the container's own DNS, which can cause issues.
- Docker tracks container health using this data

### Example: Compose entry
```yaml
restart: unless-stopped
healthcheck:
  test: ["CMD", "wget", "--spider", "--quiet", "http://127.0.0.1:1516"]
  interval: 10s
  timeout: 5s
  retries: 3
  start_period: 10s
```

### Key Fix:
> Replacing `localhost` with `127.0.0.1` was necessary in healthcheck scripts to avoid container DNS issues.

### Confirmed Results:
- Container correctly reports `healthy`
- Manual crash (e.g., `kill 1`) causes auto-restart
- Healthcheck failures alone do **not** trigger restarts (Docker behavior) unless the container crashes or exits.

---

## 📃 Makefile Automation

```makefile
# Monitor container status like 'watch'
watch:
	while true; do clear; docker ps; sleep 3; done
```

Includes:
- Dev commands (`up-dev`, `up-react-only-dev`, etc.)
- Prod commands (`up-prod`, `up-react-only-prod`, etc.)
- Make clean for removing all containers, images, and volumes
- Watch command for monitoring container status using `make watch`
---

## 🧰 Lessons, Fixes & Gotchas

| Challenge                            | Fix / Learning Outcome                                     |
|-------------------------------------|-------------------------------------------------------------|
| Healthcheck failed on `localhost`   | Used `127.0.0.1` to avoid internal container DNS issues     |
| React healthcheck failing silently  | Nginx wasn't exposing port in config, verified via `netstat`|
| `distroless` too minimal             | Switched to `node:alpine` to allow CLI testing & access     |
| Deleting code didn't crash the app  | Node kept code in memory — crash needs to kill PID 1       |
| `restart: unless-stopped` confusion | Only restarts on actual exit/crash, not healthcheck alone   |
| `mv` failed in container            | Files owned by root, container user was `node`              |

---

## 💡 What This Practice Demonstrates

- 🌍 Realistic multi-language, multi-service architecture
- 🔧 Secure and optimized Docker builds
- 🔎 Secrets management best practices
- 🪡 Healthcheck and restart integration
- ⚖️ Permission and ownership awareness in builds
- 📈 Monitoring and debugging strategy using `docker ps`, logs, and simulated failures

---

## ✅ Final Summary

This practice has gone beyond simple containerization into **operational robustness**:
- Services report their health
- Crashes auto-recover
- Production is monitored and stable

---

🗏️ *README content based on actual project implementation and testing. Format and documentation assisted by ChatGPT for clarity and completeness.*