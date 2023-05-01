# build stage
FROM golang:1.19-alpine3.16 AS builder
# Set destination for 
WORKDIR /app
# Copy the hole directory
COPY . .
# Download Go modules
RUN go mod download
# Build the code
RUN  go build -o ./out/dist cmd/api/main.go
# Binding the port
EXPOSE 8080
# Run
CMD ["/app/out/dist"]