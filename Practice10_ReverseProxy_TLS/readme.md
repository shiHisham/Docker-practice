# 🛡️ Practice 10 - Reverse Proxy with SSL (Nginx)

---

## 📚 Overview

This practice builds on top of [Practice 7](../Practice7_Small_multilang_app_Dockerized_MultiStage), where each container (frontend, backend) handled its own routing. In Practice 10, we introduced a **dedicated reverse proxy layer using Nginx** to handle:

- 🔒 SSL termination (HTTPS support with self-signed certs)
- 🔁 Centralized request routing from client to backend services
- 🚪 A single secure entrypoint to serve frontend + backend APIs

> This architecture mirrors real-world cloud deployments using ALBs, API gateways, or ingress controllers.

---

## 🧩 Updated Architecture

```
Browser
  ↓ HTTPS (443)
Reverse Proxy (Nginx w/ SSL)
  ├── /           → React frontend (api-react)
  ├── /api-node/  → Node.js API (api-node)
  └── /api-golang/→ Golang API (api-golang)
```

### 🔁 Difference from Practice 7

| Feature                         | Practice 7                              | Practice 10                            |
|--------------------------------|------------------------------------------|-----------------------------------------|
| Reverse Proxy                  | Embedded in `client-react` container     | Dedicated Nginx container               |
| HTTPS Support                  | ❌ None                                  | ✅ Self-signed certs with TLS           |
| Static file serving            | Via Nginx inside React container         | Proxied from reverse proxy             |
| Proxy Logic Location           | Inside client-react (dev & prod)         | In `reverse-proxy` container only      |
| Production-readiness           | 🚧 Partial                               | ✅ Clean separation + HTTPS             |

---

## 🔐 SSL Certificate Setup

Generate a **self-signed certificate** using OpenSSL:

```bash
openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout server.key \
  -out server.crt \
  -subj "/CN=localhost"
```

> This cert is stored inside `nginx/ssl/` and mounted read-only (ro) into the reverse proxy container.

📁 File structure:
```
Practice10_ReverseProxy_TLS/
├── nginx/
│   ├── nginx.conf
│   └── ssl/
│       ├── server.crt
│       └── server.key
├── client-react/
├── api-node/
├── api-golang/
...
```

---

## ⚙️ nginx.conf (Reverse Proxy)

```nginx
server {
  listen 443 ssl;
  server_name localhost;

  ssl_certificate     /etc/nginx/ssl/server.crt;
  ssl_certificate_key /etc/nginx/ssl/server.key;

  location / {
    proxy_pass http://api-react:1516/;
    proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
  }

  location /api-node/ {
    proxy_pass http://api-node:3000/;
    proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
  }

  location /api-golang/ {
    proxy_pass http://api-golang:8080/;
    proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
  }
}
```

> React, Node, and Go run in separate containers. The reverse proxy routes to them by internal Docker DNS names.

---

## 🐳 Docker Compose Update (prod only)

```yaml
  reverse-proxy:
    image: nginx:1.25-alpine
    container_name: reverse-proxy
    ports:
      - "443:443"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/nginx.conf:ro
      - ./nginx/ssl:/etc/nginx/ssl:ro
    depends_on:
      - api-node
      - api-golang
      - api-react
    networks:
      - MSSmall_app
```

---

## 🧪 How to Test

### 1. Start Production Environment
```bash
make up-prod
```

### 2. Visit the app in your browser:
```
https://localhost
```
> You’ll see a browser warning ⚠️ because of the self-signed certificate — accept it to continue.

### 3. Confirm API Routes:
- `https://localhost/api-node/ping`
- `https://localhost/api-golang/ping`

### 4. Or use curl:
```bash
curl -k https://localhost/api-node/ping
```

---

## 🧠 What learned on this practice

- How to create and mount SSL certs
- How to set up Nginx as a standalone reverse proxy
- How to route React, Node, and Go services through a single HTTPS entrypoint
- Difference between frontend-embedded proxy vs centralized proxy layer

---

## 📌 Takeaways

- In Practice 7, the `client-react` container handled proxying internally.
- In Practice 10, we separated proxy logic into a dedicated container with HTTPS —.

> This paves the way for adding rate-limiting, caching, auth headers, or moving to Traefik / Ingress in Kubernetes later.

---

🧾 *README content written based on personal implementation. Formatted and structured with ChatGPT for clarity.*