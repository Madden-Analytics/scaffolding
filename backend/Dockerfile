FROM golang:1.27.1 AS base

WORKDIR /backend

# Hot reloading mod
RUN go install github.com/air-verse/air@v1.67.4
EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]
