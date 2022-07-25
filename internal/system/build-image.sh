#!/bin/sh
set -ex

readonly ROOT_DIR=`pwd`
# readonly TAG=${TAG:=$(git rev-parse --short HEAD)}
readonly TAG="1.0.0"
readonly NAMESPACE=${NAMESPACE:="anhnguyen0809"}
readonly SERVICE_PATH=${ROOT_DIR}/deploy/dev/docker/web
readonly SERVICE_NAME=rin-system-service

echo "Namespace is ${NAMESPACE} and tag is ${TAG}"
echo "Start to build ${SERVICE_NAME}..."

docker build -f $SERVICE_PATH/Dockerfile \
    -t $NAMESPACE/$SERVICE_NAME:$TAG \
    -t $NAMESPACE/$SERVICE_NAME:latest .
