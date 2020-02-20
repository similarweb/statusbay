import React, { useMemo } from 'react';
import { useHistory } from 'react-router-dom';
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
  const filters = useMemo(() => {return {
    distinct: false
  }}, []);
  return (
    <PageContent>
      <Table onRowClick={onRowClick} filters={filters} title="Applications" />
    </PageContent>
  );
};

export default Applications;
