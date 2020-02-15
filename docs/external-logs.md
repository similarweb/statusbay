## External Logging System

StatusBay application can ship the application logs to external logging system (for example ELK stack).

<hr>

### Configuration

In order to ship StatusBay's application logs to an external logging system, configure the following API settings:

```bash
api:
  application:    
    log:
      level: INFO
      gelf_address: 127.0.0.1
```


### Using Helm?

Go to the [helm chart configuration](https://github.com/similarweb/statusbay-helm/blob/master/values.yaml) and set `api.application.log.gelf_address` value.
