Here's a basic README for the ReferralAPI project on GitHub. It includes an overview, installation instructions, configuration details, and usage examples to get started with the API.

---

# ReferralAPI

ReferralAPI is a referral link management service built with Golang and Gin. It allows users to create accounts, log in, and manage referrals by generating and validating unique referral codes.

## Features

- **User Registration**: Standard sign-up and referral-based sign-up.
- **Login Authentication**: Secure login with JWT-based token generation.
- **Referral Code Management**: Generate, retrieve, and delete unique referral codes.
- **Documentation**: Swagger documentation for API endpoints.

## Technologies Used

- **Go**: Core language
- **Gin**: Web framework for building RESTful APIs
- **GORM**: ORM for database interaction
- **Swaggo**: Generate and serve Swagger documentation
- **JWT**: For user authentication
- **MySQL**: Relational database for storing user and referral data

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [MySQL](https://www.mysql.com/downloads/)
- [Swagger](https://swagger.io/tools/swagger-ui/) (optional for viewing API docs)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/left-try/ReferalAPI.git
   cd ReferalAPI
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up MySQL database**:

   Create a new MySQL database for the project and configure the database connection in `database/config.go`:

   ```go
   const (
       DBUser     = "your_mysql_username"
       DBPassword = "your_mysql_password"
       DBName     = "referral_db"
   )
   ```

4. **Run database migrations**:

   Run the necessary migrations to set up the database schema.

5. **Generate Swagger Documentation**:

   Install swaggo CLI and generate the Swagger documentation:

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   swag init
   ```

6. **Run the API server**:

   ```bash
   go run main.go
   ```

   The API will start on `localhost:8000`.

## API Documentation

Swagger documentation is available at:

```
http://localhost:8000/docs/index.html
```

### Example API Endpoints

#### 1. **User Signup**

- **Endpoint**: `POST /signup`
- **Description**: Registers a new user.
- **Request Body**:
  ```json
  {
      "email": "user@example.com",
      "password": "userpassword"
  }
  ```
- **Response**:
  ```json
  {
      "message": "User created",
      "user": {
          "id": 1,
          "email": "user@example.com",
          "referrerId": null
      }
  }
  ```

#### 2. **Login**

- **Endpoint**: `POST /login`
- **Description**: Logs in a user and returns a JWT token.
- **Request Body**:
  ```json
  {
      "email": "user@example.com",
      "password": "userpassword"
  }
  ```
- **Response**:
  ```json
  {
      "message": "User logged in",
      "token": "jwt-token-string"
  }
  ```

#### 3. **Create Referral Code**

- **Endpoint**: `POST /create_code`
- **Description**: Generates a new referral code for a user.
- **Request Body**:
  ```json
  {
      "code": "REF123"
  }
  ```
- **Response**:
  ```json
  {
      "code": "REF123"
  }
  ```

#### 4. **Delete Referral Code**

- **Endpoint**: `DELETE /delete_code/:code`
- **Description**: Deletes a specific referral code.
- **Response**:
  ```json
  {
      "message": "Code deleted successfully"
  }
  ```

## Running Tests

1. **Run Server**: Start up the server as shown higher
2. **Run Tests**: Manually run http request located in the `tests` folder