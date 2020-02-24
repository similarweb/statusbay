import React from 'react';
import Grid from '@material-ui/core/Grid';
import * as PropTypes from 'prop-types';
import Box from '@material-ui/core/Box';
import Widget from '../components/Widget/Widget';
import SpotChart from '../components/Charts/Spot/SpotChart';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';
import AlertsChartContainer from './AlertsChartContainer';
import MetricIntegrationModal from '../components/MetricIntegrationModal';

const Alerts = ({ kindIndex }) => {
  const { data } = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  const { alerts } = data.kinds[kindIndex];
  let content;
  if (alerts === []) {
    content = (
      <Grid item xs={12}>
        <Widget>
          <Box
            mt={2}
            mb={2}
            display="flex"
            justifyContent="space-around"
          >
            <MetricIntegrationModal />
          </Box>
        </Widget>
      </Grid>
    );
  } else {
    content = (
      <Grid container spacing={2}>
        {
        alerts.map(({ provider, tags }) => (
          <Grid key={`${provider}-${tags}`} item xs={12}>
            <Widget title={provider}>
              <AlertsChartContainer deploymentTime={data.time} provider={provider} tags={tags} />
            </Widget>
          </Grid>
        ))
      }
      </Grid>
    );
  }
  return (
    <DeploymentDetailsSection title="Alerts" defaultExpanded>
      {content}
    </DeploymentDetailsSection>
  );
};

Alerts.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default Alerts;
