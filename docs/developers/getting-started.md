# Getting started
This document describes how to setup your local development environment for StatusBay development.

## Preparation

Make sure the following software is installed:

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

## Run the MySQL
```
$ docker run -p 3306:3306 -e MYSQL_DATABASE=statusbay -e MYSQL_ROOT_PASSWORD=1234 -d mysql:5.7
```

# Run the watcher
This command will run StatusBay watcher application.

Please refer to [configuration example](/examples/configuration/kubernetes.yaml) file.

```
$ go run main.go -config ./examples/configuration/kubernetes.yaml -mode kubernetes -kubeconfig ~/.kube/config
```

# Run the API server
This command will run the API server. 

Please refer to [configuration example](/examples/configuration/api.yaml) file.

```
$ go run main.go -config ./examples/configuration/api.yaml -mode api
```

# Test deployment example

See example of [Helm application](/examples/apply/README.md)
