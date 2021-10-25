#!/bin/bash

set -e

IMAGE_NAME="golang:1.16-alpine"

docker run -it --rm \
  --name cron_parser \
  -w /app \
  -v "$PWD":/app \
  ${IMAGE_NAME} \
  go \
  $*