FROM golang:latest

RUN mkdir -p $GOPATH/src/htt-back
WORKDIR $GOPATH/src/htt-back/
RUN go get godep && \
    go get -u github.com/gin-gonic/gin && \
    go get -u gorm.io/gorm && \
    go get -u gorm.io/driver/postgres
