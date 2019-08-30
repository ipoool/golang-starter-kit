FROM golang:1.11 as builder

LABEL maintainer="asep saepulloh <asepulloh0109@gmail.com>"

WORKDIR /go/src/github.com/ipoool/golang-starter-kit
ADD . /go/src/github.com/ipoool/golang-starter-kit/

RUN sh -c "curl https://glide.sh/get | sh"

RUN cd /go/src/github.com/ipoool/golang-starter-kit \
    && glide install \
    && CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o golang-starter-kit -tags static_all .

EXPOSE 8078
CMD ["./golang-starter-kit", "serve"]