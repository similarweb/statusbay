import API from './index';

export const getMetaData = async () => {
  const { data: { allClusters, allNamespaces, allStatuses } } = await API('/api/metadata');
  return {
    clusters: allClusters,
    namespaces: allNamespaces,
    statuses: allStatuses,
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
    ApplyID,
  }) => ({
    name: Name,
    status: Status,
    cluster: Cluster,
    namespace: Namespace,
    deployBy: DeployBy,
    time: Time,
    id: ApplyID,
  })),
});
