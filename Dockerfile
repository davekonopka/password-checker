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
RUN CGO_ENABLED=0 GOOS=linux go build -o password-checker .

# Second stage: Build a small image
# Start from scratch
FROM scratch

# Document that the service listens on port 8080
EXPOSE 8080

# Copy the binary from builder
COPY --from=builder /app/password-checker /password-checker

# Run the binary
ENTRYPOINT ["/password-checker", "-d"]
