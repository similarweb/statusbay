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
| `daemonset.createService: false` | Create service resource | false |
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
| `deployment.createService: false` | Create service resource | false |
| `deployment.name` | prefix of the deployment name | `statusbay-deployment` |
| `deployment.image.repository` | container image repository | `nginx` |
| `deployment.image.tag` | container image tag | `latest` |
| `deployment.image.pullPolicy` | container image pull policy | `IfNotPresent` |
| `deployment.resources` | the [resources] to allocate for a pod | undefined |
| `deployment.livenessProbe` | readiness health check | `HTTP 80 /` |
| `deployment.readinessProbe` | readiness health check | `HTTP 80 /` |
| `deployment.annotations` | the statusbay annotations to set | `list of annotations` |
| `statefulset.count` | number of statefulsets to simulate | `0` |
| `statefulset.replicas` | number of replicas in each statefulset | `1` |
| `statefulset.createService: false` | Create service resource | false |
| `statefulset.name` | prefix of the statefulset name | `statefulset` |
| `statefulset.image.repository` | container image repository | `nginx` |
| `statefulset.image.tag` | container image tag | `latest` |
| `statefulset.image.pullPolicy` | container image pull policy | `IfNotPresent` |
| `database.internal.resources` | The [resources] to allocate for container | undefined
| `statefulset.persistence.persistentVolumeClaim.accessMode` | The access mode of the volume | `ReadWriteOnce`
| `statefulset.persistence.persistentVolumeClaim.storageClass` | Specify the `storageClass` used to provision the volume. Or the default StorageClass will be used(the default). Set it to `-` to disable dynamic provisioning | `-`
| `statefulset.persistence.persistentVolumeClaim.size` | The size of the volume | `1Gi`
| `statefulset.annotations` | the statusbay annotations to set | `list of annotations` |
| `statefulset.livenessProbe` | readiness health check | `HTTP 80 /` |
| `statefulset.readinessProbe` | readiness health check | `HTTP 80 /` |
