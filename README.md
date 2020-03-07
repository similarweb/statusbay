# StatusBay 

![Go](https://github.com/similarweb/statusbay/workflows/Go/badge.svg?event=push)
[![Coverage Status](https://coveralls.io/repos/github/similarweb/statusbay/badge.svg?branch=master)](https://coveralls.io/github/similarweb/statusbay?branch=master)

<img src="https://github.com/similarweb/statusbay/raw/master/docs/images/logo.png" width="400">

---

## Deployment visibility like a pro
Key features:
- Watch every step of K8S deployment.
- Get Slack reports on deployment progress.
- Out of the box integrations to measure your deployment quality.
- Deployed on k8s with [Helm][0].
- Easily extensible.
- Streamline the trouble shooting experience in K8S.

### What is StatusBay?
StatusBay is an open source tool that provides the missing visibility into the K8S deployment process. 
It does that by subscribing to K8S cluster(s), collecting all the relevant events from K8S and providing a step by step "zoom-in" into the deployment process.
The main goal is to ease the experience of troubleshooting and debugging services in K8S and provide confidence while making changes. 

StatusBay is designed to be dynamic and extensible, you can easily integrate with different metric providers to monitor the quality of the deployment over time. 

We've also created an API to provide an easy way to access the data and built a UI on top of it.

![Statusbay](/docs/images/statusbay.gif)

## Getting Started

1. The quickest way to get started with StatusBay is by using K8S. [Get started with StatusBay Helm Chart](https://github.com/similarweb/statusbay-helm).
2. Deploy your application. If you'd like to adopt all StatusBay features, see available configuration options [in this example](/docs/how-to-use.md).

[See DockerHub registry](https://hub.docker.com/r/similarweb/statusbay)

## Documentation & Guides

* [Developer Guide](/docs/developers/README.md): If you are interested in contributing, read the developer guide.
* [Working with Multiple Clusters](/docs/clusters/README.md): If you have multiple K8S clusters and you wish to have a unified deployment view, take a look at this guide.
* [Integrations](/docs/integrations.md): List of StatusBay supported integrations.
* [External Logging System](/docs/external-logs.md): Ship StatusBay logs to your centralized logging system.
* [Telemetry metrics](/docs/telemetry.md): StatusBay exposes metrics for you to pick up, see the telemetry read me to get started.

## How does it work?

StatusBay **watcher** subscribes to K8S cluster event stream and watches for resource changes (CREATE/UPDATE/DELETE).
Upon a change, such as new application deployment, it starts monitoring the progress of all the resource kinds (deployment, statefulset, daemonset, etc) associated with that deployment, notifies the relevant persona on success/failure/timeout and provides detailed report through the UI.

**Example Scenario**:

Someone has deployed an Nginx through Helm or Kubectl.
```bash
$ helm install {{NGINX_APP}} .

# OR

$ kubectl create deployment --image nginx my-nginx
```

The watcher will immediately start monitoring the deployment named `my-nginx` and report to the user using the notifications channels configured (slack, email, etc).

The following annotations can be attached to deployment to configure the different features StatusBay has to offer.

#### [Read more on StatusBay deployment configuration annotations](/docs/how-to-use.md)


## Built With

* [GO](https://golang.org/).
* [K8S Client Library](https://github.com/kubernetes/client-go/).

## Contributing

Thank you for your interest in contributing! Please refer to [CONTRIBUTING.md](./CONTRIBUTING.md) for guidance.

[0]: https://github.com/similarweb/statusbay-helm
[1]: https://github.com/similarweb/statusbay/wiki