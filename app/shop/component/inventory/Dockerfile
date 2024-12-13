FROM golang:1.17.5 AS go_build
ENV CGO_ENABLED=0
WORKDIR /go-app
COPY src/main.go /go-app
RUN go mod init go-app
RUN go build -o server

FROM alpine:3.15
WORKDIR /go-app
COPY --from=go_build /go-app/server /go-app/server
EXPOSE 9991
CMD ["./server"]