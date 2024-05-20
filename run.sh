#!/bin/bash
echo "--- Ping-Pong test application ---"
echo "building ping service..."
docker build -t ping -f docker/ping/Dockerfile .
echo "building pong service..."
docker build -t pong -f docker/pong/Dockerfile .
echo "setting up kube storage..."
kubectl apply -f kube/storage.yaml
echo "setting up pong service..."
kubectl apply -f kube/pong_pod.yaml
echo "setting up ping service..."
kubectl apply -f kube/ping_cron.yaml