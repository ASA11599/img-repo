FROM golang:1.14

WORKDIR /go/src/api

COPY ./img-search-api .

RUN go get
RUN go build main.go

CMD ["./main"]
