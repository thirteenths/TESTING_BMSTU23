FROM golang:1.21.2-alpine

COPY ../../ .
RUN go mod download

WORKDIR /cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o api

CMD ["/cmd/api/api"]