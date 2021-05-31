#!/bin/bash

podman build -t webhook-receiver:latest .
podman images
