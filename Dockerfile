ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy the entire application source code
COPY . .

# Build the Go application specifying the main package location
RUN go build -v -o /run-app ./cmd/app

FROM debian:bookworm

# Copy the compiled binary from the builder stage
COPY --from=builder /run-app /usr/local/bin/

# Copy the production configuration as env.yml
COPY production.env.yml /env.yml

# Specify the default command
CMD ["run-app"]
