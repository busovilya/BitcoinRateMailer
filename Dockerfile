FROM golang:1.17-alpine
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/busovilya/BitcoinRateMailer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o BitcoinRateMailer .
EXPOSE 10000

ENTRYPOINT ["./BitcoinRateMailer"]