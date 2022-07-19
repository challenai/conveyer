#syntax=docker/dockerfile:1.2

FROM golang:1.18 as builder

WORKDIR /go/src/github.com/challenai/lightflow
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -v -ldflags '-extldflags -static' -o lightflow cmd/main.go

####################################################################################################

FROM scratch as argocli

USER 8737

WORKDIR /home/lightflow

COPY --from=builder /go/src/github.com/challenai/lightflow/lightflow /bin/

ENTRYPOINT [ "lightflow" ]
