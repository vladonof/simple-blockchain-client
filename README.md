# Blockchain Client Home Task

## Overview

This is a simple blockchain client API. The API provides endpoints to interact with blockchain data, including retrieving the current block number and fetching details of a block by its number.

## Endpoints

### 1. Get Block Number

- **URL:** `/blockNumber`
- **Method:** `POST`
- **Description:** Retrieves the block number.

### 2. Get Block By Number

- **URL:** `/blockByNumber`
- **Method:** `POST`
- **Description:** Retrieves the details of a block by its number.
- **Request Body:**

  ```json
  {
    "blockNumber": "0x134e82a"
  }
  ```

## How to Start the Server

### Start Locally

3. **Download Dependencies:**

   ```sh
   go mod download
   ```

4. **Run the Server:**

   ```sh
   go run cmd/server/main.go
   ```

### Start with Docker

1. **Build the Docker Image:**

   ```sh
   docker build -t simple-blockchain-app .
   ```

2. **Run the Docker Container:**

   ```sh
   docker run -p 8080:8080 simple-blockchain-app
   ```

## How to Run Tests

```sh
go test ./internal/...
```

## What Should Be Added for the API to be "Prod Ready"

- Implement structured logging to capture detailed logs for monitoring and debugging
- Use a configuration management tool or environment variables to manage different configurations/secrets for different environments
- Improve error handling to provide more descriptive error messages and proper HTTP status codes
- Implement authentication and authorization to secure the API endpoints
- Ensure all inputs are properly validated.
- Add rate limiting to prevent abuse of the API
- CI/CD Pipeline
  - Run Tests
  - Static Code Analysis
  - Dependency scan
  - Build and push docker image to ECR
- Set up monitoring(both log ingesting and infrastructure data) and alerting to track the performance and availability.
- Review the HCL, make it more dynamic and configurable, stuff like:
  - Proper VPC/subnetting
  - Security Groups
  - SSL
  - IAM
  - Autoscaling
- Implement caching layers to improve performance (if applicable)
- Provide comprehensive API documentation
- Comprehensive load test
