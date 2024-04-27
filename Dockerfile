FROM golang:1.21.9-alpine as build-base

WORKDIR /go/app/base

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 

RUN go build -o ./out/go-app .

FROM alpine:3.16.2
COPY --from=build-base /app/out/go-app /app/go-app

CMD ["/app/go-app"]