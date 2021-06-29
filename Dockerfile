FROM golang:1.16

WORKDIR /go/src/app
COPY . .
# timezone change
RUN cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN go build -o main .

RUN apt clean -y
RUN apt update -y 
RUN apt upgrade -y 
RUN apt install -y sudo cron

ADD cron.d /etc/cron.d/
RUN chmod 0644 /etc/cron.d/*
RUN crontab /etc/cron.d/*

RUN touch /var/log/hoge.log

CMD env > /root/root.env && cron && tail -f /var/log/hoge.log
