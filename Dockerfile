FROM golang:1.7

COPY ./vendors/ /go/src/
COPY ./src/ /go/src/

WORKDIR /go

RUN go build src/gform/gform.go

EXPOSE 8080

CMD ./gform
