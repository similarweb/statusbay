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
import makeStyles from '@material-ui/core/styles/makeStyles';
import IconButton from '@material-ui/core/IconButton';
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

const useStyles = makeStyles((theme) => ({
  chips: {
    '& > *': {
      margin: theme.spacing(0.5),
    },
  },
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
    history.goBack();
  };
  const classes = useStyles();
  return (
    <DeploymentDetailsContextProvider id={`${deploymentId}`}>
      {
        ({ data, loading }) => (loading ? <Box m={2} flexGrow={1} justifyContent="space-around" display="flex" flexDirection="column"><Loader /></Box> : (
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
                    {moment.unix(data.time).utc().format('DD/MM/YYYY HH:MM:ss')}
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
