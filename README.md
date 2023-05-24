# OpenWeather ROT-13 proxy

Demo istio service mash project with ingress/egress gateways.

# Deploy

```bash
# Create secret with api key
kubectl create secret generic openweather --from-literal=api-key='[YOUR OWM API KEY]'
# Generate certs
./certgen.sh
# Run build & deploy script
./deploy.sh
# Verify installation
./request.sh
# Should return current temperature in Moscow
```