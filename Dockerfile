FROM golang:alpine as build
RUN mkdir -p $GOPATH/src/github.com/ahmdrz/rp
WORKDIR $GOPATH/src/github.com/ahmdrz/rp
ADD . .

RUN go build -o /app/rp

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build /app/rp /usr/local/bin/rp
