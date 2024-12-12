FROM golang:1.23-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go install github.com/air-verse/air@latest

COPY . ./

CMD ["air", "-c", "/app/.air.toml"]
