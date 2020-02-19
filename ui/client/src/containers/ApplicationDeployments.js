import React, { useMemo } from 'react';
import Box from '@material-ui/core/Box';
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
    distinct: true,
    name: appName,
  }), []);
  const title = `Application: ${appName}`;
  return (
    <PageContent>
      <Table hideNameFilter={true} filters={filters} onRowClick={onRowClick} title={title} />
    </PageContent>
  );
};

export default ApplicationDeployments;
