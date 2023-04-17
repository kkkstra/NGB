FROM golang:alpine AS builder
ENV GOOS=linux\
  GO111MODULE=on \
  GOPROXY=https://goproxy.cn,direct
COPY . /src
WORKDIR /src
RUN cd cmd && go build -o /bin/app .

FROM alpine:latest
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories
RUN apk update \
  && apk upgrade \
  && apk add --no-cache bash \
  bash-doc \
  bash-completion \
  && rm -rf /var/cache/apk/* \
  && /bin/bash
COPY ./configs /configs
COPY --from=builder /bin/app /bin/app
CMD /bin/app