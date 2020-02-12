import React from 'react';
import Box from '@material-ui/core/Box';
import { useTranslation } from 'react-i18next';
import { useHistory } from 'react-router-dom';
import Table from '../components/Table/Table';
import PageContent from '../components/Layout/PageContent';
import PageTitle from '../components/Layout/PageTitle';

export default () => {
  const { t } = useTranslation();
  const history = useHistory();
  const onRowClick = (row) => () => {
    // redirect to details page
    history.push({
      pathname: `/application/${row.name}/${row.time}`,
    });
  };
  return (
    <PageContent>
      <Box m={3}>
        <PageTitle>
          {t('all deployments')}
        </PageTitle>
      </Box>
      <Table onRowClick={onRowClick} />
    </PageContent>
  );
};
