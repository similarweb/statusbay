Example helm chart for Statusbay deployments

Example of multiple deployment for debugging Statusbay behavers

## Install the chart

```
helm upgrade -i foo .
```

## Uninstallation
```
helm delete foo
```

## Configuration
Parameter | Description | Default
--- | --- | ---
`daemonset.count` | Number of Daemonset deployments | `1`
`daemonset.name` | Prefix of Daemonset | `statusbay-daemonset`
`daemonset.image.repository` | container image repository | `nginx`
`daemonset.image.tag` | container image tag | `latest`
`daemonset.image.pullPolicy` | container image pull policy | `nginx`
`daemonset.resources` | The [resources] to allocate for container | undefined
`daemonset.livenessProbe` | The indicates whether the Container is running. | `Pass configuration check`
`daemonset.readinessProbe` | The indicates whether the Container is ready to service requests | `Pass configuration check`
`daemonset.annotations` The annotations used in ingress | `list of annotations`
`deployment.count` | Number of Deplotments deployments | `1`
`deployment.replicas` | Number of Deplotments replicas | `3`
`deployment.progressDeadlineSeconds` | Maximum deployment time | `300`
`deployment.name` | Prefix of Daemonset | `statusbay-deployment`
`deployment.image.repository` | container image repository | `nginx`
`deployment.image.tag` | container image tag | `latest`
`deployment.image.pullPolicy` | container image pull policy | `nginx`
`deployment.resources` | The [resources] to allocate for container | undefined
`deployment.livenessProbe` | The indicates whether the Container is running. | `Pass configuration check`
`deployment.readinessProbe` | The indicates whether the Container is ready to service requests | `Pass configuration check`