# Practice 6 - Small Multi-Language App (Fully Dockerized)

---

## About This Practice

This practice is based on a project from [DevOps Directive Docker Course](https://github.com/sidpalas/devops-directive-docker-course/tree/main/05-example-web-application).

It is a **3-tier web application** where all components have been **fully Dockerized**:

- **React frontend**: Dockerized React/Vite client.
- **Node.js API**: Dockerized Node.js Express backend.
- **Golang API**: Dockerized Go backend.
- **PostgreSQL Database**: Dockerized relational database.

All services are orchestrated using **Docker Compose**.

---

## Architecture

```
[React Frontend] --> [Node.js API]
                    [Golang API] --> [PostgreSQL Database]
```

---

## Technologies Used

- React
- Node.js
- Golang
- PostgreSQL
- Docker
- Docker Compose

---

## How to Run the Project

### 1. Clone the Repository

```bash
git clone https://github.com/shiHisham/Docker-practice.git
cd Docker-practice/Practice6_Small_multilang_app_Dockerized
```

### 2. Build and Start All Services

```bash
docker-compose up --build
```

This command will build the Docker images for all services and start the containers.

---

## Services and URLs

| Service         | URL                         |
|-----------------|------------------------------|
| React Frontend  | http://localhost:1515         |
| Node.js API     | http://localhost:3000         |
| Golang API      | http://localhost:8080         |
| PostgreSQL      | localhost:5432 (Database port) |

---

## Dockerized Components

| Component       | Details                      |
|-----------------|-------------------------------|
| React Client    | Built using `node:18-alpine`   |
| Node.js API     | Built using `node:20`          |
| Golang API      | Built using `golang:1.20`      |
| PostgreSQL      | Image `postgres:15.1-alpine`   |

---

## Notes

- All services are fully containerized and connected via Docker networks.
- PostgreSQL data is persisted using a Docker volume (`pgdata`).
- Environment variables like `DATABASE_URL` are passed to backend services at runtime.

---

## Credit

Project based on the example from [DevOps Directive Docker Course](https://github.com/sidpalas/devops-directive-docker-course/tree/main/05-example-web-application).

---