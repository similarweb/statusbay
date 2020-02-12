## How To Use

### Deployment Example
After you've installed and configured StatusBay, lets take a look at this deployment example:

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: statusbay-example
  labels:
    app.kubernetes.io/name: statusbay-example
annotations:
    statusbay.io/application-name: "statusbay-application"
    statusbay.io/progress-deadline-seconds: "60"
    statusbay.io/report-slack-channels: "#foo-channel"
    statusbay.io/report-deploy-by: similarweb@similarweb.com
    statusbay.io/alerts-pingdom: nginx,us-east-1
    statusbay.io/metrics-datadog-2xx: sum:web.http.2xx{*}
    statusbay.io/metrics-datadog-5xx: sum:web.http.5xx{*}
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: statusbay-example
  template:
    metadata:
      labels:
        app.kubernetes.io/name: statusbay-example
    spec:
      containers:
        - name: nginx
          image: nginx:latest
          imagePullPolicy: IfNotPresent
```


### Available Annotations
| Name | Description | Required | Example | 
| ---- | ----------- | -------- | ------- | 
| statusbay.io/application-name | Used to group different kinds under one logical StatusBay unit | No | `statusbay.io/application-name: my-amazing-app` |
| statusbay.io/progress-deadline-seconds | Use this to override the default deployment timeout before it's marked as failed | No | `statusbay.io/progress-deadline-seconds: 600` |
| statusbay.io/report-slack-channels | Comma separated list of Slack channels to send deployment notifications to. Note: the bot must be present in the channels | No | `statusbay.io/report-slack-channels: #deployments,#devops` |
| statusbay.io/alerts-pingdom | Comma separated Pingdom tags associated with this deployment | No | `statusbay.io/alerts-pingdom: nginx,us-east-1` |
| statusbay.io/alerts-statuscake | Comma separated StatusCake tags associated with this deployment | No | `statusbay.io/alerts-statuscake: nginx,us-east-1` |
| statusbay.io/metrics-datadog-{custom-metric-name} | Datadog metric associated with your deployment | No | `statusbay.io/metrics-datadog-2xx: sum:nginx.2xx{environment:production}` |
| statusbay.io/metrics-prometheus-{custom-metric-name} | Prometheus metric associated with your deployment | No | `statusbay.io/metrics-prometheus-5xx: prometheus_http_requests_total{code="200"}` |

