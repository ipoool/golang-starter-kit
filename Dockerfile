FROM golang:1.11 as builder

LABEL maintainer="asep saepulloh <asepulloh0109@gmail.com>"

WORKDIR /go/src/github.com/ipoool/laporcuranmor-api
ADD . /go/src/github.com/ipoool/laporcuranmor-api/

RUN sh -c "curl https://glide.sh/get | sh"

RUN cd /go/src/github.com/ipoool/laporcuranmor-api \
    && glide install \
    && CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o laporcuranmor-api -tags static_all .

EXPOSE 8078
CMD ["./laporcuranmor-api", "serve"]