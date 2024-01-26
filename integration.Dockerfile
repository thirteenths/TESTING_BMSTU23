FROM golang:1.21

WORKDIR /app/

ENV CGO_ENABLED 1
ENV GOPATH /go
#ENV GOCACHE /go-build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod /app/
COPY go.sum /app/
RUN go mod download && go mod verify

COPY . /app/

CMD ["make", "integration-test"]