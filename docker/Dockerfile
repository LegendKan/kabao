# kabao kastore.cc

FROM ubuntu:14.04

MAINTAINER kanchuanqi <kanchuanqi@gmail.com>

RUN apt-get update && apt-get install -y nginx wget git

RUN mv /etc/nginx/nginx.conf /home/nginx.conf.bak

COPY ./nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /home/ssl

COPY ./ssl /home/ssl

#RUN service nginx restart

RUN wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz \
    && rm go1.6.2.linux-amd64.tar.gz \
    && mkdir -p /home/go/src

ENV GOROOT /usr/local/go

ENV GOPATH /home/go

ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /home/go/src

RUN git clone https://LegendKan:19901124ubuntu@github.com/LegendKan/kabao.git 

WORKDIR kabao

RUN go get github.com/astaxie/beego \
	&& go get github.com/beego/bee \
	&& go get

EXPOSE 80 443

ENTRYPOINT service nginx restart && bee run
