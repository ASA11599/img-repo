FROM golang:1.14

RUN mkdir /images

WORKDIR /go/src/img-repo-api

COPY ./img-repo-api .

RUN go get
RUN go build main.go

CMD ["./main"]
