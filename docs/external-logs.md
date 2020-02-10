# External Logging System

The StatusBay application can expose the application logs to external logging system (For example elasticsearch stack).


## How to setup?

StatusBay can ship the application logs to external logging system by using Gelf protocol. 


### Using helm deployment?

Go to [Helm Deployment](https://github.com/similarweb/statusbay-helm) and configure the `api.application.log.gelf_address` value.
