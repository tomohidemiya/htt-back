FROM golang:latest

RUN mkdir -p $GOPATH/src/htt-back
WORKDIR $GOPATH/src/htt-back/
RUN go get -u github.com/codegangsta/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/jinzhu/gorm/dialects/postgres

CMD gin -i run
