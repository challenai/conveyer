#syntax=docker/dockerfile:1.2

FROM golang:1.18 as builder

WORKDIR /go/src/github.com/challenai/lightflow
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

####################################################################################################

FROM scratch as fsmysql

USER 8737

WORKDIR /home/lightflow

COPY --from=builder /go/src/github.com/challenai/lightflow/fsmysql /bin/

ENTRYPOINT [ "fsmysql" ]
