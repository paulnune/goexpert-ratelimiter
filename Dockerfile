FROM golang:1.22.5 AS builder
WORKDIR /app
COPY . .
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o rate-limit cmd/api/main.go
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rate-limit cmd/api/main.go

FROM scratch
EXPOSE 8080
COPY --from=builder /app/rate-limit .
CMD ["./rate-limit"]