#!/bin/bash

export INGRESS_NAME=istio-ingressgateway
export INGRESS_NS=istio-system
export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export SECURE_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="https")].port}')

curl -v -HHost:welcome.local --resolve "welcome.local:$SECURE_INGRESS_PORT:$INGRESS_HOST" \
  --cacert certs/example.com.crt --cert certs/client.example.com.crt --key certs/client.example.com.key \
  "https://welcome.local:$SECURE_INGRESS_PORT/weather/zbfpbj"
