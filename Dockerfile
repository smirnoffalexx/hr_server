FROM --platform=linux/amd64 golang:1.24-alpine AS build

WORKDIR /src

COPY . .

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g internal/api/http/routing/routing.go --output ./internal/api/http/docs

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /src/server /app/server

EXPOSE 8080

ENTRYPOINT ["/app/server"]
