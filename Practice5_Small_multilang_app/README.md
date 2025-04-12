# Practice 5 - Small Multi-Language App (Partially Dockerized)

## About This Practice

This practice is based on a project from [DevOps Directive Docker Course](https://github.com/sidpalas/devops-directive-docker-course/tree/main/05-example-web-application).

It is a simple **3-tier web application** composed of:

- **React frontend**: A web client built with React and Vite.
- **Node.js API**: Backend API developed with Express.js.
- **Golang API**: Backend API developed with Go.
- **PostgreSQL Database**: A relational database.

In this version, **only the PostgreSQL database** was Dockerized.  
The frontend and backend services (React, Node.js, and Go) were launched **locally** on the host machine.

---

## Architecture

- **Frontend** (React) ➔ communicates with ➔ **Node.js API** and **Golang API**
- **APIs** ➔ fetch data from ➔ **PostgreSQL** (running inside Docker)

---

## Technologies Used

- React
- Node.js
- Golang
- PostgreSQL (Dockerized)

---

## How to Run the Project (Practice 5)

1. **Start PostgreSQL in Docker:**
   Follow Makefile instructions to run PostgreSQL in a Docker container. This will set up the database for the APIs to connect to.

2. **Run Node.js API:**
   - Navigate to the `api-node` directory.
   - Install dependencies using `npm install`.
   - Start the API using `make run-api-node`.

3. **Run Golang API:**
   - Navigate to the `api-golang` directory.
   - Install dependencies using `go mod download`.
   - Start the API using `make run-api-golang`.

4. **Run React Frontend:**
   - Navigate to the `client-react` directory.
   - Install dependencies using `npm install`.
   - Start the frontend using `make run-client-react`.

5. **Access the Application:**
   - Open your web browser and navigate to `http://localhost:3000` to access the React frontend.
    - The frontend will communicate with the Node.js and Golang APIs to fetch data from the PostgreSQL database.