# Istio Security Setup
- By default the Istio CA generates a self-signed root certificate and key and uses them to sign the workload certificates.

## mTLS Connection with Your Own CA Certificates
- Source: https://istio.io/latest/docs/tasks/security/cert-management/plugin-ca-cert/

### Create a Secret 'cacerts' and deploy PeerAuth
```bash
cd istio/security

export SECRET_NAME=cacerts # only this name working

kubectl create secret generic $SECRET_NAME -n istio-system \
      --from-file=cacerts/ca-cert.pem \
      --from-file=cacerts/ca-key.pem \
      --from-file=cacerts/root-cert.pem \
      --from-file=cacerts/cert-chain.pem

kubectl get secret $SECRET_NAME -n istio-system -o yaml

# k8s svc
curl cnk.com:9990
curl cnk.com:9990/ml

# istio gateway
curl gateway.cnk.com:80
curl gateway.cnk.com:80/ml
```

### Deploy PeerAuthentication
```bash
# activate mTLS
kubectl apply -f peerAuthentication.yaml

# k8s svc -> NOT WORKING
curl cnk.com:9990
curl cnk.com:9990/ml

# istio gateway -> WORKING
curl gateway.cnk.com:80
curl gateway.cnk.com:80/ml
```

### Connenction Test in Frontend-POD
```bash
FRONTEND_POD=$(k get pod -l app=frontend -o jsonpath='{.items[0].metadata.name}')
k exec -c istio-proxy $FRONTEND_POD -it -- bash
```

```bash
istio-proxy@cnk-frontend-64f77c4865-nf96d:/$ curl cnk.com:9990 -v
*   Trying 172.18.0.152:9990...
* Connected to cnk.com (172.18.0.152) port 9990 (#0)
> GET / HTTP/1.1
> Host: cnk.com:9990
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Recv failure: Connection reset by peer
* Closing connection 0
curl: (56) Recv failure: Connection reset by peer
```

```bash
istio-proxy@cnk-frontend-64f77c4865-nf96d:/$ curl gateway.cnk.com:80 -v
*   Trying 172.18.0.150:80...
* Connected to gateway.cnk.com (172.18.0.150) port 80 (#0)
> GET / HTTP/1.1
> Host: gateway.cnk.com
> User-Agent: curl/7.81.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< content-type: application/json
< date: Sun, 19 Jan 2025 13:46:41 GMT
< content-length: 117
< x-envoy-upstream-service-time: 0
< server: istio-envoy
< 
{"appName":"frontend","language":"golang","version":"1.1.2","podID":"cnk-frontend-64f77c4865-nf96d","env":"K8S-DEV"}
* Connection #0 to host gateway.cnk.com left intact
```