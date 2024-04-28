FROM golang:1.22.1
WORKDIR /app
# RUN go install github.com/spf13/cobra-cli@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 1323
CMD ["go", "run", "main.go"]
