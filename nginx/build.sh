#!/bin/sh
docker build --no-cache . -t nginx-1.21.3:v1.0
docker tag nginx-1.21.3:v1.0 ghcr.io/issei-fujimoto-accelia/nginx-1.21.3:v1.0
docker push ghcr.io/issei-fujimoto-accelia/nginx-1.21.3:v1.0
