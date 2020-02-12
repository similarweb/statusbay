const crypto = require("crypto");

function generate() {
    const id = crypto.randomBytes(16).toString("hex");
    return id;
}

const logs = [{
    title: 'Readiness probe failed: dial tcp 10.60.179.55:5555: i/o timeout',
    time: 1576584132,
    content: 'The pod not pass the hatcheck: <ul><li>Indicates whether the Container is ready to service requests. If the readiness probe fails, the endpoints controller removes the Pod’s IP address from the endpoints of all Services that match the Pod. </li><li><a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate" target="_blank">Read more on Readiness details</a></li></ul>',
    error: true,
},
    {
        title: 'Liveness probe failed: dial tcp 10.60.179.55:5555: i/o timeout',
        time: 1576584133,
        error: false,
    },
    {
        title: 'Container stacky-celery failed liveness probe, will be restarted',
        time: 1576584139,
        error: false,
    },
    {
        title: 'Container image "placr/flower:0.9.1" already present on machine',
        time: 1576584169,
        error: false,
    },
    {
        title: 'Readiness probe failed: dial tcp 10.60.179.55:5555: i/o timeout',
        time: 1576584132,
        content: 'The pod not pass the hatcheck: <ul><li>Indicates whether the Container is ready to service requests. If the readiness probe fails, the endpoints controller removes the Pod’s IP address from the endpoints of all Services that match the Pod. </li><li><a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate" target="_blank">Read more on Readiness details</a></li></ul>',
        error: true,
    },
    {
        title: 'Liveness probe failed: dial tcp 10.60.179.55:5555: i/o timeout',
        time: 1576584133,
        error: false,
    },
    {
        title: 'Container stacky-celery failed liveness probe, will be restarted',
        time: 1576584139,
        error: false,
    },
    {
        title: 'Container image "placr/flower:0.9.1" already present on machine',
        time: 1576584169,
        error: false,
    },
    {
        title: 'Readiness probe failed: dial tcp 10.60.179.55:5555: i/o timeout',
        time: 1576584132,
        content: 'The pod not pass the hatcheck: <ul><li>Indicates whether the Container is ready to service requests. If the readiness probe fails, the endpoints controller removes the Pod’s IP address from the endpoints of all Services that match the Pod. </li><li><a href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate" target="_blank">Read more on Readiness details</a></li></ul>',
        error: true,
    },
    {
        title: 'Liveness probe failed: dial tcp 10.60.179.55:5555: i/o timeout',
        time: 1576584133,
        error: false,
    },
    {
        title: 'Container stacky-celery failed liveness probe, will be restarted',
        time: 1576584139,
        error: false,
    },
    {
        title: 'Container image "placr/flower:0.9.1" already present on machine',
        time: 1576584169,
        error: false,
    }];

const points = [
    [
        1577734320000,
        1,
    ],
    [
        1577734350000,
        1,
    ],
    [
        1577734380000,
        2,
    ]
];
module.exports = {
    getAll() {
        if (points.length <= 20) {
            const last = points[points.length - 1][0];
            points.push([
                last + 30000,
                Math.floor(Math.random() * 10) + 1,
            ]);
        }
        return {
            data: {
                name: 'hare-deployment-5',
                status: 'running',
                kinds: [
                    {
                        name: 'overwatch-api',
                        type: 'deployment',
                        stats: {
                            desired: Math.floor(Math.random() * 10) + 1,
                            current: Math.floor(Math.random() * 10) + 1,
                            updated: Math.floor(Math.random() * 10) + 1,
                            ready: Math.floor(Math.random() * 10) + 1,
                            available: Math.floor(Math.random() * 10) + 1,
                            unavailable: Math.floor(Math.random() * 10) + 1,
                        },
                        replicaSets: [],
                        deploymentEvents: [],
                        podEvents: [
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            }
                        ],
                        metrics: {
                            deploymentTime: 1577734380000,
                            http5xx4xx: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                            http2xx: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                            avgResponseTime: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                            memory: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                        },
                        alerts: {
                            deploymentTime: 1578127835,
                            providers: [
                                {
                                    provider: 'Statuscake',
                                    data: [
                                        {
                                            name: 'statuscake1',
                                            link: '',
                                            data: [
                                                {
                                                    from: 1578127665,
                                                    to: 1578127835,
                                                },
                                                {
                                                    from: 1578127935,
                                                    to: 1578128000,
                                                },
                                            ],
                                        },
                                        {
                                            name: 'statuscake2',
                                            link: '',
                                            data: [
                                                {
                                                    from: 1578127665,
                                                    to: 1578127835,
                                                },
                                            ],
                                        },
                                        {
                                            name: 'statuscake3',
                                            link: '',
                                            data: [
                                                {
                                                    from: 1578127665,
                                                    to: 1578127835,
                                                },
                                            ],
                                        },
                                    ],
                                }
                            ]
                        }
                    },
                    {
                        name: 'overwatch-api-1',
                        type: 'daemonSet',
                        stats: {
                            desired: Math.floor(Math.random() * 10) + 1,
                            current: Math.floor(Math.random() * 10) + 1,
                            updated: Math.floor(Math.random() * 10) + 1,
                            ready: Math.floor(Math.random() * 10) + 1,
                            available: Math.floor(Math.random() * 10) + 1,
                            unavailable: Math.floor(Math.random() * 10) + 1,
                        },
                        replicaSets: [],
                        deploymentEvents: [],
                        podEvents: [
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-60f45e89',
                                status: 'running',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-eb92dc03',
                                status: 'successful',
                                logs: logs.slice(0, 10)
                            },
                            {
                                name: 'web-pro-task-0982abac',
                                status: 'failed',
                                logs: logs.slice(0, 10)
                            }
                        ],
                        metrics: {
                            deploymentTime: 1577734380000,
                            http5xx4xx: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                            http2xx: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                            avgResponseTime: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                            memory: [
                                {
                                    name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
                                    points
                                },
                                {
                                    name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
                                    points
                                }
                            ],
                        },
                        alerts: {
                            deploymentTime: 1578127835,
                            providers: [
                                {
                                    provider: 'Statuscake',
                                    data: [
                                        {
                                            name: 'statuscake1',
                                            link: '',
                                            data: [
                                                {
                                                    from: 1578127665,
                                                    to: 1578127835,
                                                },
                                                {
                                                    from: 1578127935,
                                                    to: 1578128000,
                                                },
                                            ],
                                        },
                                        {
                                            name: 'statuscake2',
                                            link: '',
                                            data: [
                                                {
                                                    from: 1578127665,
                                                    to: 1578127835,
                                                },
                                            ],
                                        },
                                        {
                                            name: 'statuscake3',
                                            link: '',
                                            data: [
                                                {
                                                    from: 1578127665,
                                                    to: 1578127835,
                                                },
                                            ],
                                        },
                                    ],
                                }
                            ]
                        }
                    }
                ]
            }
        }
    }
}
