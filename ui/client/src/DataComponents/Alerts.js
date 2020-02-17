import React from 'react';
import Grid from '@material-ui/core/Grid';
import Widget from '../components/Widget/Widget';
import SpotChart from '../components/Charts/Spot/SpotChart';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import * as PropTypes from 'prop-types';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const Alerts = ({ kindIndex }) => {
  const {data} = useDeploymentDetailsContext()
  if (!data) {
    return null;
  }
  const { deploymentTime, providers } = data.kinds[kindIndex].alerts;
  return (
    <DeploymentDetailsSection title="Alerts" defaultExpanded><Grid container spacing={2}>
      {
        providers.map((provider) => (
          <Grid key={provider.provider} item xs={12}>
            <Widget title={provider.provider}>
              <SpotChart deploymentTime={deploymentTime} series={provider.data}/>
            </Widget>
          </Grid>
        ))
      }
    </Grid></DeploymentDetailsSection>
  );
};

Alerts.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default Alerts;
