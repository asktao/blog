FROM golang:1.12

#Setup timezone
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

CMD mkdir /blog

ADD . /blog

WORKDIR /blog

RUN GO111MODULE=on

RUN go build -o app .


CMD ["./app"]
#ENTRYPOINT["./app"]

EXPOSE 8080