import React from 'react';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import * as PropTypes from 'prop-types';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import EventsViewLogs from '../components/EventsView/EventsViewLogs';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const titles = {
  deployment: 'Deployment Events',
  daemonSet: 'DaemonSet Events',
  StatefulSet: 'StatefulSet Events',
};

const DeploymentEvents = ({ kindIndex }) => {
  const { data } = useDeploymentDetailsContext();
  if (!data || (Array.isArray(data.kinds[kindIndex].deploymentEvents) && data.kinds[kindIndex].deploymentEvents.length === 0)) {
    return null;
  }
  return (
    <DeploymentDetailsSection title={titles[data.kinds[kindIndex].type]} defaultExpanded>
      <Grid container>
        <Grid item xs={12}>
          <Paper>
            <EventsViewLogs logs={data.kinds[kindIndex].deploymentEvents}/>
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
