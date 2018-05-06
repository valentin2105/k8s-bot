# k8s-bot

[![Go Report Card](https://goreportcard.com/badge/github.com/valentin2105/k8s-bot)](https://goreportcard.com/report/github.com/valentin2105/k8s-bot)
[![Build Status](https://travis-ci.org/valentin2105/k8s-bot.svg?branch=master)](https://travis-ci.org/valentin2105/k8s-bot)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/dwyl/esta/issues)



`k8s-bot` is a small Golang Bot that allows you to interact with your Kubernetes cluster via Hipchat channel.

## How install it ?
You can install it in your cluster using `Helm` and your Chat Application credentials:

```
git clone https://github.com/valentin2105/k8s-bot.git && cd k8s-bot/helm

# Edit values.yaml

pod: k8s-bot
image: valentinnc/k8s-bot:latest
serviceAccount: k8s-bot
provider: Hipchat
room: "3999999" # Change Room ID
token: "thenah7oCaishei7een2it7iu9jucie5uphoa1ohg2ain2iengooShoo8Ee7eeNa" # Change Token

helm install -n k8s-bot --namespace k8s-bot .
```

You should generate a Hipchat User's token (`Account Settings -> API Access -> Generate token`)

## How to use it ?
```
!k - get cs

!k kube-system get deploy

!k all get pod
```

<img src="https://i.imgur.com/9qNRiiT.png" width="463" height="236">


`!k [ns] [verb] [resource]`

## How to build ?
```
go get github.com/tbruyelle/hipchat-go/hipchat
cd $GOPATH/src
git clone https://github.com/valentin2105/k8s-bot.git
go build
./k8s-bot -h
```

k8s-bot requires `kubectl` and these environment variables:
```
KUBERNETES_SERVICE_HOST=10.32.0.1	# By default on Kubernetes
KUBERNETES_SERVICE_PORT=443		# By default on Kubernetes
PROVIDER=Hipchat
TOKEN=<hipchat-user-token>
ROOM=<hipchat-room-id>
```
All of the above are included in the Helm chart and Dockerfile.

## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/valentin2105/k8s-bot/badge.svg?style=beer-square)](https://beerpay.io/valentin2105/k8s-bot)  [![Beerpay](https://beerpay.io/valentin2105/k8s-bot/make-wish.svg?style=flat-square)](https://beerpay.io/valentin2105/k8s-bot?focus=wish)