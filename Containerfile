FROM golang:latest as builder

WORKDIR /build

COPY . .

RUN pwd

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./gosockets ./cmd/main

RUN ls -lrta

FROM alpine:latest as app-runner

WORKDIR /root/

COPY --from=builder /build/gosockets .

RUN ls -lrta

RUN pwd

EXPOSE 80

ENTRYPOINT ["./gosockets"]



