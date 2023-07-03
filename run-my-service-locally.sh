#!/bin/bash
set -e

docker build -t myservice .
docker run -p 8080:8080 myservice
