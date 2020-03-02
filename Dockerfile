FROM golang:alpine

LABEL maintainer="kanhaiya1501@gmail.com"
LABEL version="1.0.0"

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep
RUN mkdir /go/src/gra-gin
COPY . /go/src/gra-gin
WORKDIR /go/src/gra-gin
RUN dep ensure && go build main.go
EXPOSE 8080

CMD [ "./main" ]