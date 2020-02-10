# Example Helm chart for Statusbay deployments

Example of multiple deployment for debugging Statusbay behavers

## Install

```
helm upgrade -i foo .
```

## Uninstall
```
helm delete foo
```

## Configuration
| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `daemonset.count` | number of daemonset deployments | `1` |
| `daemonset.name` | prefix of the daemonset name | `statusbay-daemonset` |
| `daemonset.image.repository` | container image repository | `nginx` |
| `daemonset.image.tag` | container image tag | `latest` |
| `daemonset.image.pullPolicy` | container image pull policy | `IfNotPresent` |
| `daemonset.resources` | the [resources] to allocate for a pod | undefined |
| `daemonset.livenessProbe` | liveness health check | `HTTP 80 /` |
| `daemonset.readinessProbe` | readiness health check | `HTTP 80 /` |
| `daemonset.annotations` | the statusbay annotations to set | `list of annotations` |
| `deployment.count` | number of deployments to simulate | `1` |
| `deployment.replicas` | number of replicas in each deployment | `3` |
| `deployment.progressDeadlineSeconds` | maximum deployment time | `300` |
| `deployment.name` | prefix of the deployment name | `statusbay-deployment` |
| `deployment.image.repository` | container image repository | `nginx` |
| `deployment.image.tag` | container image tag | `latest` |
| `deployment.image.pullPolicy` | container image pull policy | `IfNotPresent` |
| `deployment.resources` | the [resources] to allocate for a pod | undefined |
| `deployment.livenessProbe` | readiness health check | `HTTP 80 /` |
| `deployment.readinessProbe` | readiness health check | `HTTP 80 /` |