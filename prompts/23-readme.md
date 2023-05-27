# Generate Readme

## Prompt

Can you help me create a README.md file that explains the high level summary of this project and the tooling included and how to work with it?

## Response

Sure, here's a basic `README.md` file:

```markdown
# Password Strength Checker

The Password Strength Checker is a command-line and HTTP-based service written in Go that checks the strength of a password based on certain rules. It is built to be production-ready with operational features including health checks, metrics, Docker support, and load testing configuration.

## Functional Requirements

The application checks if a password is strong based on the following criteria:
- It has at least 6 characters and at most 20 characters.
- It contains at least one lowercase letter, at least one uppercase letter, and at least one digit.
- It does not contain three repeating characters in a row (i.e., "aaa").

## Features

### Command-line Interface

You can use the password checker in the command-line with a simple input password.

### HTTP Interface

The password checker can also be run as a HTTP server that takes in password as a parameter and returns the steps required to make it strong.

### Health Check Endpoint

The HTTP server also includes a `/healthcheck` endpoint that returns a 200 success HTTP status value.

### Metrics Endpoint

The HTTP server includes a `/metrics` endpoint that exposes Prometheus metrics. It currently tracks the total number of requests made to the password strength check endpoint.

### Docker Support

This service can be built and run as a Docker container using the provided Dockerfile and docker-compose.yml files.

### Load Testing Configuration

A k6 configuration file (`loadtest.js`) is provided for load testing the HTTP password checker web service.

### Task Runner

A Go-Task `Taskfile.yml` is provided for automation of common tasks like running tests, building the application, and more.

### Logging

The HTTP service logs each request made, with adjustable log level setting to control the verbosity of the logs.

## Usage

### CLI

Run `go run main.go <password>` to check the strength of `<password>`.

### HTTP server

Run `go run main.go -d` to start the HTTP server. You can then send a GET request to `localhost:8080/check/<password>` to check the strength of `<password>`.

## Build

You can use Docker Compose to build and run the service:

```bash
docker-compose up --build
```

## Test

Unit tests can be run with the Go test command:

```bash
go test
```

## Task Runner

You can use the Go-Task task runner to automate common tasks. For example, to run tests:

```bash
task test
```

## Load Testing

You can use k6 for load testing. First install k6, then run:

```bash
k6 run loadtest.js
```
```

Feel free to modify the above README file to fit your project needs better. It's important to keep your README updated as your project evolves, as it's usually the first thing people see when they come to your project.