# gRPC Server

This branch hosts a gRPC server implemented in Go, providing a robust foundation for handling remote procedure calls efficiently.

## Overview

The gRPC server in this branch establishes a robust infrastructure to manage a `PingService` that handles ping requests. This service allows clients to interact with the server and receive responses.

## Setup

To run and test this gRPC server, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/MrAzharuddin/grpc_tutorial.git
   ```
2. **Navigate to the server directory:**

   ```bash
   cd grpc_tutorial
   ```

3. **Install dependencies:**

   Make sure you have Go installed and then execute:

   ```bash
   go mod tidy
   ```

4. **Run the server:**

   Execute the following command to start the server:

   ```bash
   go run main.go
   ```

## Dependencies

- **Go:** Ensure that Go is installed to compile and execute the server.
- **gRPC (google.golang.org/grpc):** Utilizes the gRPC package for managing remote procedure calls efficiently.

## Usage

The gRPC server exposes the `PingService` for handling ping requests. By sending a ping, gRPC clients can receive responses from the server.

### gRPC Services

- **PingService**: Manages ping requests and responses.

## Server Configuration

The server configuration is specified in the `main.go` file:

## Contributing

Contributions to this project are welcome. To contribute, please follow these guidelines:

- Fork the repository.
- Create a new branch for your feature/bug fix.
- Make changes and ensure tests pass.
- Submit a pull request to the `grpc_server` branch of the repository.

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](./CODE_OF_CONDUCT.md). We welcome and encourage everyone to participate with mutual respect.

## License

This project is licensed under the MIT License. For more information, see the [LICENSE](./LICENSE) file.

We encourage you to contribute, open issues, and suggest improvements to this gRPC server.
