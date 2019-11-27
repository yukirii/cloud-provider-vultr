FROM golang:1.13 AS builder
RUN  apt-get update \
       && apt-get -y install make
ADD . /go/src/github.com/yukirii/vultr-cloud-provider
WORKDIR /go/src/github.com/yukirii/vultr-cloud-provider
RUN ["make", "clean", "build"]

FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /go/src/github.com/yukirii/vultr-cloud-provider/bin/vultr-cloud-controller-manager .
CMD ["/vultr-cloud-controller-manager"]
