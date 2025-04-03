#!/bin/bash

set -e

IMAGE_NAME="golang:1.24-alpine3.20"

docker run -it --rm \
  --name cron_parser \
  -w /app \
  -v "$PWD":/app \
  "${IMAGE_NAME}" \
  go "$@"
