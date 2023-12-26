# Gorilla

Go backend running gin gonic api server.

## Getting started

### install docker desktop if not installed yet
```bash
brew install --cask docker
```

### Spining up the local cluster

Open docker for desktop and enable the kubernetes engine. To do so:

1. Go to Settings -> Kubernetes -> enable Kubernetes ✅
2. Restart docker desktop

### Setting the kubectl context right

If you have already installed kubectl and it is pointing to some other environment, such as minikube or a EKS cluster, ensure you change the context so that kubectl is pointing to docker-desktop:

```bash
 kubectl config get-contexts
 kubectl config use-context docker-desktop
```



## Make targets
Here is a list of all available make targets:

- `generate`: Generate golang code
- `build-app`: Build docker image
- `deploy-all`: Deploy all apps to kubernetes
- `port-forward`: Port forward to api server at port 8080
