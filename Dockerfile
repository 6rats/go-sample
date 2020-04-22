FROM golang:1.13.7-buster

RUN mkdir /go/src/go-sample
WORKDIR /go/src/go-sample

RUN curl -fLo /go/bin/air https://git.io/linux_air \
  && chmod +x /go/bin/air

COPY . ${WORKDIR}

CMD air
