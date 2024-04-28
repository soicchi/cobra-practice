FROM golang:1.22.1 AS builder
WORKDIR /app
# RUN go install github.com/spf13/cobra-cli@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /bin/main .

FROM golang:1.22.1-bullseye AS runtime
WORKDIR /app
COPY --from=builder /go/pkg /go/pkg
COPY --from=builder /bin/main /bin/main
EXPOSE 1323
CMD ["/bin/main"]
