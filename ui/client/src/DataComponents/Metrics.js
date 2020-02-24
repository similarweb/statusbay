import React from 'react';
import Grid from '@material-ui/core/Grid';
import * as PropTypes from 'prop-types';
import Box from '@material-ui/core/Box';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';
import MetricChartContainer from './MetricChartContainer';
import Widget from '../components/Widget/Widget';
import MetricIntegrationModal
  from '../components/IntergationModals/MetricIntegrationModal/MetricIntegrationModal';

const Metrics = ({ kindIndex }) => {
  const { data } = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  const { metrics } = data.kinds[kindIndex];
  let content;
  if (metrics.length === 0) {
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
    content = metrics.map((metric) => (
      <MetricChartContainer
        key={metric.name}
        name={metric.name}
        provider={metric.provider}
        query={metric.query}
        deploymentTime={data.time}
      />
    ));
  }
  return (
    <div>
      <DeploymentDetailsSection title="Metrics" defaultExpanded>
        <Grid container spacing={2}>
          {
          content
        }
        </Grid>
      </DeploymentDetailsSection>
    </div>
  );
};

Metrics.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default Metrics;
