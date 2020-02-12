# Getting started
This document describes how to setup your local development environment for StatusBay.

## Preparation

Make sure the following tools are installed:

* Docker
* Golang 1.12.0+ ([installation manual](https://golang.org/dl/))
* Node.js 12+ and npm 6+ ([installation with nvm](https://github.com/creationix/nvm#usage))
* MySQL 5.7
* Minikube ([installation manual](https://kubernetes.io/docs/tasks/tools/install-minikube/))
* (Optional) Helm ([installation manual](https://helm.sh/docs/intro/install/))

Fork StatusBay project

## Run Minikube

If you want to run the watcher application, you should have a running K8S environment. 
Minikube is one the options to achieve that.

```
$ minikube start
```

## Run MySQL
```
$ docker run -p 3306:3306 -e MYSQL_DATABASE=statusbay -e MYSQL_ROOT_PASSWORD=1234 -d mysql:5.7
```

# Run StatusBay watcher
This command will run StatusBay watcher.

Please refer to [configuration example](/examples/configuration/kubernetes.yaml) file to see additional configurations.

```
$ go run main.go -config ./examples/configuration/kubernetes.yaml -mode kubernetes -kubeconfig ~/.kube/config


# The reason we pass -kubeconfig in the command above is in order to define the cluster we wish the watcher to subscribe to it's event stream.
```

# Run the API server
This command will run the API server. 

Please refer to [configuration example](/examples/configuration/api.yaml) file to see additional configurations..

```
$ go run main.go -config ./examples/configuration/api.yaml -mode api
```

# Deployment example via Helm

Use this [Helm application](/examples/apply/README.md) in order to test your deployment.

