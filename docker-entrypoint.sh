#!/bin/bash

if [ -z "$KUBERNETES_TOKEN" ]; then
	KUBERNETES_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
fi

kubectl config set-cluster k8s --server=https://$KUBERNETES_SERVICE_HOST:$KUBERNETES_SERVICE_PORT --insecure-skip-tls-verify=true
kubectl config set-credentials scheduler --token=$KUBERNETES_TOKEN
kubectl config set-context default-context --cluster=k8s --user=scheduler
kubectl config use-context default-context
kubectl get cs

./k8s-bot --provider=$PROVIDER --token=$TOKEN --room=$ROOM
