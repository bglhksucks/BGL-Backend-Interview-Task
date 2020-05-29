FROM golang:1.14.2-alpine

WORKDIR /go/src/bgl
COPY . .

RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o main .

CMD ["./main"]