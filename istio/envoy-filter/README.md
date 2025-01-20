# Envoy Filter

## Setup Env

```bash
# Deploy sleep app
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.23/samples/sleep/sleep.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.23/samples/httpbin/httpbin.yaml

export SOURCE_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})
export DESTINATION_SVC=$(kubectl get svc -l app=httpbin -o jsonpath={.items..metadata.name})
export DESTINATION_URL=$DESTINATION_SVC:8000

# Deploy telemetry to activate access logs
kubectl apply -f telemetry.yaml

# Access httpbin
kubectl exec $SOURCE_POD -c sleep -- curl -sS -v $DESTINATION_URL/status/200

# Check the traffic from access log
kubectl logs -l app=sleep -c istio-proxy --tail=100
```

## Envoy Filter 1 - Add X-Envoy-Ip-Tags
```bash
kubectl apply -f filters/ev-filter-ip-tagging.yaml
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS -H "Authorization: Bearer MY_USER_TOKEN" -v $DESTINATION_URL/get
# headers -> "X-Envoy-Ip-Tags": "my-filter",
kubectl delete -f ev-filter-ip-tagging.yaml # cleanup
```

## Envoy Filter 2 - Remove X-Envoy Param
```bash
# you see x-envoy-upstream-service-time
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS -v $DESTINATION_URL/status/200

k apply -f filters/remove-x-envoy-param.yaml

# x-envoy-upstream-service-time is disappeared
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS -v $DESTINATION_URL/status/200
```

## Envoy Filter 3 - Updates Header Values 
```bash
# check the Mymessage and X-Request-Id
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS -v -H "myMessage: testYUYA" $DESTINATION_URL/get  | jq '.headers'
k apply -f filters/update-headers.yaml

# check the Mymessage and X-Request-Id
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS -v -H "myMessage: testYUYA" $DESTINATION_URL/get  | jq '.headers'
```
