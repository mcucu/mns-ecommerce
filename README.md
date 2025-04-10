# MNSolution E-Commerce

## Overview
This project implements an online ordering system that allows users to select items and place orders. The system includes rules for packaging items based on their price and weight, ensuring efficient delivery.

## Project Structure
```
mns-ecommerce
├──api
│   ├──cmd
│   │   └── root.go                 # Register all api dependencies
│   ├── internal
│   │   ├── handlers                # Handles HTTP requests
│   │   │   └── order_handler.go    # Handles HTTP requests related to orders
│   │   ├── models
│   │   │   └── order.go            # Defines the Order struct and related methods
│   │   ├── services
│   │   │   └── order_service.go    # Logic for packaging orders based on rules
│   │   └── utils
│   │       └── validation.go       # Utility functions for validating order data
│   ├── go.mod                      # Go module file for dependencies
│   ├──main.go                      # Entry point of the application
│   ├──ecommerce.dh                 # sqlite db
│   └── README.md                   # API documentation
├── web
│   ├── assets
│   │   ├── css
│   │   │   └── styles.css      # CSS styles for the web application
│   │   └── js
│   │       └── app.js          # JavaScript code for handling item selection
│   ├── index.html              # Main HTML file for the web application
│   ├── components
│   │       └── order_form.html     # HTML structure for the order form
│   └── README.md                   # API documentation
└── README.md                   # Project documentation
```

## Setup Instructions
### Clone Project
- Clone the repository:
   ```
   git clone <repository-url>
   cd mns-ecommerce
   ```

### Setup dotenv
- create .env file for api service : `cp ./api/.env.example ./api/.env`
- create .env file for web service : `cp ./web/.env.example ./web/.env`

### Run all services
- run docker composer : `docker-compose up --build`

### Run API Project
1. Go to api folder: `cd api/`
2. Install Go dependencies:
   ```
   go mod tidy
   ```
3. Run the application:
   ```
   go run main.go
   ```
4. Open your web browser and navigate to `http://localhost:3000` to access the api.

### Run Web Project
1. Go to web folder: `cd web/`
2. Install js dependencies:
   ```
   npm install
   ```
3. Run the application:
   ```
   npm start
   ```
4. Open your web browser and navigate to `http://localhost:8080` to access the online ordering system.

## Features
- User-friendly interface for item selection and order placement.
- Dynamic packaging of items based on price and weight.
- Validation of order data to ensure compliance with packaging rules.

## Technologies Used
- Go for the backend server and business logic.
- HTML5, CSS, and JavaScript for the frontend interface.
- Bootstrap for responsive design (if included in the assets).

## Demo App
- api service : `https://mns-ecommerce-api.herokuapp.com/`
- web service : `https://mns-ecommerce-web.herokuapp.com/`

## System Diagram
![Alt text](docs/ecommerce.svg?raw=true "System Diagram")
