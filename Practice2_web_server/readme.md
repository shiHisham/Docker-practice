# 🌐 Practice 2: Dockerized Web Server with Flask

In this practice, we run a simple web application inside a Docker container using the Flask framework.

The goal is to understand how to:
- Dockerize a basic Python web server
- Use `EXPOSE` to define accessible ports
- Access a containerized web service from the browser

## 📦 What’s Inside

- A small Flask web app responding to `/`
- A Dockerfile that installs Flask and runs the app on port `1516`

## 🚀 How to Run

```bash
# Step 1: Build the Docker image
docker build -t flask-web-app .

# Step 2: Run the container and expose port 1516
docker run --rm -p 1515:1516 flask-web-app
```

## 🌍 Open in Browser

Go to [http://localhost:1515](http://localhost:1515)

You should see:

```
Hello from Dockerized Web Server! using flask Framework 👋
```

## 🧠 What This Covers

- Running web servers in containers
- Installing Python packages with `pip` during image build
- Exposing ports from containers using `EXPOSE` and `-p`
- Serving HTTP content from a Flask app

---