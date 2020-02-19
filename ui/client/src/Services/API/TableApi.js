import API from './index';

const defaultMetaData = {
  allClusters: [],
  allNamespaces: [],
  allStatuses: [],
}

export const getMetaData = async () => {
  let { data } = await API('/api/metadata');
  // use default meta data when server is down, or in case of error
  if (!data) {
    data = defaultMetaData
  }
  const { allClusters, allNamespaces, allStatuses } = data;
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
