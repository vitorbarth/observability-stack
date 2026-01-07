FROM golang:1.25 AS build

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o account-service ./cmd/server

FROM gcr.io/distroless/base-debian12

COPY --from=build /app/account-service /usr/local/bin/account-service

EXPOSE 8081

ENTRYPOINT ["account-service"]