FROM golang:1.17.2-buster as builder

ENV NGINX_VERSION=1.21.3

WORKDIR /work

RUN apt update && apt install -y libpcre3 libpcre3-dev zlib1g zlib1g-dev openssl libssl-dev
RUN git clone https://github.com/cubicdaiya/ngx_dynamic_upstream.git
RUN cd /work/ngx_dynamic_upstream && make build NGINX_VERSION=${NGINX_VERSION}
RUN cd /work/ngx_dynamic_upstream/tmp/nginx/${NGINX_VERSION}/nginx-${NGINX_VERSION} && make install

FROM nginx:1.21.3

COPY --from=builder /usr/local/nginx/sbin/nginx /usr/sbin/nginx
COPY --from=builder /usr/local/nginx /usr/local/nginx


CMD nginx -g "daemon off;"