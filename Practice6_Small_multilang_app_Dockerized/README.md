# 🧪 Practice 6 - Small Multi-Language App (Fully Dockerized)

---

## 📚 About This Practice

This practice is based on a project from [DevOps Directive Docker Course](https://github.com/sidpalas/devops-directive-docker-course/tree/main/05-example-web-application).

It is a complete **3-tier web application**, now with all components **fully Dockerized** and orchestrated using **Docker Compose**:

- **React frontend**: Dockerized React/Vite client
- **Node.js API**: Dockerized Node.js Express backend
- **Golang API**: Dockerized Go backend
- **PostgreSQL**: Dockerized relational database with persistent storage

> This practice builds directly on [Practice 5](../Practice5_Small_multilang_app), where only PostgreSQL was Dockerized and all other components were running locally.  
> Now, everything runs inside containers.

---

## 🧩 Architecture

```
[ React Frontend ]
        ↓
[ Node.js API ] <--> [ Golang API ]
        ↓                  ↓
      [ PostgreSQL Database ]
```

- All services run in isolated Docker containers.
- Backend services communicate with PostgreSQL through the internal Docker network.
- Frontend connects to backend services through HTTP proxies using service names defined in Docker Compose.

---

## 🛠️ Technologies Used

- **React** (Vite)
- **Node.js** (Express)
- **Golang**
- **PostgreSQL**
- **Docker**
- **Docker Compose**

---

## 🚀 How to Run the Project

### 1. Clone the Repository

```bash
git clone https://github.com/shiHisham/Docker-practice.git
cd Docker-practice/Practice6_Small_multilang_app_Dockerized
```

### 2. Build and Start All Services

```bash
docker-compose up --build
```

This command will:
- Build Docker images for React, Node.js, and Go APIs
- Launch all services and connect them through a shared network
- Start the PostgreSQL container with persistent volume

---

## 🌐 Services and URLs

| Service         | URL                          |
|-----------------|------------------------------|
| React Frontend  | [http://localhost:1515](http://localhost:1515) |
| Node.js API     | [http://localhost:3000](http://localhost:3000) |
| Golang API      | [http://localhost:8080](http://localhost:8080) |
| PostgreSQL      | `localhost:5432` (database connection)          |

---

## ⚙️ Vite Proxy Configuration (React Frontend)

To allow the React frontend to communicate with the backend APIs from inside Docker, the **Vite dev server proxy** was updated:

```js
port: 1516,
proxy: {
  '/api/golang': {
    target: 'http://api-golang:8080',
    changeOrigin: true,
    rewrite: (path) => path.replace(/^\/api\/golang/, ''),
    secure: false,
  },
  '/api/node': {
    target: 'http://api-node:3000',
    changeOrigin: true,
    rewrite: (path) => path.replace(/^\/api\/node/, ''),
    secure: false,
  },
},
```

### 🔁 Key Points:
- The **hostnames** (`api-golang` and `api-node`) match the service names defined in `docker-compose.yml`
- Port `1516` matches the port exposed in the **React Dockerfile**
- This configuration allows internal service communication without relying on `localhost`

---

## 📦 Dockerized Components

| Component       | Dockerfile Base Image         |
|-----------------|-------------------------------|
| React Client    | `node:18-alpine`              |
| Node.js API     | `node:20`                     |
| Golang API      | `golang:1.20`                 |
| PostgreSQL      | `postgres:15.1-alpine`        |

Each service has its own `Dockerfile` and runs in isolation while sharing the same network and environment.

---

## 🧠 What This Practice Covers

- Full Dockerization of a multi-language application
- Building and managing containers with Dockerfiles
- Running multi-service apps with Docker Compose
- Setting up communication between frontend and backend inside containers
- React proxy setup for service discovery by container name
- Volume persistence and containerized PostgreSQL access

---

## 🙏 Credit

Project based on the example from [DevOps Directive Docker Course](https://github.com/sidpalas/devops-directive-docker-course/tree/main/05-example-web-application).

---


---

🧾 *README content written based on my own implementation and learning. Formatting and structure assisted by ChatGPT for clarity and presentation.*
