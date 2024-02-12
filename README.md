# Online Store MVP

## Project Description

This project is a backend API for an online store MVP ordering system created using the Go language.


## Key Technologies

1. **Go**: Used as the main programming language to develop the backend API. Golang was chosen for its high performance, ease of code management, and support for large-scale application development.

2. **Fiber**: A lightweight and fast web framework for Golang. Gofiber is used to build API endpoints with optimal performance.

3. **Postgres**: Postgres relational database is used to store and manage data related to ticket booking. This provides reliability and flexibility in data management.

4. **JWT**: Used for user authentication and authorization. JWT provides a secure way to transmit authentication information between parties involved.

5. **Midtrans**: Engaged to handle the payment process. Midtrans is an integrated payment gateway, allowing the application to accept payments with various payment methods.



## Running the Project

1. Make sure Golang is installed on your system.
2. configuration settings for the database and others according to the .env.example file
3. Install dependencies using the `go mod tidy` command.
4. Adjust the Postgres database and Midtrans API configuration in the configuration file.
5. Run the application with the `go run main.go` command.
6. The backend API will run on localhost as you set in the .env file.


## Documentation

1. https://dbdiagram.io/d/online-store-mvp-db-65c3907dac844320aea87fde
2. https://app.swaggerhub.com/apis/A7520971/online-store-mvp/1.0.0
3. https://documenter.getpostman.com/view/32137512/2sA2r3ZRFj


