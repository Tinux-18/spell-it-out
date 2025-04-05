#!/bin/sh
IMAGE=spell-it-out
CONTAINER=spell-it-out

# Remove existing container in case it already exists.
docker container rm $IMAGE --force || true

# Build a new image using a custom Docker file.
docker image build -t $IMAGE -f Dockerfile .

# Run a container from the image.
# Connect the container the the host network.
# Name the container.
# Mount the current working directory to the working directory.
# Automatically remove exited containers.
docker run -it --network host --name $CONTAINER --rm $IMAGE
