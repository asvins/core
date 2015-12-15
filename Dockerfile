FROM golang:1.5.1

WORKDIR /go/src/app
COPY . /go/src/app

RUN mkdir upload
RUN go get .
CMD app

EXPOSE 8080
