# Web-application template

This repository contains a basic template for a web application built with Go (backend) and Vue.js (frontend).
The goal is to provide a starting point for future projects with common features, such as user authentication,
database connectivity with PostgreSQL, API data fetching, and a modern, responsive frontend using Bootstrap.

The project is designed to be containerized using Docker,
ensuring easy deployment and consistent development environments.

## Purpose

I created this project to explore Go environment and enhance my Vue.js skills.
The idea was to create a simple, scalable web application that could be reused for various other projects in the future.

## Features

* **Auth**: registration and login functionality,
* **PostgreSQL integration**: database for persistance,
* **Basic API**: this includes routes for user management and any other custom data operations,
* **Modular structure**: the backend code is modular and tries to follow good Go practices, making it easy to extend with additional routes and services.
* **Responsive design**: Vue.js + Bootstrap 5, giving it a simple but modern look that works well on different devices,

## Setup

### Requirements

* docker and docker-compose
* go
* nodejs
* vue-cli

### Steps

1. Clone this repository
2. Setup required variables (and ensure that read writes to this file are set correctly)

```
# docker-compose.yml
DB_HOST=localhost
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
DB_PORT=5432
PORT=8080
JWT_SECRET=yourjwtsecret
JWT_SECRET=your_jwt_secret
```

3. Start the application

```
docker-compose up --build
```

4. Access the application

* Backend: `http://localhost:8080/api/v1/`
* Frontend: `http://localhost:80/`

### Tests

Just run `go test ./... -v` from backend directory. Firstly you may need to install go dependencies and libraries via `go mod download`.

### Development

You may want to develop on the host machine just to make things faster and easier:

* `npm run serve --live` - live reload for frontend.
* `go build -o main . && ./main` - as far as I know there is no native way to do live reload but I think it is not necessary. One may use some library for this (air?)
