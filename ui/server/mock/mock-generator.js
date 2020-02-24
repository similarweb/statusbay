const faker = require('faker');
const moment = require('moment');
const clusters = require('./clusters.mock').getAll().map(x => x.Name);
const nameSpaces = require('./name-spaces.mock').getAll().map(x => x.Name);
const users = require('./users.mock').getAll().map(x => x.Name);
/*
In the server folder
node -e 'console.log(require("./mock/mock-generator.js").getAll(100))' > mock/mock-data.json
node -e 'console.log(JSON.stringify(require("./mock/mock-generator.js").getDetails(), null, 4))' > mock/mock-details.json
*/

const appNames = [...Array(10)].map(() => faker.lorem.slug());
module.exports = {
    getAll(number){
        return [...Array(number)].map(()=>{
            return {
                "Name": faker.random.arrayElement(appNames),
                "Status": faker.random.arrayElement(["successful","failed","running", "timeout", "deleted"]),
                "Cluster": faker.random.arrayElement(clusters),
                "Namespace": faker.random.arrayElement(nameSpaces),
                "DeployBy": faker.random.arrayElement(users),
                "Time": (new Date(faker.date.between(moment(),moment().subtract(1, 'weeks')))).getTime(),
                "ApplyId": faker.random.number().toString(),
            }
        });
    },
    getDetails() {
        return {
            "Name":"cormorant-deployment-7",
            "Cluster":"tokyo",
            "Namespace":"heavy",
            "Status":"running",
            "Time":"",
            "Details":{
                "Resources":{
                    "Deployments":{
                        "deployment1":{
                            "MetaData":{
                                "Name":"deployment1",
                                "Namespace":"default",
                                "ClusterName":"",
                                "Labels":{
                                    "app.kubernetes.io/managed-by":"me",
                                    "app.kubernetes.io/name":"deployment1"
                                },
                                "DesiredState":1,
                                "Alerts":null,
                                "Metrics":null
                            },
                            "DeploymentEvents":[
                                {
                                    "Message":"Scaled up replica set deployment1-9ff6b5676 to 1",
                                    "Time":1580032133000000000,
                                    "Action":"",
                                    "ReportingController":"",
                                    "MarkDescriptions":[

                                    ]
                                }
                            ],
                            "Metrics":null,
                            "Pods":{
                                "deployment1-9ff6b5676-c7nml":{
                                    "Phase":"Running",
                                    "CreationTimestamp":"0001-01-01T00:00:00Z",
                                    "Events":[
                                        {
                                            "Message":"Successfully assigned default/deployment1-9ff6b5676-c7nml to minikube",
                                            "Time":1580032133000000000,
                                            "Action":"Binding",
                                            "ReportingController":"default-scheduler",
                                            "MarkDescriptions":[

                                            ]
                                        },
                                        {
                                            "Message":"Container image \"nginx:latest\" already present on machine",
                                            "Time":1580032135000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        },
                                        {
                                            "Message":"Created container nginx",
                                            "Time":1580032135000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        },
                                        {
                                            "Message":"Started container nginx",
                                            "Time":1580032135000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        }
                                    ]
                                }
                            },
                            "Replicaset":{
                                "deployment1-9ff6b5676":{
                                    "Events":[
                                        {
                                            "Message":"Created pod: deployment1-9ff6b5676-c7nml",
                                            "Time":1580032133000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        }
                                    ]
                                }
                            },
                            "Status":{
                                "ObservedGeneration":1,
                                "Replicas":1,
                                "UpdatedReplicas":1,
                                "ReadyReplicas":1,
                                "AvailableReplicas":1,
                                "UnavailableReplicas":0
                            }
                        },
                        "elad":{
                            "MetaData":{
                                "Name":"root",
                                "Namespace":"default",
                                "ClusterName":"",
                                "Labels":{
                                    "app.kubernetes.io/managed-by":"me",
                                    "app.kubernetes.io/name":"root"
                                },
                                "DesiredState":1,
                                "Alerts":null,
                                "Metrics":null
                            },
                            "DeploymentEvents":[
                                {
                                    "Message":"Scaled up replica set root-76d67fd644 to 1",
                                    "Time":1580032133000000000,
                                    "Action":"",
                                    "ReportingController":"",
                                    "MarkDescriptions":[

                                    ]
                                }
                            ],
                            "Metrics":null,
                            "Pods":{
                                "elad-76d67fd644-cxnl7":{
                                    "Phase":"Running",
                                    "CreationTimestamp":"0001-01-01T00:00:00Z",
                                    "Events":[
                                        {
                                            "Message":"Successfully assigned default/root-76d67fd644-cxnl7 to minikube",
                                            "Time":1580032133000000000,
                                            "Action":"Binding",
                                            "ReportingController":"default-scheduler",
                                            "MarkDescriptions":[

                                            ]
                                        },
                                        {
                                            "Message":"Container image \"nginx:latest\" already present on machine",
                                            "Time":1580032135000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        },
                                        {
                                            "Message":"Created container nginx",
                                            "Time":1580032135000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        },
                                        {
                                            "Message":"Started container nginx",
                                            "Time":1580032136000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        }
                                    ]
                                }
                            },
                            "Replicaset":{
                                "root-76d67fd644":{
                                    "Events":[
                                        {
                                            "Message":"Created pod: root-76d67fd644-cxnl7",
                                            "Time":1580032133000000000,
                                            "Action":"",
                                            "ReportingController":"",
                                            "MarkDescriptions":[

                                            ]
                                        }
                                    ]
                                }
                            },
                            "Status":{
                                "ObservedGeneration":1,
                                "Replicas":1,
                                "UpdatedReplicas":1,
                                "ReadyReplicas":1,
                                "AvailableReplicas":1,
                                "UnavailableReplicas":0
                            }
                        }
                    },
                    "Daemonsets":{

                    },
                    "Statefulsets": {

                    }
                }
            }
        }
    }
};
