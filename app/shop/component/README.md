## SHOP APPS

### Quick Debug
```bash
#---------------------
# frontend Service
#---------------------

cd app/shop/component/frontend
export IMAGE_NAME="yuyatinnefeld/cnk-frontend"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9990:9990 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

# verify
curl http://localhost:9990
curl http://localhost:9990/health
curl http://localhost:9990/fetch-item?index=1

# push image
docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.1.2
docker image push ${IMAGE_NAME}:1.1.2

#---------------------
# Inventory Service
#---------------------

cd app/shop/component/inventory
export IMAGE_NAME="yuyatinnefeld/cnk-backend-inventory"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9991:9991 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

# verify
curl http://localhost:9991
curl http://localhost:9991/health
curl http://localhost:9991/items/2

# push image
docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.1.4
docker image push ${IMAGE_NAME}:1.1.4

#---------------------
# ML Service
#---------------------

cd app/shop/component/ml
export IMAGE_NAME="yuyatinnefeld/cnk-backend-ml"
docker build -t ${IMAGE_NAME} .
docker run -it -p 9992:9992 -e ENV=DOCKER-DEV --net=host ${IMAGE_NAME}

docker image tag ${IMAGE_NAME} ${IMAGE_NAME}:1.1.2
docker image push ${IMAGE_NAME}:1.1.2
```

### Image Build Pipeline
- Github Action