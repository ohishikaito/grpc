FROM golang:1.15.2-alpine AS builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine

WORKDIR /app

COPY --from=builder /src/main .

CMD ["./main"]