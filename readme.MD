# 🧃 Products API

This project is a products API built with Go, following a clean architecture approach, with gRPC communication between microservices and JWT-based authentication.

## 🚀 Technologies Used

- **Go (Golang)** — main language of the application
- **Gin** — lightweight and fast web framework
- **gRPC** — efficient communication between microservices
- **Protocol Buffers (Protobuf)** — defines gRPC contracts
- **PostgreSQL** — relational database
- **Prisma** — ORM used in the user microservice
- **Docker (optional)** — for running dependencies (DB, etc.)
- **JWT (JSON Web Tokens)** — for service-to-service authentication

## 🧩 Microservices Structure

- **users-api** (another service):

  - Responsible for user creation and authentication
  - Generates and signs JWT tokens
  - Exposes a gRPC service (`UserService`) that returns user data
  - Link to repo: https://github.com/Natan-Lucena/users-sharding-api

- **Files-API** (another service):

  - Responsible saving users profile pictures in AWS S3
  - Send messages to a RabbitMQ Queue that is listened by users-api
  - Link to repo:https://github.com/Natan-Lucena/Rabbit-Image-Uploader

- **products-api** (this project):

  - Responsible for product creation and management
  - Validates JWTs received in REST requests
  - Uses gRPC to communicate with the users service
  - Creates products linked to the authenticated user

## ▶️ Running the Project

To run the **products-api**, follow these steps:

1. **Start the `users-api` service**

   Make sure the `users-api` service is up and running (along with its PostgreSQL shards). Inside its project directory, run:

   ```bash
    $ docker-compose up -d

    $ go mod download


    $ go run cmd/rest/main.go

   ```
