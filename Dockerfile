FROM golang:1.21.4-bullseye
LABEL authors="tiago98751"

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -tags=jsoniter src/main.go

CMD ["./main"]