# CRM Backend API

This is a **CRM Backend API** written in Go. It provides a RESTful service to manage customer data, including functionalities to list, retrieve, create, update, and delete customers.

## Features

- **List all customers**: Retrieve a list of all customers in JSON format.
- **Get a specific customer**: Retrieve details of a single customer using their ID.
- **Add a new customer**: Create a new customer with a unique ID.
- **Update an existing customer**: Modify the details of an existing customer.
- **Delete a customer**: Remove a customer from the list using their ID.
- **Serve static files**: Optionally serves static files from a `./static` directory.

## Prerequisites

- [Go (version 1.22 or later)](https://golang.org/doc/install)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/crm-backend-api.git
   cd crm-backend-api
   ```

2. **Build the project:**

   ```bash
   go build
   ```

## Running the Application

1. **Run the server:**

   ```bash
   go run main.go
   ```

2. The server will start on port **3000**:

   ```
   Server is starting on port 3000...
   ```

3. **Access the API** using the following base URL:

   ```
   http://localhost:3000
   ```

## API Endpoints

### 1. List All Customers

- **Endpoint**: `/customers`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "role": "Admin",
      "email": "john@example.com",
      "phone": "1234567890",
      "contacted": false
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "role": "User",
      "email": "jane@example.com",
      "phone": "0987654321",
      "contacted": true
    }
  ]
  ```

### 2. Get a Specific Customer by ID

- **Endpoint**: `/customers/{id}`
- **Method**: `GET`
- **Example**: `/customers/1`
- **Response**:
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "role": "Admin",
    "email": "john@example.com",
    "phone": "1234567890",
    "contacted": false
  }
  ```

### 3. Add a New Customer

- **Endpoint**: `/customers`
- **Method**: `POST`
- **Request Body** (JSON):
  ```json
  {
    "name": "Bob Marley",
    "role": "User",
    "email": "bob@example.com",
    "phone": "4445556666",
    "contacted": false
  }
  ```
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "role": "Admin",
      "email": "john@example.com",
      "phone": "1234567890",
      "contacted": false
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "role": "User",
      "email": "jane@example.com",
      "phone": "0987654321",
      "contacted": true
    },
      {
    "id": 4,
    "name": "Bob Marley",
    "role": "User",
    "email": "bob@example.com",
    "phone": "4445556666",
    "contacted": false
  }
  ]
  ```

### 4. Update an Existing Customer

- **Endpoint**: `/customers/{id}`
- **Method**: `PUT`
- **Example**: `/customers/2`
- **Request Body** (JSON):
  ```json
  {
    "id": 2,
    "name": "Jane Doe",
    "role": "User",
    "email": "janedoe@example.com",
    "phone": "7778889999",
    "contacted": true
  }
  ```
- **Response**:
  ```json
  {
    "id": 2,
    "name": "Jane Doe",
    "role": "User",
    "email": "janedoe@example.com",
    "phone": "7778889999",
    "contacted": true
  }
  ```

### 5. Delete a Customer

- **Endpoint**: `/customers/{id}`
- **Method**: `DELETE`
- **Example**: `/customers/3`
- **Response**: `204 No Content`

## Error Handling

The API returns appropriate HTTP status codes for different types of errors, such as:

- `400 Bad Request` for invalid input or malformed JSON.
- `404 Not Found` if the customer ID does not exist.
- `409 Conflict` if trying to add a customer with an existing ID.
- `405 Method Not Allowed` for unsupported HTTP methods.

## Example Usage with `curl`

1. **List all customers**:
   ```bash
   curl -X GET http://localhost:3000/customers
   ```

2. **Get a specific customer**:
   ```bash
   curl -X GET http://localhost:3000/customers/1
   ```

3. **Add a new customer**:
   ```bash
   curl -X POST http://localhost:3000/customers -H "Content-Type: application/json" -d '{"id":4,"name":"Bob Marley","role":"User","email":"bob@example.com","phone":"4445556666","contacted":false}'
   ```

4. **Update a customer**:
   ```bash
   curl -X PUT http://localhost:3000/customers/2 -H "Content-Type: application/json" -d '{"name":"Jane Doe","role":"User","email":"janedoe@example.com","phone":"7778889999","contacted":true}'
   ```

5. **Delete a customer**:
   ```bash
   curl -X DELETE http://localhost:3000/customers/3
   ```

## Project Structure

```
.
├── main.go
├── static
│   └── index.html
└── README.md
```

API documentation is also available at the root endpoint http://localhost:3000

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any improvements or bug fixes.
