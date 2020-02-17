import React from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import querystring from 'query-string';
import {
  useLocation,
  useHistory,
  useParams,
} from 'react-router-dom';
import LinearProgress from '@material-ui/core/LinearProgress';
import PageTitle from '../components/Layout/PageTitle';
import PageContent from '../components/Layout/PageContent';
import ReplicasStats from '../DataComponents/ReplicasStats';
import PodEvents from '../DataComponents/PodEvents';
import Metrics from '../DataComponents/Metrics';
import Alerts from '../DataComponents/Alerts';
import DeploymentStatus from '../DataComponents/DeploymentStatus';
import Kinds from '../DataComponents/Kinds';
import DeploymentEvents from '../DataComponents/DeploymentEvents';
import {
  DeploymentDetailsContextProvider,
} from '../context/DeploymentDetailsContext';
import * as moment from 'moment';

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
      {
        ({ data, loading }) => (loading ? <Box m={2} flexGrow={1} justifyContent="space-around" display="flex" flexDirection="column"><LinearProgress /></Box> : (
          <PageContent>
            <Box mt={3} mb={3}>
              <Grid container spacing={2} justify="space-between" alignContent="center">
                <Grid item xs={12} xl={6}>
                  <PageTitle>
                    {data.name}
                    <Typography variant="body2">Namespace: {data.namespace}</Typography>
                    <Typography variant="body2">Cluster: {data.cluster}</Typography>
                    <Typography variant="body2">{moment.unix(data.time).utc().format('DD/MM/YYYY HH:MM:ss')}</Typography>
                  </PageTitle>
                </Grid>
                <Grid container xs={12} xl={6} alignContent="center" direction="row-reverse">
                  <Grid item>
                    <DeploymentStatus />
                  </Grid>
                </Grid>
              </Grid>
            </Box>
            <Kinds selectedTab={parseInt(tab)} onTabChange={handleTabChange} />
            <Box mt={2} mb={2}>
              <ReplicasStats kindIndex={parseInt(tab)} />
            </Box>
            <Box mt={2} mb={2}>
              <PodEvents kindIndex={parseInt(tab)} />
              {/* <Metrics kindIndex={parseInt(tab)} /> */}
              <DeploymentEvents kindIndex={parseInt(tab)} />
              {/* <Alerts kindIndex={parseInt(tab)} /> */}
            </Box>
          </PageContent>

        ))

      }
    </DeploymentDetailsContextProvider>
  );
};

export default DeploymentDetails;
