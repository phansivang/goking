# User Management API

This is a **User Management API** built with **Golang**, **Gin**, and **GORM**. It follows the **N-tier architecture** and provides CRUD operations for users and user profiles. The API supports filtering, pagination, and standardized response formats.

---

## Features

- **User Management**:
    - Create, read, update, and delete users.
    - Filter users by name, description, and creation date.
    - Paginate user listings.

- **User Profile Management**:
    - Create, read, update, and delete user profiles.

- **Standardized Responses**:
    - All responses follow a consistent format:
      ```json
      {
        "data": [],
        "message": "OK"
      }
      ```

- **Validation**:
    - Input validation using DTOs (Data Transfer Objects).

- **Environment Variables**:
    - Database credentials are loaded from a `.env` file.

---

## Technologies Used

- **Golang**: Backend programming language.
- **Gin**: HTTP web framework.
- **GORM**: ORM for database operations.
- **PostgreSQL**: Relational database.
- **godotenv**: Load environment variables from a `.env` file.