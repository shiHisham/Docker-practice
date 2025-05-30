# 🐳 Practice 1: Hello Docker World

This is a basic introduction to Docker using a simple Python app.

The goal is to understand how to:
- Write a Dockerfile
- Build a Docker image
- Run a container

## 📦 What’s Inside

- A Python script that prints a message
- A Dockerfile based on the official `python:alpine` image

## 🚀 How to Run

```bash
# Step 1: Build the Docker image
docker build -t docker-hello-world .

# Step 2: Run the container
docker run --rm docker-hello-world
```

## ✅ Output

```bash
Hello from inside a docker container !
```

## 🧠 What This Covers

- Basic Dockerfile instructions: `FROM`, `WORKDIR`, `COPY`, `CMD`
- How to build an image with `docker build`
- How to run a container with `docker run`

---


---

🧾 *README content written based on my own implementation and learning. Formatting and structure assisted by ChatGPT for clarity and presentation.*
