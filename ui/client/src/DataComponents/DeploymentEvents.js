import React from 'react';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import EventsViewLogs from '../components/EventsView/EventsViewLogs';
import * as PropTypes from 'prop-types';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const titles = {
  deployment: 'Deployment Events',
  daemonSet: 'DaemonSet Events',
};

const DeploymentEvents = ({ kindIndex }) => {
  const data = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  return (
    <DeploymentDetailsSection title={titles[data.kinds[kindIndex].type]} defaultExpanded>
      <Grid container>
        <Grid item xs={12}>
          <Paper>
            <EventsViewLogs logs={data.kinds[kindIndex].deploymentEvents} />
          </Paper>
        </Grid>
      </Grid>
    </DeploymentDetailsSection>
  );
};

DeploymentEvents.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default DeploymentEvents;
