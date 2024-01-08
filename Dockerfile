FROM golang:1.21 as builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./main

FROM ubuntu:24.04

# RUN apt-get update -y

WORKDIR /app

COPY --from=builder /usr/src/app/main .

EXPOSE 8085

ENTRYPOINT [ "./main" ]