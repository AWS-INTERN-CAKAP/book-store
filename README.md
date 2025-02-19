# Book Store App

# Introduction
Book Store App is a full-stack application consisting of a frontend built with React.js and a backend developed in Golang. The application is containerized using Docker and can be easily deployed using Docker Compose.

# Tech Stack
- Frontend: React.js

- Backend: Golang (echo)

- Database: MySQL

- Containerization: Docker, Docker Compose

# Prerequisites
Before running the application, ensure you have the following installed on your system:

- [Docker](https://www.docker.com/)

- [Docker Compose](https://docs.docker.com/compose/install/)

# Setup & Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd book-store
```

2. Create a .env file (optional, if needed)
If your application requires environment variables, copy the env.example file in both the backend and frontend directories and rename it to .env

```bash
cp backend/env.example backend/.env
cp frontend/env.example frontend/.env
```

3. Run the application using Docker Compose:

```bash
docker-compose up --build -d
```

4. Access the application:

```bash
Frontend: http://localhost:3000
Backend API: http://localhost:8080
```
