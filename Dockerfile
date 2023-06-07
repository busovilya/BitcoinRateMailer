FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.mod
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

FROM debian AS runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 10000
ENTRYPOINT ["./main"]