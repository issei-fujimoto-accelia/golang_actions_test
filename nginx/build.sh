#!/bin/sh
docker build --no-cache . -t nginx-1.21.3:v1.0
docker tag nginx-1.21.3:v1.0 ghcr.io/builed-nginx/nginx-1.21.3:v1.0
docker push ghcr.io/builded-nginx/nginx-1.21.3:v1.0
