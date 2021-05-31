#!/bin/bash

mkdir -p logs
podman run -d -p 8080:8080 --volume $(pwd)/logs:/app/log:Z --name webhook-receiver webhook-receiver:latest
