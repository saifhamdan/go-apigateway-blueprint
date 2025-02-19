# GO apigateway blueprint by Saif Hamdan

This project is a Go RESTful API dockerized gateway blueprint using my favorite web framework Fiber, along with GORM, PostgreSQL, Redis, and NATS. 

**Note:** This blueprint is for demonstration purposes and is not production-ready. The main purpose of this blueprint is to serve as a foundation for your next projects and to provide some inspiration.

I was inspired by the official Go [article](https://go.dev/doc/modules/layout).

## Project Overview

This project demonstrates a modular and scalable layout for a Go-based API gateway. It includes various components and configurations to help you get started quickly.

### Key Features

- **Fiber:** My favourite web framework in Go.
- **GORM:** An ORM library for Go, providing a powerful database abstraction layer.
- **PostgreSQL:** A robust, open-source relational database.
- **Redis:** An in-memory data structure store, used as a database, cache, and message broker.
- **NATS:** A simple, secure, and high-performance messaging system for cloud-native applications.

## Project Layout

The project is organized into the following directories:

1. **cmd:** Contains the main application entry points.
2. **config:** Configuration files and settings.
3. **data:** Data-related files, such as database migrations.
4. **internal:** Internal application code, including business logic and services.
5. **logs:** Log files generated by the application.
6. **models:** Data models and schemas.
7. **pkg:** Reusable packages and utilities.
8. **public:** Publicly accessible files, such as static assets.
9. **swagger:** API documentation generated by Swagger.

## Getting Started

To get started with this project, follow these steps:

1. **Clone the repository:**
    ```sh
    git clone https://github.com/saifhamdan/go-apigateway-blueprint.git
    cd go-apigateway-blueprint
    ```

2. **Install dependencies:**
    ```sh
    make init
    ```

3. **Create .env file:**
    ```sh
    nano .env
    ```
    then add your custom environment variables

4. **run docker compsoe stack for development:**
    ```sh
    make stackup
    ```    

5. **Run the application:**
    ```sh
    make run
    ```


## Seed Data
To populate the database with initial seed data, you can use the following command:

  ```sh
  make seed
  ```

This command will execute the seeding script defined in the Makefile. The seeding script typically inserts predefined data into the database, which can be useful for development and testing purposes. Ensure that your database is properly configured and running before executing this command.

## Generate Swagger Docs
To generate Swagger documentation for your API, you can use the following command:

  ```sh
  make swag
  ```

This command will execute the Swagger generation tool, which scans your Go code for annotations and generates the corresponding Swagger documentation. The generated documentation can be used to visualize and interact with your API using tools like Swagger UI.