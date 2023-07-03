FROM golang:1.14-stretch as builder
RUN mkdir /mnt/workspace
COPY . /mnt/workspace
WORKDIR /mnt/workspace
RUN go get -v -u github.com/gorilla/mux \
    && CGO_ENABLED=0 go build -o MyService .
RUN go test

FROM alpine:3.10
RUN apk --no-cache -U add ca-certificates=20191127-r2
EXPOSE 8080
COPY --from=builder /mnt/workspace/MyService /
COPY --from=builder /mnt/workspace/VERSION /
CMD ["/MyService"]
