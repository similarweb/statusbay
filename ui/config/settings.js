const mock = {
    webserver_endpoint: "http://localhost:8081",
    nomad_address: "http://127.0.0.1:4646"
}
const development = {
    webserver_endpoint: "http://localhost:8080",
}
const production = {
    webserver_endpoint: "https://statusbay-api.op-us-east-1.pe.svc.int.similarweb.io",
    nomad_address: "http://nomad-server-production.service.consul:4646",
}

var configuration = {}
switch(process.env.NODE_ENV){
    case 'mock':
        configuration = Object.assign({}, mock)
        break;
    case 'development':
            configuration = Object.assign(mock, development)
        break;
    case 'production':
        configuration = Object.assign(mock, development, production)
        break;
  }

module.exports = configuration

