# Builder for Go
FROM golang as builder

WORKDIR ${GOPATH}/src/github.com/user/app
COPY . .
RUN set -x && \ 
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN echo ${GOPATH}/src/github.com/user/app

# Golang app
FROM scratch

WORKDIR /root/
COPY --from=builder /go/src/github.com/user/app/app .
CMD ["./app"]  
