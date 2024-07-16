## Task-Management

## Overview
This is a project that securely manages user tasks with features for authentication, task creation, updates, and retrieval.

## Features

- User authentication and authorization
- Create and update tasks to a specific user
- Retrieve tasks by user ID.


## Technologies Used
- Go
- postgres


## API Endpoint
```sh
POST /auth/signup
POST /auth/login

GET /api/users
GET /api/users/{id}

POST /api/tasks/{id}
POST /api/tasks/update/{id}
```

## API Endpoint validation
- You can use postman to test endpoints.


## Prerequisites
- Before you begin, ensure you have Go installed
- postgres installed
- psql SQL shell runing
- Ensure you have your connection details configured in your environment.

## Installation
- Clone this repository to your local machine:
```sh
git clone https://github.com/Udehlee/Task-Management.git
```
- Navigate to the project directory:
- cd Task-Management

- Run the server:
```sh
go run main.go
```
- The server will be listening on port 8080.

