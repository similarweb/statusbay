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
import * as moment from 'moment';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import Chip from '@material-ui/core/Chip';
import IconButton from '@material-ui/core/IconButton';
import makeStyles from '@material-ui/core/styles/makeStyles';
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
import Loader from '../components/Loader/Loader';
import ReplicaSetEvents from '../DataComponents/ReplicaSetEvents';
import ServiceSetEvents from '../DataComponents/ServiceSetEvents';
import NoData from '../components/Table/NoData';

const useStyles = makeStyles((theme) => ({
  chips: {
    '& > *': {
      margin: theme.spacing(0.5),
    }
  },
  notFound: {
    position: 'fixed',
    top: '50%',
    left: '50%',
    transform: 'translateX(-50%)translateY(-50%)',
  }
}));

const DeploymentDetails = () => {
  const location = useLocation();
  const history = useHistory();
  const { tab = '0' } = querystring.parse(location.search);
  const { deploymentId } = useParams();
  const handleTabChange = (event, newValue) => {
    history.replace({
      pathname: location.pathname,
      search: `?${new URLSearchParams({
        tab: newValue,
      })}`,
    });
  };
  const onClickBack = () => {
    if (history.length <= 2) {
      history.push('/');
    } else {
      history.goBack();
    }
  };
  const classes = useStyles();
  return (
    <DeploymentDetailsContextProvider id={`${deploymentId}`}>
      {
        ({ data, loading, error }) => {
          if (error) {
            return <div className={classes.notFound}><NoData  message="Deployment not found" /></div>;
          }
          if (loading) {
            return (
              <Box
                m={2}
                flexGrow={1}
                justifyContent="space-around"
                display="flex"
                flexDirection="column"
              >
                <Loader />
              </Box>
            );
          }
          return (
            <PageContent>
              <Box mt={3} mb={3}>
                <Typography variant="h3">
                  <IconButton aria-label="back" onClick={onClickBack}>
                    <ArrowBackIcon fontSize="large" />
                  </IconButton>
                  {data.name}
                </Typography>
                <Box mt={1} mb={1} className={classes.chips}>
                  <DeploymentStatus />
                  <Chip label={(
                    <Typography>
                      Namespace:
                      {data.namespace}
                    </Typography>
                  )}
                  />
                  <Chip label={(
                    <Typography>
                      Cluster:
                      {data.cluster}
                    </Typography>
                  )}
                  />
                  <Chip label={(
                    <Typography>
                      Deployment Time:
                      {moment.unix(data.time).format('DD/MM/YYYY HH:mm:ss')}
                    </Typography>
                  )}
                  />
                </Box>
              </Box>
              <Kinds selectedTab={parseInt(tab)} onTabChange={handleTabChange} />
              <Box mt={3} mb={3}>
                <ReplicasStats kindIndex={parseInt(tab)} />
              </Box>
              <Box mt={3} mb={3}>
                <PodEvents kindIndex={parseInt(tab)} deploymentId={deploymentId} />
                <DeploymentEvents kindIndex={parseInt(tab)} />
                <ReplicaSetEvents kindIndex={parseInt(tab)} />
                <ServiceSetEvents kindIndex={parseInt(tab)} />
                <Metrics kindIndex={parseInt(tab)} />
                <Alerts kindIndex={parseInt(tab)} />
              </Box>
            </PageContent>
          );
        }
      }
    </DeploymentDetailsContextProvider>
  );
};

export default DeploymentDetails;
