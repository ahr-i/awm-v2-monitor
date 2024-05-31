#!/bin/bash

# Variables
NETWORK_NAME="host"
#NETWORK_SUBNET="101.0.0.0/24"
IMAGE_NAME="awm-v2-monitor"
CONTAINER_NAME="awm-v2-monitor"
#CONTAINER_IP="101.0.0.2"

# Step 1: Docker Create Network
#docker network create --subnet $NETWORK_SUBNET $NETWORK_NAME
#if [ $? -ne 0 ]; then
#  echo "Network creation failed. An existing network already exists."
#fi

# Step 2: Docker Image Build
docker build -t $IMAGE_NAME . || { echo "Docker image build failed."; exit 1; }

# Step 3: Docker Container Start
docker run -it --rm --name $CONTAINER_NAME --network $NETWORK_NAME $IMAGE_NAME
if [ $? -ne 0 ]; then
  echo "Container execution failed."
  exit 1
fi
