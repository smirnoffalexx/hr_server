FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init -g internal/api/http/routing/routing.go --output ./internal/api/http/docs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go/bin/app cmd/app/main.go

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata && \
    mkdir -p /hr-server

COPY --from=build /go/bin/app /hr-server/app

WORKDIR /hr-server

ENV TZ=UTC

EXPOSE 8080

CMD ["/hr-server/app"]