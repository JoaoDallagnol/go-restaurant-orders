# go-restaurant-orders

**`go-restaurant-orders`** is a microservices-based project built with **Go (Golang)** to efficiently manage restaurant ordering systems. It is structured into several independent services, each with its own clear responsibilities, enabling scalability and modularity.

### 🧩 Architecture Overview

The system is composed of the following microservices:

- **Auth Service** – Handles authentication and user management
- **Menu Service** – Manages menu items and categories
- **Order Service** – Handles order creation and tracking
- **Payment Service** – Manages payment processing and status

### Auth Service 
  The `auth-service` provides endpoints for user authentication and JWT-based access control.

  - 📚 **API Endpoints** (following **OpenAPI 3.0.3**):
    - `POST /register` – Register a new user  
    - `POST /login` – Authenticate and receive a JWT  
    - `GET /users` – List all users *(JWT protected)*  
    - `GET /users/{id}` – Get user by ID *(JWT protected)*  
    - `PUT /users/{id}` – Update user by ID *(JWT protected)*  
    - `DELETE /users/{id}` – Delete user by ID *(JWT protected)*  

  - 🛠️ **Technologies Used**:
    - **Go 1.24.4** – Core programming language  
    - **Gin** – Lightweight HTTP web framework  
    - **GORM** – ORM for database interaction  
      - `gorm.io/gorm`  
      - `gorm.io/driver/postgres`  
    - **JWT** – Authentication via [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)  
    - **Swagger / OpenAPI 3.0** – API documentation and client generation  
    - **Viper** – Configuration management  
    - **Go Modules** – Dependency management  
