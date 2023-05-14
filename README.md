# OpenWeather ROT-13 proxy

Demo istio service mash project with ingress/egress gateways.

# Deploy

```bash
# Create secret with api key
kubectl create secret generic openweather --from-literal=api-key='[YOUR OWM API KEY]'
# Run build & deploy script
./deploy.sh
# Verify installation (Set INGRESS_HOST / INGRESS_PORT before) 
curl -HHost:welcome.local "http://$INGRESS_HOST:$INGRESS_PORT/weather/zbfpbj" -v
# Should return current temperature in Moscow (ROT-13 zbfpbj).
```