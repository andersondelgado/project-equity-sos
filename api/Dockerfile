FROM golang:latest
WORKDIR /go/src/app
COPY server.go /go/src/app
COPY vendor /go/src/app/vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl
WORKDIR /root/
COPY --from=0 go/src/app/server .
CMD ["./server"]
LABEL version=demo-3