# 🧪 Practice 7 - Full Stack Multi-Language Microservices App  
### 🐳 Fully Dockerized · Multi-Stage Builds · Dev/Prod Separation · Secrets · Makefile Automation

---

## 📚 Overview

This project is a refined version of [Practice 6](../Practice6_Small_multilang_app_Dockerized), evolving the same small multi-language, 3-tier application into a **secure, optimized, production-ready stack** using Docker best practices. It includes:

- 🏗️ Multi-stage Docker builds  
- 🔐 Secure container images (Distroless, Scratch)  
- ⚙️ Dev and Prod separation via Docker Compose  
- 🔁 Hot reload in dev with tools like `air`, `nodemon`, and Vite  
- 🔑 Secrets management using `.env` and Docker secrets  
- 📦 Minimal image sizes, non-root containers, and custom networking  
- 🧾 Full Makefile automation  

This stack is ideal for both learning and production-like deployments.

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

## 🧩 Architecture & Communication

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

## 🧪 Development Environment

- Each service has its own multi-stage `Dockerfile` with a `dev` target.
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
  - `api-node`: `distroless`
  - `api-react`: built static site served via `nginxinc/nginx-unprivileged`
- PostgreSQL runs the same image in both dev and prod, with a custom entrypoint script for secret injection.

---

## 📉 Docker Image Size Comparison

| Service        | Practice 6 Size | Practice 7 Size | Reduction          |
|----------------|------------------|------------------|---------------------|
| React Frontend | 290MB            | 41.3MB           | ✅ ~85% smaller      |
| Node.js API    | 1.14GB           | 167MB            | ✅ ~85% smaller      |
| Golang API     | 978MB            | 16.2MB           | ✅ ~98% smaller      |
| PostgreSQL     | 243MB            | 243MB            | ≈ Same               |

> Achieved through multi-stage builds, production-only dependencies, and lean base images.

---

## 🔐 Secrets & Environment Management

- **Development** uses `.env` files and `env_file:` in Compose.
- **Production** uses Docker secrets securely mounted at runtime.

### Example: Golang Secret Handling (./api-golang/main.go)
```go
// Dev
os.Getenv("DATABASE_URL")

// Prod
os.ReadFile("/run/secrets/DATABASE_URL")
```

> Apps use fallback logic to detect their current environment.

---

## 🧾 Makefile Automation

A custom `Makefile` simplifies local development, production setup, and cleaning:

```makefile
# Development
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

# Clean
clean:
	docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml down -v
	docker-compose -f docker-compose.yaml -f docker-compose.prod.yaml down -v
```

---

## 🧠 What This Practice Demonstrates

- Real-world microservices architecture with frontend/backend/db layers
- Secure, minimal Docker images using multi-stage builds
- Clean and consistent dev/prod environment separation
- Automated workflows using Makefile
- Safe secret injection with Docker secrets
- Multi-container orchestration with Docker Compose
- Custom networking and DNS-based communication

---

## ✅ Why This Setup Matters

This project isn't just a Docker tutorial — it’s a **blueprint for professional-grade full-stack app deployment**:

- 🔐 Secure by default (non-root, secrets, minimal images)
- 🧱 Modular and scalable architecture
- 🧪 Dev-friendly with live reload and isolated services
- 🚀 Prod-ready with optimized builds and secret handling
- 🔄 Easy to extend with CI/CD, reverse proxies (Nginx, Traefik), or orchestration (Kubernetes, ECS)

---