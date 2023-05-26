# 18 Dockerfile

## Prompt

Can you help me create a multi-stage Dockerfile that contains the built go application binary in a minimalist base image with an entrypoint that runs the app in daemon mode?

## Response

Sure, here is an example of a multi-stage Dockerfile that compiles your Go app and then packages it into a small scratch (empty) container:

```Dockerfile
# First stage: Build the application
FROM golang:1.17 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Second stage: Build a small image
# Start from scratch
FROM scratch

# Document that the service listens on port 8080
EXPOSE 8080

# Copy the binary from builder
COPY --from=builder /app/main /main

# Run the binary
ENTRYPOINT ["/main", "-d"]
```

This Dockerfile is fairly straightforward:

We are starting from the official golang image to build the Go application. The builder image has all the necessary tools to build a Go app.
We are setting the working directory to /app.
We are copying the go.mod and go.sum files and downloading the dependencies using go mod download. This means that Docker will cache the dependencies unless the go.mod or go.sum files change.
Then we are copying the rest of the application code and building it with go build.
For the second stage, we are starting from an empty image to keep it small.
Finally, we are copying the compiled binary from the builder stage and defining an entrypoint to run the binary in daemon mode when the container starts.
You can build the Docker image with the following command:

```bash
docker build -t password-checker .
```

And then you can run the container with:

```bash
docker run -p 8080:8080 password-checker
```

Please make sure to replace main in the Dockerfile with the actual name of your main executable (the output of the go build command).
