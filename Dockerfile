FROM golang:latest
MAINTAINER tubone24 <tubo.yyyuuu@gmail.com>
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN make install
RUN make build
EXPOSE 8585
