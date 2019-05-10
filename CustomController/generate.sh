#!/bin/bash -e

CURRENT_DIR=$(pwd)
GEN_DIR=$(dirname $0)
REPO_DIR="$CURRENT_DIR/$GEN_DIR"

PROJECT_MODULE="github.com/enixdark/kubernetes-programming-deepdive/CustomController"
IMAGE_NAME="kubernetes-codegen:latest"

CUSTOM_RESOURCE_NAME="controller"
CUSTOM_RESOURCE_VERSION="v1alpha1"

echo "Building codegen Docker image..."
docker build -f "${GEN_DIR}/Dockerfile.code-generate" \
             -t "${IMAGE_NAME}" \
             "${REPO_DIR}"

cmd="./generate-groups.sh all \
    "$PROJECT_MODULE/pkg/client" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

echo "Generating client codes..."
echo $REPO_DIR
echo $PROJECT_MODULE
echo $cmd

docker run --rm \
           -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd

sudo chown $USER:$USER -R $REPO_DIR/pkg