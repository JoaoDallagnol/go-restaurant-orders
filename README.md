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

### Menu Service
  The `menu-service` manages restaurants and their dishes, ensuring a clean relationship between entities (1:N).  

  - 📚 **API Endpoints** (following **OpenAPI 3.0.3**):
    - **Restaurants**
      - `POST /restaurants` – Create a new restaurant  
      - `GET /restaurants` – List all restaurants *(with dishes)*  
      - `GET /restaurants/{id}` – Get a restaurant by ID *(with dishes)*  
      - `PUT /restaurants/{id}` – Update a restaurant by ID  
      - `DELETE /restaurants/{id}` – Delete a restaurant by ID  
    - **Dishes**
      - `POST /dishes` – Create a new dish (linked to a restaurant)  
      - `GET /dishes` – List all dishes  
      - `GET /dishes/{id}` – Get a dish by ID *(with restaurant ID)*  
      - `PUT /dishes/{id}` – Update a dish by ID  
      - `DELETE /dishes/{id}` – Delete a dish by ID  

  - 🛠️ **Technologies Used**:
    - **Go 1.24.4** – Core programming language  
    - **Gin** – Lightweight HTTP web framework  
    - **GORM** – ORM for database interaction  
      - `gorm.io/gorm`  
      - `gorm.io/driver/postgres`  
    - **Swagger / OpenAPI 3.0** – API documentation and client generation  
    - **Viper** – Configuration management  
    - **Go Modules** – Dependency management  
    - **shopspring/decimal** – High-precision decimal for dish prices  

### Order Service
  The `order-service` handles the creation, management, and tracking of customer orders, integrating user and menu data to manage the entire order lifecycle.

  - 📚 **API Endpoints** (following **OpenAPI 3.0.3**):
    - **Orders**
      - `POST /orders` – Create a new order  
      - `GET /orders` – List all orders  
      - `GET /orders/{id}` – Get order by ID  
      - `PUT /orders/{id}` – Update an existing order  
      - `DELETE /orders/{id}` – Delete an order  
    - **OrderItems**
      - `GET /order-items` – List all order items  
      - `GET /order-items/{id}` – Get order item by ID

  - 🛠️ **Technologies Used**:
    - **Go 1.24.6** – Core programming language  
    - **Gin** – Lightweight HTTP web framework  
    - **GORM** – ORM for database interaction  
      - `gorm.io/gorm`  
      - `gorm.io/driver/postgres`  
    - **Swagger / OpenAPI 3.0** – API documentation and client generation  
    - **Viper** – Configuration management  
    - **Go Modules** – Dependency management  
    - **shopspring/decimal** – High-precision decimal for monetary values  