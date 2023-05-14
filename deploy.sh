#!/bin/bash

# Reuse Minikube’s built-in Docker daemon
# eval $(minikube docker-env)

docker build -t welcome:1.0 -f welcome/Dockerfile .
docker build -t openweather-proxy:1.0 -f openweather-proxy/Dockerfile .

kubectl label namespace default istio-injection=enabled

# kubectl delete --all deployments --namespace=default

kubectl apply -f welcome.yaml
kubectl apply -f openweather-proxy.yaml

kubectl apply -f ingress.yaml
kubectl apply -f egress.yaml
