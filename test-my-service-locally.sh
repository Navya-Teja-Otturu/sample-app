#!/bin/bash
set -e

docker build -f testDockerfile -t test-my-service .
docker run --network host test-my-service
