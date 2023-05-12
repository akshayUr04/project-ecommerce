# build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache git \
    && go mod download \
    && go build -o ./out/dist cmd/api/main.go

# production stage
FROM alpine:3.16
COPY --from=builder /app/out/dist /app/
COPY template ./template
WORKDIR /app
EXPOSE 3000
CMD ["./dist"]

#go embed 