FROM golang:apline
WORKDIR /go
ADD . /go
CMD ["go", "run", "main.go"]