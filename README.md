# Go URL Shortener

## Project Structure

The project is organized as follows:

| Directory/File            | Description                                      |
|---------------------------|--------------------------------------------------|
| `go-url-shortener/`        | Root directory of the project                   |
| `├── cmd/`                 | Main application entry point                     |
| `│   └── server/`          | Main app entry point                             |
| `│       └── main.go`      | The main Go file for running the server          |
| `├── internal/`            | Internal logic and modules                       |
| `│   ├── handler/`         | HTTP handlers (API layer)                        |
| `│   ├── service/`         | Business logic                                  |
| `│   ├── repository/`      | DB interactions                                  |
| `│   ├── model/`           | Structs and data models                          |
| `│   ├── utils/`           | Helper functions (e.g., generating short codes)  |
| `├── config/`              | Configuration files (e.g., environment variables, secrets) |
| `├── Dockerfile`           | Docker configuration for containerization        |
| `├── go.mod`               | Go module file                                  |
| `├── go.sum`               | Go module checksum file                         |
| `└── README.md`            | Project documentation (Readme.md file)                |

## Description

This project is a URL shortener implemented in Go. It provides a simple API to shorten long URLs and redirect users to the original URL using the shortened code.

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/work-kumar-rajesh/go-url-shortener.git


