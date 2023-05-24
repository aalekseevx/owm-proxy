#!/bin/bash

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=example.com' -keyout certs/example.com.key -out certs/example.com.crt

openssl req -out certs/welcome.local.csr -newkey rsa:2048 -nodes -keyout certs/welcome.local.key -subj "/CN=welcome.local/O=httpbin organization"
openssl x509 -req -sha256 -days 365 -CA certs/example.com.crt -CAkey certs/example.com.key -set_serial 0 -in certs/welcome.local.csr -out certs/welcome.local.crt

openssl req -out certs/client.example.com.csr -newkey rsa:2048 -nodes -keyout certs/client.example.com.key -subj "/CN=client.example.com/O=client organization"
openssl x509 -req -sha256 -days 365 -CA certs/example.com.crt -CAkey certs/example.com.key -set_serial 1 -in certs/client.example.com.csr -out certs/client.example.com.crt


kubectl -n istio-system delete secret welcome-credential
kubectl create -n istio-system secret generic welcome-credential \
  --from-file=tls.key=certs/welcome.local.key \
  --from-file=tls.crt=certs/welcome.local.crt \
  --from-file=ca.crt=certs/example.com.crt
