#!/bin/bash

# Using git to version control the images
VERSION=$(git rev-parse HEAD)

# These are optional arguments
# usage - ./deploy.sh <docker-username> <go-api-name>
# example - ./deploy.sh someuser weather-service-go-api
DOCKER_USERNAME=${1-happilymarrieddadudemy}
GO_API_REPO=${2-weather-service-go-api}

## Push up and set GO API
docker build -t ${DOCKER_USERNAME}/${GO_API_REPO}:${VERSION} -f ./Dockerfile .
docker push ${DOCKER_USERNAME}/${GO_API_REPO}:${VERSION}

docker build -t ${DOCKER_USERNAME}/${GO_API_REPO}:latest -f ./Dockerfile .
docker push ${DOCKER_USERNAME}/${GO_API_REPO}:latest


sleep 1

echo "Completed!"