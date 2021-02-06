FROM golang:latest

RUN mkdir -p $GOPATH/src/htt-back
WORKDIR $GOPATH/src/htt-back/
RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/jinzhu/gorm/dialects/postgres
