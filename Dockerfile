FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./bin/main cmd/main.go


FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/bin/main .

EXPOSE 50051
EXPOSE 2112

CMD [ "./main" ]