########### builder ###########
FROM golang AS builder

ENV GOPATH=

ENV CGO_ENABLED=0

RUN go mod init github.com/ShubhamTatvamasi/gin-redis

RUN go get github.com/gin-gonic/gin

RUN go get github.com/go-redis/redis/v8

COPY main.go .

RUN go build main.go

########### gin-redis ###########
FROM scratch

COPY --from=builder /go/main .

EXPOSE 80

CMD ["/main"]
