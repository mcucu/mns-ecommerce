# Start from golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

# Copy go mod and sum files
COPY ./go.mod ./
COPY ./go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd

# Finally our multi-stage to build a small image
# Start a new stage from alpine
FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .
COPY ./.env.example ./.env

# Run executable
CMD ["./main"]
