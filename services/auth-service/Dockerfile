FROM golang:1.25 AS build

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o auth-service ./cmd/server

FROM gcr.io/distroless/base-debian12

COPY --from=build /app/auth-service /usr/local/bin/auth-service

EXPOSE 8080

ENTRYPOINT ["auth-service"]
