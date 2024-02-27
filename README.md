[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/LECuYE4o)
# Bookstore API

## Overview

This is a RESTful API for a bookstore, built using Go, Docker, PostgreSQL, and Nginx. The API supports operations like creating users, adding books to inventory, and writing reviews.

## Features

- User authentication
- Inventory management
- Writing and viewing book reviews
- Admin functionalities
- HTTPS support via Nginx
- Reverse proxy using Nginx

## Requirements

- Docker
- Docker Compose
- Go 1.21.0 or higher

## Installation

1. Clone the repository
    ```bash
    git clone https://github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-35C4n0r
    ```

2. Navigate to the project folder
    ```bash
    cd vit-2025-summer-engineering-internship-task-35C4n0r
    ```

3. HTTPS Configuration for Local Development

    This project uses HTTPS for secure communication. For local development, you can generate a self-signed SSL certificate using OpenSSL. Run the following command:
   Place them in a folder ```certs``` in the root directory of the project.
    ```bash
    openssl req -x509 -newkey rsa:4096 -keyout localhost.key -out localhost.crt -days 365 -nodes -subj "/CN=localhost"
    ```
   **I've already added the certificates in certs/\* for ease of use.**

4. Environment Variables

    The project uses the following environment variables:
    - `DB_HOST`: Database host name
    - `DB_USER`: Database user
    - `DB_PASSWORD`: Database password
    - `SERVER_PORT`: Port for the Go application
    - `SERVER_SECRET`: Secret for JWT generation
    - `ADMIN_SECRET`: Secret for admin functionalities
    
    **I've already added all of the above variables to the ```docker-compose.yml``` for ease of use**

5. Build and start the Docker containers
    ```bash
    docker-compose build
    docker compose up
    ```

## Usage

Once the containers are up, the API will be available at `https://localhost/api/`.

## API Endpoints

### Public Routes

- `GET /`: Sanity check to test the server.
- `POST /register`: Register a new user.
- `POST /login`: Login and receive a JWT token.

### Protected Routes

These routes require a JWT token in the `Authorization` header.

#### General Routes

- `GET /middleware`: Checks if middleware is functioning correctly.
- `POST /deactivate`: Deactivate a user's account.
- `DELETE /delete`: Delete a user's account.

#### Book Routes

- `GET /books`: Search for books.
- `POST /cart`: Add a book to the cart.
- `PUT /cart`: Update the quantity of an item in the cart.
- `DELETE /cart/:isbn`: Remove a book from the cart.
- `GET /cart`: Retrieve items in the cart.
- `POST /purchase`: Purchase items in the cart.

#### Review Routes

- `POST /review`: Add a review for a book.
- `GET /review`: Retrieve reviews for a book.

#### Download Route

- `GET /download/:isbn`: Download a book.

### Admin Routes

These routes require admin privileges.

- `POST /admin/add`: Add a new book to the database.
- `DELETE /admin/delete/:isbn`: Remove a book from the inventory.
- `PUT /admin/update`: Update the details of a book.




