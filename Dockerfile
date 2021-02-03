FROM golang:1.15.7-alpine

WORKDIR /test-server

ADD main.go /test-server

RUN apk add git
RUN go get -v -u github.com/go-chi/chi
RUN cd /test-server && go build

EXPOSE 80

ENTRYPOINT ./test-server 
