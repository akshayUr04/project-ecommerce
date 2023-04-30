# build stage
FROM golang:1.19-alpine3.16 AS builder
#maintainer info
LABEL maintainer="akshayur <akshayur@gmail.com>"
#installing git
RUN apk update && apk add --no-cache git

WORKDIR /Job-Portal

COPY . .

RUN apk add --no-cache make

RUN make deps
RUN go mod vendor
RUN make build



# Run stage
FROM alpine:3.16

WORKDIR /project-ecommerce
COPY go.mod .
COPY go.sum .
COPY views ./views
COPY --from=builder /Job-Portal/build/bin/api .


CMD [ "/ecommerce/api"] 