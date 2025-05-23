# build stage
FROM golang:1.24-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o todolist

# final stage
FROM alpine:latest
COPY --from=builder /build/todolist /todolist
COPY .env .env
EXPOSE 8080
CMD ["/todolist"]
