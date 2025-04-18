# API Gateway

This is the API Gateway for the e-commerce microservices platform.

## Features

- Authentication (JWT)
- Request routing to Inventory and Order services
- Request logging
- Basic telemetry

## Environment Variables

- `JWT_SECRET`: Secret key for JWT tokens (default: "very-secret-key")
- `INVENTORY_SERVICE_URL`: URL of Inventory Service (default: "http://localhost:8081")
- `ORDER_SERVICE_URL`: URL of Order Service (default: "http://localhost:8082")
- `PORT`: Port to run on (default: 8080)

## Endpoints

### Auth
- `POST /auth/login` - Get JWT token
- `POST /auth/register` - Register new user

### Inventory (protected)
- `POST /products` - Create product
- `GET /products` - List products
- `GET /products/:id` - Get product
- `PATCH /products/:id` - Update product
- `DELETE /products/:id` - Delete product

### Orders (protected)
- `POST /orders` - Create order
- `GET /orders` - List orders
- `GET /orders/:id` - Get order
- `PATCH /orders/:id` - Update order

### Users (protected)
- `POST /users/register` - Create user
- `GET /users` - List user
- `GET /users/:id` - Get user