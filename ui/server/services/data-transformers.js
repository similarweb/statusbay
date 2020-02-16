const kinds = {
  Deployment: 'deployment',
  DaemonSet: 'daemonSet',
}

const createDeploymentData = ([name, rawData]) => {
  return {
    name: name,
    type: kinds.Deployment,
    stats: transformers.status(rawData),
    // replicaSets: [],
    deploymentEvents: transformers.deploymentEvents(rawData),
    podEvents: transformers.podEvents(rawData),
    // metrics: {
    //   deploymentTime: 1577734380000,
    //   http5xx4xx: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    //   http2xx: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    //   avgResponseTime: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    //   memory: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    // },
    // alerts: {
    //   deploymentTime: 1578127835,
    //   providers: [
    //     {
    //       provider: 'Statuscake',
    //       data: [
    //         {
    //           name: 'statuscake1',
    //           link: '',
    //           data: [
    //             {
    //               from: 1578127665,
    //               to: 1578127835,
    //             },
    //             {
    //               from: 1578127935,
    //               to: 1578128000,
    //             },
    //           ],
    //         },
    //         {
    //           name: 'statuscake2',
    //           link: '',
    //           data: [
    //             {
    //               from: 1578127665,
    //               to: 1578127835,
    //             },
    //           ],
    //         },
    //         {
    //           name: 'statuscake3',
    //           link: '',
    //           data: [
    //             {
    //               from: 1578127665,
    //               to: 1578127835,
    //             },
    //           ],
    //         },
    //       ],
    //     }
    //   ]
    // }
  }
}
const createDaemonSetData = ([name, rawData]) => {
  return {
    name: name,
    type: kinds.DaemonSet,
    stats: transformers.status(rawData),
    // replicaSets: [],
    deploymentEvents: transformers.daemonSetEvents(rawData),
    podEvents: transformers.podEvents(rawData),
    // metrics: {
    //   deploymentTime: 1577734380000,
    //   http5xx4xx: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    //   http2xx: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    //   avgResponseTime: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    //   memory: [
    //     {
    //       name: 'sum:sw.haproxy.backend.response.5xx{dc:op-us-east-1,service:be_pe',
    //       points
    //     },
    //     {
    //       name: 'sum:sw.haproxy.backend.response.4xx{dc:op-us-east-1,service:be_p',
    //       points
    //     }
    //   ],
    // },
    // alerts: {
    //   deploymentTime: 1578127835,
    //   providers: [
    //     {
    //       provider: 'Statuscake',
    //       data: [
    //         {
    //           name: 'statuscake1',
    //           link: '',
    //           data: [
    //             {
    //               from: 1578127665,
    //               to: 1578127835,
    //             },
    //             {
    //               from: 1578127935,
    //               to: 1578128000,
    //             },
    //           ],
    //         },
    //         {
    //           name: 'statuscake2',
    //           link: '',
    //           data: [
    //             {
    //               from: 1578127665,
    //               to: 1578127835,
    //             },
    //           ],
    //         },
    //         {
    //           name: 'statuscake3',
    //           link: '',
    //           data: [
    //             {
    //               from: 1578127665,
    //               to: 1578127835,
    //             },
    //           ],
    //         },
    //       ],
    //     }
    //   ]
    // }
  }
}

const convertDeploymentDetailsData = (data) => {
  return {
    name: data.Name,
    status: data.Status,
    time: data.Time,
    namespace: data.Namespace,
    cluster: data.Cluster,
    kinds: [
      ...Object.entries(data.Details.Resources.Deployments).map(createDeploymentData),
      ...Object.entries(data.Details.Resources.Daemonsets).map(createDaemonSetData),
    ]
  }
};
const transformers = {
  status: (rawData) => {
    return {
      desired: rawData.Status.ObservedGeneration,
      current: rawData.Status.Replicas,
      updated: rawData.Status.UpdatedReplicas,
      ready: rawData.Status.ReadyReplicas,
      available: rawData.Status.AvailableReplicas,
      unavailable: rawData.Status.UnavailableReplicas,
    }
  },
  deploymentEvents: (rawData) => {
    return rawData.Events.map(event => {
      return {
        title: event.Message,
        time: event.Time,
        content: event.MarkDescriptions && event.MarkDescriptions.length > 0 && event.MarkDescriptions[0],
        error: event.MarkDescriptions ? event.MarkDescriptions.length > 0 : false,
      }
    })
  },
  daemonSetEvents: (rawData) => {
    return rawData.Events.map(event => {
      return {
        title: event.Message,
        time: event.Time,
        content: event.MarkDescriptions && event.MarkDescriptions.length > 0 && event.MarkDescriptions[0],
        error: event.MarkDescriptions ? event.MarkDescriptions.length > 0 : false,
      }
    })
  },
  podEvents: (rawData) => {
    return Object.entries(rawData.Pods).map(([name, pod]) => {
      return {
        name,
        status: pod.Phase.toLowerCase(),
        time: pod.CreationTimestamp,
        logs: pod.Events.map(event => {
          return {
            title: event.Message,
            time: event.Time,
            content: event.MarkDescriptions && event.MarkDescriptions.length > 0 && event.MarkDescriptions[0],
            error: event.MarkDescriptions ? event.MarkDescriptions.length > 0 : false,
          }
        })
      }
    })
  }
}

module.exports = {convertDeploymentDetailsData};
