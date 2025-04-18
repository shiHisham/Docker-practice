# 🐳 Practice 3: Docker Compose App with PostgreSQL

This practice demonstrates how to connect a Python application to a PostgreSQL database using Docker Compose.

The goal is to:
- Set up multiple containers (Python app + PostgreSQL)
- Enable communication between services using Compose networking
- Install system dependencies to support PostgreSQL client libraries (`psycopg2`)

---

## 📦 What’s Inside

- A Python app that connects to PostgreSQL after a short wait
- A `Dockerfile` with build steps and necessary system dependencies
- A `docker-compose.yml` file that starts both services together
- A `requirements.txt` listing Python packages to install (`psycopg2`)

---

## ⚙️ Build Fixes / Notes

> While working on this app, the following fixes were applied:

- Changed base image from `python:3.13-slim` to `python:3.13` to support `apt-get`
- Installed required system libraries (`gcc`, `libpq-dev`) to allow building `psycopg2` from source
- Ensured PostgreSQL is available before the Python app starts (with a 5-second delay)

---

## 🚀 How to Run

```bash
# Start the app and database
docker-compose up --build
```

## ✅ Expected Output

Once both services are running, you should see in the logs:

```
Successfully connected to the database!
```

---

## 🧠 What This Covers

- Using `docker-compose` for multi-container apps
- Dependency order using `depends_on`
- Installing system packages in Docker for Python libs like `psycopg2`
- Service-to-service communication using Docker Compose network

---