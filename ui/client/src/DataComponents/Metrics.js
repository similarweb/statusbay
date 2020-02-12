import React from 'react';
import Grid from '@material-ui/core/Grid';
import MetricChart from '../components/MetricChart/MetricChart';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import * as PropTypes from 'prop-types';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const Metrics = ({ kindIndex }) => {
  const data = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  const { deploymentTime, ...allMetrics } = data.kinds[kindIndex].metrics;
  return (
    <DeploymentDetailsSection title="Metrics" defaultExpanded><Grid container spacing={2}>
      {
        Object.entries(allMetrics).map(([metricName, series]) => (
          <Grid key={metricName} item xs={6}>
            <MetricChart metric={metricName} series={series} deploymentTime={deploymentTime}/>
          </Grid>
        ))
      }
    </Grid></DeploymentDetailsSection>
  );
};

Metrics.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default Metrics;
