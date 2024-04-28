FROM golang:1.21.9-alpine as build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test --tags=unit -v ./...

RUN go build -o ./out/go-app .

FROM alpine:3.19.1
COPY --from=build-base /app/out/go-app /app/go-app

EXPOSE 8080

CMD ["/app/go-app"]