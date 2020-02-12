export default [{
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
