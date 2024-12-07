# Go HTTP Server

This is a simple HTTP server written in Go that responds with "hello, world" when accessed.

## Prerequisites

- Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

## Setup

1. Clone the repository or create a new directory for the project:

   ```bash
   mkdir go-http-server
   cd go-http-server
   ```

2. Initialize the Go module:

   ```bash
   go mod init go-http-server
   ```

3. Create a file named `main.go` and add the server code.

## Running the Server

1. Run the server using the following command:

   ```bash
   go run main.go
   ```

2. Open a web browser and go to `http://localhost:3000` to see the "hello, world" message.

## Code Overview

- The server listens on port 3000 by default.
- The `events` function handles requests to the `/events` URL and responds with an event stream.

## Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t go-http-server .
   ```

2. Run the Docker container:

   ```bash
   # Start the service
   docker-compose up -d
   ```

3. View logs

   ```bash
   docker-compose logs -f
   ```

4. Check container status

   ```bash
   docker-compose ps
   ```

5. Restart

   ```bash
   docker-compose restart

   ```

6. Update and rebuild

   ```bash
   docker-compose up -d --build
   ```
