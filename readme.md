# Client Manager Web API App Using GO

This project is a Go-based web API that follows a specific structure. It consists of packages for domain logic, services, and repositories, along with a main application in the `cmd` directory.

## Project Structure

The project structure is organized as follows:

- `cmd/`: Contains the main application code and configuration files.

  - `go.mod`: Go module file for the main application.
  - `go.sum`: Go module checksum file.
  - `main.go`: Entry point of your web API.

- `pkg/`: Contains various packages for different parts of your application.

  - `domain/`: Contains domain logic and models.
  - `service/`: Contains services that handle business logic.
  - `repository/`: Contains repository code for data storage/retrieval.
  - `go.mod`: Go module file for the packages.

- `go.work`: The workspace file for the project.

## Getting Started

1. Make sure you have Go installed. If not, you can download it from [https://golang.org/dl/](https://golang.org/dl/).

2. Clone this repository:

   ```sh
   git clone https://github.com/hussain18/client-manager.git
   ```
