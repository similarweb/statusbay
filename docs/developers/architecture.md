# Architecture
![Architecture Diagram](../images/architecture.png)

## UI 
- Written in [React](https://reactjs.org/).
- Using [React Material](https://material-ui.com/) for UI components.
- Every API request goes to UI backend.

## UI Backend
- Written in Node.js.
- This component is in charge of the communication between the UI and the API.

## API
- Written in [Golang](https://golang.org/).
- The API component communicates with the database **(Read-Only access)**.
- See the [list of available API endpoints](api-endpoints.md).

## Watcher
- Written in [Golang](https://golang.org/).
- Subscribes to resource changes (CREATE/UPDATE/DELETE) in a single K8S cluster, collects the information and saves the results to the database. 


