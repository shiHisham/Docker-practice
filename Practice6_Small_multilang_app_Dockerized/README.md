# Practice 6 - Small Multi-Language App (Fully Dockerized)
---
# 📄 README for **Practice 6** (Fully Dockerized App)

```markdown
# Practice 6 - Small Multi-Language App (Fully Dockerized)

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

```plaintext
[React Frontend] --> [Node.js API]
                     [Golang API] --> [PostgreSQL Database]
```

## Technologies Used
* React
* Node.js
* Golang
* PostgreSQL (Database)
* Docker
* Docker Compose

## How to Run the Project (Practice 6)
    ```bash
    git clone https://github.com/shiHisham/Docker-practice.git
    cd Docker-practice/Practice6_Small_multilang_app_Dockerized
    ```

**Build and Start All Services**:
   ```bash
   docker-compose up --build
   ```
   This command will build the Docker images for all services and start the containers.