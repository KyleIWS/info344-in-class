#!/usr/bin/env bash
docker run -d  \
-p 443:443 \
--name zipsvr \
-v $(pwd)/TLS:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
kylews/zipsvr
