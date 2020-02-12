import API from './index';

export const getMetaData = async () => {
  const { data: { allClusters, allNamespaces } } = await API('/api/metadata');
  return {
    clusters: allClusters,
    namespaces: allNamespaces,
    statuses: ['successful', 'failed', 'running', 'timeout', 'deleted'],
  };
};

export const transformTableData = (results, totalCount) => ({
  totalCount,
  rows: results.map(({
    Name,
    Status,
    Cluster,
    Namespace,
    DeployBy,
    Time,
  }) => ({
    name: Name,
    status: Status,
    cluster: Cluster,
    namespace: Namespace,
    deployBy: DeployBy,
    time: Time,
  })),
});
