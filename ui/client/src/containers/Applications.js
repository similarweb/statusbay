import React, { useMemo } from 'react';
import Box from '@material-ui/core/Box';
import { useHistory } from 'react-router-dom';
import PageTitle from '../components/Layout/PageTitle';
import Table from '../components/Table/Table';
import PageContent from '../components/Layout/PageContent';

const Applications = () => {
  const history = useHistory();
  const onRowClick = (row) => (event) => {
    // redirect to application deployments page
    history.push({
      pathname: `/applications/${row.name}`,
    });
  };
  const filters = useMemo(() => ({
    distinct: false,
  }), []);
  return (
    <PageContent>
      <Table onRowClick={onRowClick} filters={filters} title="Applications" />
    </PageContent>
  );
};

export default Applications;
