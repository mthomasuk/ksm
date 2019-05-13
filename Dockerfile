FROM golang

WORKDIR /go/src/github.com/getfiit/ksm

ADD . /go/src/github.com/getfiit/ksm

RUN go get -t -v ./...

CMD ["go", "run", "example/basic.go"]

EXPOSE 8080