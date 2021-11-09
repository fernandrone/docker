FROM golang:1.16

WORKDIR /go/src/github.com/fernandrone/bootcamp-docker/sextoandar

COPY static static
COPY main.go main.go
COPY go.mod go.mod

RUN go build -o website

EXPOSE 3000

ENTRYPOINT [ "./website" ]
