FROM golang:latest
MAINTAINER tubone24 <tubo.yyyuuu@gmail.com>
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN make backend-install
RUN make backend-build
EXPOSE 8585
