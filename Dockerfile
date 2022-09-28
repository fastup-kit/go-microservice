FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o app api/api.go

FROM alpine

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app ./

CMD ["./app"]