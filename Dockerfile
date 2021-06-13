FROM golang:1.14.3-alpine AS build
RUN apk add git
WORKDIR /go/src/coursera
COPY . .
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/pg.v4
RUN GO_ENABLED=0 go build -o ../bin/course
WORKDIR /go/src
EXPOSE 8080
ENTRYPOINT [ "./bin/course" ]