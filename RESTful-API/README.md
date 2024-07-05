# Go RESTful API Example

## Description

This project is an example implementation of a RESTful API backend in Go using the Gin framework. It demonstrates CRUD operations (Create, Read, Update, Delete) for managing user resources.

## Features

- Create a new user
- Retrieve all users
- Retrieve a specific user by ID
- Update an existing user
- Delete a user by ID

## Project Structure

The project is structured into separate files for each functionality:

- `main.go`: Initializes the server and sets up the Gin router.
- `create.go`: Contains the handler function to create a new user.
- `read.go`: Contains the handler functions to retrieve all users and retrieve a specific user by ID.
- `delete.go`: Contains the handler function to delete a user by ID.

## Setup

1. Clone the project:

    ```bash
    git clone https://github.com/yourusername/your-project.git
    ```

2. Initialize Database:
   
   - Create a new database. Example database name: `users_db`.
   - Run the provided SQL script to create the necessary tables (`users` table).
   
3. Change configuration:

    - Create a new file named `config.yaml` inside the `/config` folder.
    - Use `config.yaml.example` as a template or see the sample config below:

    ```yaml
    app:
        name: go-restful-api

    server:
        host: localhost
        port: 8080

    database:
        db_host: localhost
        db_port: 5432
        db_name: users_db
        db_username: your_username
        db_password: your_password

    jwt:
        expired: 3600
        issuer: go-restful-api
        secret: your_secret_key
    ```

4. Run the project:

    ```bash
    go run main.go
    ```

## API Endpoints

- **POST** `/users`: Create a new user
- **GET** `/users`: Retrieve all users
- **GET** `/users/:id`: Retrieve a specific user by ID
- **DELETE** `/users/:id`: Delete a user by ID

## Dependencies

- Gin: HTTP web framework
- Gorm: ORM library for database access
- Other libraries listed in `go.mod`

## Postman Collection

- Import `go-restful-api.postman_collection.json` to your Postman to access the API endpoints.


