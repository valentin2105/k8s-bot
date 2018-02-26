# k8s-bot

`k8s-bot` is a small Golang Bot that allow you to request your Kubernetes cluster via Hipchat channel.

[![Go Report Card](https://goreportcard.com/badge/github.com/valentin2105/k8s-bot)](https://goreportcard.com/report/github.com/valentin2105/k8s-bot)
## How install it ?
Install it quickly in your cluster using `Helm` with your Chat Application credentials :

```
git clone https://github.com/valentin2105/k8s-bot.git && cd k8s-bot/helm
helm install -n k8s-bot --namespace k8s-bot --set provider=Hipchat --set token=<token> --set room=<room> .
```

(You need to use Hipchat User's token and Room's ID like `3999999`)

## How to use it ?
```
!k - get cs

!k kube-system get deploy 

!k all get pod
```

<img src="https://i.imgur.com/9qNRiiT.png" width="463" height="236"> 

## How to build ?
```
go get github.com/tbruyelle/hipchat-go/hipchat
cd $GOPATH/src
git clone https://github.com/valentin2105/k8s-bot.git
go build
./k8s-bot -h
```

k8s-bot require `kubectl` and theses environment variables :
```
KUBERNETES_SERVICE_HOST=10.32.0.1	# By default on Kubernetes
KUBERNETES_SERVICE_PORT=443		# By default on Kubernetes
PROVIDER=Hipchat
TOKEN=<hipchat-user-token>
ROOM=<hipchat-room-id>
```
All of them are included on the Helm/Dockerfile stuffs.
