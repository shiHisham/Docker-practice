# 🗄️ Practice 4: Docker Volumes for Persistent PostgreSQL Database

This practice demonstrates how Docker volumes allow PostgreSQL data to persist even after containers are stopped or removed.

---

## 🎯 Goal

- Understand how to persist database data using Docker volumes.
- Prove that data is not lost when containers are recreated.

---

## 📦 What’s Inside

- A `docker-compose.yml` file defining a PostgreSQL service with a named volume (`postgres-data`).

---

## 🛠️ What Was Done

1. Launched PostgreSQL with a named volume:
   ```yaml
   volumes:
     - postgres-data:/var/lib/postgresql/data
   ```
2. Connected to the running container using `psql`:
   ```bash
   docker exec -it <container_id> psql -d docker_db -U docker_user
   ```
3. Created a table and inserted a record:
   ```sql
   CREATE TABLE cars (
     brand VARCHAR(255),
     model VARCHAR(255),
     year INT
   );

   INSERT INTO cars (brand, model, year) VALUES ('Ford', 'Kuga', 2018);

   SELECT * FROM cars;
   ```
4. Verified the record exists.
5. Stopped and removed containers:
   ```bash
   docker-compose down
   ```
6. Relaunched the stack:
   ```bash
   docker-compose up
   ```
7. Reconnected and confirmed that the record still exists.

---

## ✅ Result

The inserted data was successfully retained across container restarts because it was stored in the Docker volume `postgres-data`.

---

## 🧠 What This Covers

- Docker named volumes for persistent storage
- PostgreSQL data durability in containerized environments
- Safe teardown and re-creation of stateful services

---

---

🧾 *README content written based on my own implementation and learning. Formatting and structure assisted by ChatGPT for clarity and presentation.*
