FROM golang:1.15.2-alpine AS dev

WORKDIR /app

ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk git \
  && go get github.com/gravityblast/fresh
# github.com/pilu/fres から go get できなくて、↑にしたらいけた
# && go get github.com/pilu/fres

COPY . .

CMD ["fresh"]

FROM golang:1.15.2-alpine AS builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine AS prod

WORKDIR /app

COPY --from=builder /src/main .

CMD ["./main"]