#!/bin/bash

# Set the variables
LOCAL_IMAGE="vira-backend-app"
NEW_TAG="latest"
DOCKER_HUB_USERNAME="polarbe8"
REPOSITORY_NAME="vira-backend-app"

# Build the updated Docker image
echo "Building the updated Docker image..."
docker build -t $LOCAL_IMAGE:$NEW_TAG .

# Tag the image
echo "Tagging the image..."
docker tag $LOCAL_IMAGE:$NEW_TAG $DOCKER_HUB_USERNAME/$REPOSITORY_NAME:$NEW_TAG

# Push the updated image to Docker Hub
echo "Pushing the updated image to Docker Hub..."
docker push $DOCKER_HUB_USERNAME/$REPOSITORY_NAME:$NEW_TAG

# Verify the image upload
echo "The updated image has been pushed to Docker Hub: $DOCKER_HUB_USERNAME/$REPOSITORY_NAME:$NEW_TAG"