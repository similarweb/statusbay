import React, { useMemo } from 'react';
import {
  useHistory,
  useParams,
} from 'react-router-dom';
import Table from '../components/Table/Table';
import PageContent from '../components/Layout/PageContent';

const ApplicationDeployments = () => {
  const { appName } = useParams();
  const history = useHistory();
  const onRowClick = (row) => () => {
    // redirect to deployment details page
    history.push({
      pathname: `/application/${row.id}`,
    });
  };
  const filters = useMemo(() => ({
    distinct: false,
    exactName: appName,
  }), []);
  return (
    <PageContent>
      <Table hideNameFilter={true} filters={filters} onRowClick={onRowClick} title={appName} hideDistinctFilter={true} />
    </PageContent>
  );
};

export default ApplicationDeployments;
