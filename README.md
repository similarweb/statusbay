# StatusBay 

![Go](https://github.com/similarweb/statusbay/workflows/Go/badge.svg?event=push)


## Deployment visibility like a pro
> - Watch every step of K8S deployment.
> - Get slack reports on deployment progress.
> - Out of the box integration to measure your deployment quality.
> - Run on K8S (with [Helm][0]).
> - Easily extensible.
> - Streamline the trouble shooting experience in k8s.  

### What is StatusBay?
StatusBay is an open source tools that provides visibility into the K8S deployment process. 
It does it by directly subscribing a k8s cluster. 
The watcher collects all the relevant events from k8s internals and provides a step by step "zoom-in" into the deployment process of multiple resources. 
The main goal is to ease the experience of troubleshooting and debugging a new/existing service in k8s. 

StatusBay is designed to be extensible and also provides a way to integrate with different metric providers to monitor the quality of the deployment overtime. 

All this info is served by an API and one central dashboard. 

![Statusbay](/docs/images/statusbay.gif)

## Getting Started

1. The quickest way to get started with StatusBay is K8S. To deploy StatusBay on K8S, [see Helm chart deployments][0]
2. Deploy you application. (Application example: [Deployment](/docs/how-to-use.md)). 

[See DockerHub registry](https://hub.docker.com/r/similarweb/statusbay)

## Documentation

* [Developer Guide](/docs/developers/README.md): If you are interested in contributing, read the developer guide.
* [Support Multiple Clusters](/docs/clusters/README.md): Watch on multiple clusters.
* [Providers Integrations](/docs/how-to-use.md): List of K8S deployment integrations.
* [External Logging System](/docs/external-logs.md): Explain how to ship StatusBay application logs to logging system.
* [Telemetry metrics](/docs/telemetry.md): Expose StatusBay metrics

## How it works?

StatusBay registers and watches to resource changes (CREATE/UPDATE/DELETE) in K8S clusters. 
This component is a service named `watcher`.
Upon a change such as a new application deployment it starts monitoring the progress of all the resource kinds (deployment, statefulset, daemonset, etc), notifies about success/failure/timeout and basically all the related events.

**For example:**

A developer has deployed a Nginx via Helm.
```
$ helm install {{NGINX_APP}} .
```

OR

Via Kubectl
```
$ kubectl create deployment --image nginx my-nginx
```

The watcher will immediately start monitoring a deployment named `my-nginx` and report to all consumers (slack, api, etc).

The following annotations can be attached to a deployment in order to configure the different features StatusBay has to offer.

### [Read more about K8S deployment integrations](/docs/how-to-use.md)


## Built With

* [GO](https://golang.org/).
* [K8S Client Library](https://github.com/kubernetes/client-go/).

## Contributing

Thank you for your interest in contributing! Please refer to [CONTRIBUTING.md](./CONTRIBUTING.md) for guidance.

[0]: https://github.com/similarweb/statusbay-helm
[1]: https://github.com/similarweb/statusbay/wiki