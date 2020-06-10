# gin-redis

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/shubhamtatvamasi/gin-redis)](https://hub.docker.com/r/shubhamtatvamasi/gin-redis)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/shubhamtatvamasi/gin-redis?sort=semver)](https://hub.docker.com/r/shubhamtatvamasi/gin-redis)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/shubhamtatvamasi/gin-redis/latest)](https://hub.docker.com/r/shubhamtatvamasi/gin-redis)
[![Docker Pulls](https://img.shields.io/docker/pulls/shubhamtatvamasi/gin-redis)](https://hub.docker.com/r/shubhamtatvamasi/gin-redis)
[![MicroBadger Layers (tag)](https://img.shields.io/microbadger/layers/shubhamtatvamasi/gin-redis/latest)](https://hub.docker.com/r/shubhamtatvamasi/gin-redis)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/shubhamtatvamasi/gin-redis)](https://hub.docker.com/r/shubhamtatvamasi/gin-redis)

Deploy Gin and Redis
```bash
kubectl apply -f https://raw.githubusercontent.com/ShubhamTatvamasi/gin-redis/master/deployment.yaml
```

scale deployment
```
kubectl scale deployment gin-redis --replicas 10
```

Delete Deployment
```bash
kubectl delete -f https://raw.githubusercontent.com/ShubhamTatvamasi/gin-redis/master/deployment.yaml
```

