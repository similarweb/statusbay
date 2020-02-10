# How To Use

### After the initial setup those are the different features when writing an application deployment: 

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

**Annotation (Optional):** 

`statusbay.io/application-name`

**When to use:**

Use when more then one resource should be grouped under the same application.

**For example:**

When deploying a service using a helm chart with multiple resources (Deployment, StatefulSet, Service, etc) 

---
**Annotation (Optional):** 

`statusbay.io/progress-deadline-seconds` 

**When to use:**

Use to override the default timeout of a deployment before its declared failed as failed. 

**For example (Optional):**

You have a resource that takes 5 minutes to warm up and become ready while the default progress dead line is 1 minute. 

---
**Annotation (Optional):** 

`statusbay.io/report-slack-channels`


**When to use:**

List of Slack (comma separated) channels for StatusBay notifications.
> Note: to get slack notification from StatusBay you must invite the given bot the to destination Slack channel 

**For example:**

When you deploy an application and you want to notify 3 different channels.

---
**Annotation (Optional):** 

`statusbay.io/alerts-pingdom` -

**When to use:**

List of Pingdom checks that are associated with your deployment.

---
**Annotation (Optional):** 

`statusbay.io/alerts-statuscake`

**When to use:**

List of Statuscake checks that are associated with your deployment.

---
**Annotation (Optional):** 

`statusbay.io/metrics-datadog-{custom-metric-name}`

**When to use:**

List of Datadog queries that are associated with your deployment.

---
**Annotation (Optional):** 

`statusbay.io/metrics-prometuse-errors-title-{custom-metric-name}`

**When to use:**

List of Prometuse queries that are associated with your deployment.
