FROM golang:alpine AS builder
ENV GOOS=linux\
  GO111MODULE=on \
  GOPROXY=https://goproxy.cn,direct
COPY . /src
WORKDIR /src
RUN cd cmd && go build -o /bin/app .

FROM alpine:latest
COPY ./configs /configs
COPY --from=builder /bin/app /bin/app
CMD /bin/app