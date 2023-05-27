FROM golang:1.17-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/busovilya/bitcoin_rate_mailer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bitcoin_rate_mailer .
EXPOSE 10000

ENTRYPOINT ["./bitcoin_rate_mailer"]