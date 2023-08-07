FROM ubuntu:focal

ENV NGINX_VERSION 1.23.1
ENV OPENSSL_VERSION 1.1.1h
ENV SITEGUARD_LITE_VERSION 3.40-2
ENV NGX_MRUBY_BRANCH=master
ENV NGX_MRUBY_REVISION=11956b6b4a51af1f26db520feac1d5effd2c9152
ENV TZ=Asia/Tokyo
ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /

RUN apt-get -qq update \
  && apt-get -qq install \
     libapr1 libapr1-dev libaprutil1 libaprutil1-dev build-essential bison cron libpcre3-dev libpcre++-dev debhelper \
     flex gcc make automake autoconf libtool git libreadline6-dev \
     zlib1g-dev libncurses5-dev libssl-dev rake libpam0g-dev \
     autotools-dev cgroup-lite git wget \
  && apt-get -y clean \
  && rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/matsumotory/ngx_mruby

RUN wget http://www.openssl.org/source/openssl-$OPENSSL_VERSION.tar.gz \
  && tar zxf openssl-$OPENSSL_VERSION.tar.gz \
  && rm openssl-$OPENSSL_VERSION.tar.gz

COPY nginx /nginx-1.23.1

WORKDIR /ngx_mruby
RUN ./configure \
  --with-ngx-src-root=/nginx-$NGINX_VERSION \
  --with-ngx-config-opt="--with-http_v2_module \
                         --with-http_ssl_module \
                         --with-stream --with-debug \
                         --with-http_stub_status_module \
                         --without-stream_access_module \
                         --prefix=$prefix \
                         --with-openssl-opt=enable-tls1_3" \
  --with-openssl-src=/openssl-$OPENSSL_VERSION

RUN make
RUN mkdir -p /usr/local/nginx
env DESTDIR=/usr/local/nginx
RUN make install

RUN ln -s /usr/local/nginx/conf /etc/nginx \
  && ln -s /usr/local/nginx/sbin/nginx /usr/sbin/nginx \
  && ln -s /usr/local/nginx/logs /var/log/nginx

COPY ./server.crt /etc/nginx/ssl/server.crt
COPY ./server.key /etc/nginx/ssl/server.key

COPY ./nginx.conf /etc/nginx/nginx.conf

COPY ./src /usr/share/nginx/html

WORKDIR /usr/local/nginx

# gdb debugging
RUN apt-get update && apt-get install -y gdb gdbserver
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
