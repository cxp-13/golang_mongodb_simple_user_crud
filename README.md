
# User Management API

This is a simple RESTful API for managing users, built with Go, Gin framework, and MongoDB.

## Setup

### Prerequisites

- Go 1.18 or higher
- MongoDB

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/user-management-api.git
    cd user-management-api
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

### Configuration

Ensure MongoDB is running on `localhost:27017`. You can configure the MongoDB URI in the `getMongoDBClient` function in `main.go` if needed.

## Running the Application

To run the application, execute the following command:
```sh
go run main.go
```

The server will start on `http://localhost:9000`.

## API Endpoints

- **POST** `/user` - Create a new user
- **GET** `/user/:id` - Retrieve a user by ID
- **PUT** `/user/:id` - Update a user by ID
- **DELETE** `/user/:id` - Delete a user by ID

## Example Request

### Create a User

Using curl:
```sh
curl -X POST http://localhost:9000/user \
-H "Content-Type: application/json" \
-d '{
  "name": "Alice",
  "gender": "Female",
  "age": 30
}'
```

## License

This project is licensed under the MIT License.
