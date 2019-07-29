FROM golang:alpine
RUN mkdir -p $GOPATH/src/github.com/ahmdrz/rp
WORKDIR $GOPATH/src/github.com/ahmdrz/rp
ADD . .

RUN apk add --no-cache ca-certificates
RUN go build -i -o rp
RUN cp rp /usr/local/bin
EXPOSE 8080
CMD ["rp", "-v", "true"]