# 🧪 Practice 5 - Small Multi-Language App (Partially Dockerized)

## 📚 About This Practice

This practice is based on a project from [DevOps Directive Docker Course](https://github.com/sidpalas/devops-directive-docker-course/tree/main/05-example-web-application).

It is a simple **3-tier web application** composed of:

- **React frontend**: A web client built with React and Vite.
- **Node.js API**: Backend API developed with Express.js.
- **Golang API**: Backend API developed with Go.
- **PostgreSQL Database**: A relational database.

> In this version, **only the PostgreSQL database** is Dockerized.  
> The frontend and backend services (React, Node.js, and Go) are launched **locally** on the host machine.

---

## 🧩 Architecture

```
[ React Frontend ]
        ↓
[ Node.js API ] <--> [ Golang API ]
        ↓                  ↓
      [ PostgreSQL (Dockerized) ]
```

- **Frontend** (React) communicates with both **Node.js API** and **Golang API**
- **APIs** fetch data from a shared **PostgreSQL** instance running in Docker

---

## ⚙️ Technologies Used

- React (Vite)
- Node.js (Express.js)
- Golang
- PostgreSQL (Dockerized)
- Makefile (for running services)

---

## 🏃‍♂️ How to Run the Project (Practice 5)

### 1. Start PostgreSQL in Docker

Use the Makefile command to launch a PostgreSQL container:

```bash
make run-postgres
```

This sets up the PostgreSQL database with the following values:
- `POSTGRES_PASSWORD=foobarbaz`
- Volume: `pgdata:/var/lib/postgresql/data`
- Port: `5432`

### 2. Run the Node.js API

```bash
make run-api-node
```

This will:
- Navigate to the `api-node` directory
- Start the development server with the correct `DATABASE_URL`
- Run this command if you have error with `nvm` or `npm`:
```bash
nvm ls
nvm use node 19.4
npm install
npm run dev
```
To install the dependencies and start the Node.js API server.

### 3. Run the Golang API

```bash
make run-api-golang
```

This will:
- Navigate to the `api-golang` directory
- Run the Go API server
- Run this command if you have error with `go run`:
```bash
mkdir go-workspace
export GOPATH=$PWD/go-workspace
go mod download
go run main.go
```
To download the dependencies and start the Golang API server.

### 4. Run the React Frontend

```bash
make run-client-react
```

This will:
- Navigate to the `client-react` directory
- Start the Vite development server

---

## 🌍 Access the Application

Once all services are running, open your browser:

```
http://localhost:1515
```

The React frontend will:
- Fetch data from both Node.js and Golang APIs
- Both APIs will fetch their data from the shared PostgreSQL container

---

## 📂 File Highlights

Here’s a preview of key Makefile commands:

```makefile
DATABASE_URL:=postgres://postgres:foobarbaz@localhost:5432/postgres

run-postgres:
	docker run -e POSTGRES_PASSWORD=foobarbaz \
	-v pgdata:/var/lib/postgresql/data \
	-p 5432:5432 postgres:15.1-alpine

run-api-node:
	cd api-node && DATABASE_URL=${DATABASE_URL} npm run dev

run-api-golang:
	cd api-golang && DATABASE_URL=${DATABASE_URL} go run main.go

run-client-react:
	cd client-react && npm run dev
```

---

## 🧠 What This Practice Covers

- Multi-language project structure and collaboration between services
- Running only the database in Docker while developing other services locally
- Managing and reusing environment variables across services
- Using a `Makefile` to simplify development workflow

---