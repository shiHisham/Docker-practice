# 🛠 Full Stack Dockerized Microservices App

A complete setup for a small full-stack microservices app using **Go**, **Node.js**, **React**, and **PostgreSQL**, containerized with **Docker**. The stack is clean, optimized, and ready for both development and production use.

---

## ⚙️ Stack Overview

| Service      | Language / Tool | Description                       |
|--------------|------------------|-----------------------------------|
| `api-golang` | Go (Gin)         | Backend API #1                    |
| `api-node`   | Node.js (Express)| Backend API #2                    |
| `api-react`  | React + Vite     | Frontend (SPA)                    |
| `postgres`   | PostgreSQL       | Shared database                   |

---

## 🧪 Development Environment

- Each service has its own multi-stage `Dockerfile` with a `dev` target.
- React runs with Vite hot reload on port `1516`
- Golang uses `air` for hot reload
- Node uses `nodemon`

> Services communicate using their **Docker Compose service name** (e.g. `api-node`, `api-golang`), not `localhost`.

### Example Vite Proxy Setup:
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

## 🚀 Production Setup

- `api-golang`: golang image + `/run/secrets/DATABASE_URL`
- `api-node`: distroless image + `/run/secrets/DATABASE_URL`
- `api-react`: built with Vite, served via `nginxinc/nginx-unprivileged`
- `postgres`: same container used for both environments


## 🔐 Secrets Management

- Dev: uses `.env` + `env_file:` in Compose
- Prod: uses Docker Swarm secrets (`/run/secrets/…`)

Go example:
```go
os.Getenv("DATABASE_URL") // for dev
os.ReadFile("/run/secrets/DATABASE_URL") // for prod
```

To create the secret:
```bash
docker secret create DATABASE_URL secrets/DATABASE_URL.txt
```

---

## 📦 Docker & Compose Highlights

- Each service has its own `.dockerignore`
- Multi-stage Dockerfiles
- Clean separation of dev vs prod via `docker-compose.dev.yaml` & `docker-compose.prod.yaml`
- Network aliasing via service names (e.g. `api-node`, `api-golang`)

---

## 🛠 Makefile (Optional Shortcuts)

```Makefile
# Development environment
up-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build
up-golang-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d postgres api-golang
up-node-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build postgres api-node
up-react-only-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build postgres api-react

# Production
up-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d
up-golang-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build -d postgres api-golang
up-node-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build postgres api-node
up-react-only-prod:
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml up --build postgres api-react
# Clean everything
clean:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml down -v
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml down -v
```

---

## ✅ Summary

This setup gives you:
- 🔁 Hot reload in dev
- 🔐 Secure secrets handling in prod
- 🐳 Minimal, layered images
- 🧠 Easy switching between dev/prod
- 🌍 Networked services with DNS names