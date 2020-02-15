import React from 'react';
import Box from '@material-ui/core/Box';
import querystring from 'query-string';
import {
  useLocation,
  useHistory,
  useParams,
} from 'react-router-dom';
import AppBar from '@material-ui/core/AppBar';
import PageTitle from '../components/Layout/PageTitle';
import PageContent from '../components/Layout/PageContent';
import ReplicasStats from '../DataComponents/ReplicasStats';
import PodEvents from '../DataComponents/PodEvents';
import Metrics from '../DataComponents/Metrics';
import Alerts from '../DataComponents/Alerts';
import DeploymentStatus from '../DataComponents/DeploymentStatus';
import Kinds from '../DataComponents/Kinds';
import DeploymentEvents from '../DataComponents/DeploymentEvents';
import { DeploymentDetailsContextProvider } from '../context/DeploymentDetailsContext';

const DeploymentDetails = () => {
  const location = useLocation();
  const history = useHistory();
  const { tab = '0' } = querystring.parse(location.search);
  const { deploymentId } = useParams();
  const handleTabChange = (event, newValue) => {
    history.push({
      pathname: location.pathname,
      search: `?${new URLSearchParams({
        tab: newValue,
      })}`,
    });
  };
  return (
    <DeploymentDetailsContextProvider id={`${deploymentId}`}>
      <PageContent>
        <Box m={3} display="flex" alignItems="center" justifyContent="space-between">
          <PageTitle>
            Deployment details:
            {deploymentId}
          </PageTitle>
          <DeploymentStatus />
        </Box>
        <AppBar position="static" color="primary">
          <Kinds selectedTab={parseInt(tab)} onTabChange={handleTabChange} />
        </AppBar>
        <Box m={2}>
          <ReplicasStats kindIndex={parseInt(tab)} />
        </Box>
        <Box m={2}>
          <PodEvents kindIndex={parseInt(tab)} />
          {/* <Metrics kindIndex={parseInt(tab)} /> */}
          <DeploymentEvents kindIndex={parseInt(tab)} />
          {/* <Alerts kindIndex={parseInt(tab)} /> */}
        </Box>
      </PageContent>
    </DeploymentDetailsContextProvider>
  );
};

export default DeploymentDetails;
