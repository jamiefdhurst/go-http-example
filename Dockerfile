FROM golang:1.15

WORKDIR /go/src/github.com/jamiefdhurst/go-http-example
COPY . .

RUN go install

EXPOSE 8080

CMD ["go-http-example"]
