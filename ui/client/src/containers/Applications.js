import React, { useMemo } from 'react';
import Box from '@material-ui/core/Box';
import PageTitle from '../components/Layout/PageTitle';
import Table from '../components/Table/Table';
import PageContent from '../components/Layout/PageContent';
import { useHistory } from 'react-router-dom';

const Applications = () => {
  const history = useHistory();
  const onRowClick = (row) => () => {
    // redirect to application deployment page
    history.push({
      pathname: `/applications/${row.name}`,
    });
  };
  const filters = useMemo(() => {return {
    distinct: true
  }}, []);
  return (
    <PageContent>
      <Box m={3}>
        <PageTitle>
        Applications
        </PageTitle>
      </Box>
      <Box>
        <Table onRowClick={onRowClick} filters={filters} />
      </Box>
    </PageContent>
  );
};

export default Applications;
