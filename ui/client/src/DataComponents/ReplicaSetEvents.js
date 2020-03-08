import React from 'react';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import * as PropTypes from 'prop-types';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import EventsViewLogs from '../components/EventsView/EventsViewLogs';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const ReplicaSetEvents = ({ kindIndex }) => {
  const { data } = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }

  const replicaSets = data.kinds[kindIndex].replicaSet || [];
  return (
    <>
      {
        replicaSets.map((set) => set.logs.length > 0 && (
        <DeploymentDetailsSection key={set.name} title={set.name}>
          <Grid container>
            <Grid item xs={12}>
              <Paper>
                <EventsViewLogs logs={set.logs} />
              </Paper>
            </Grid>
          </Grid>
        </DeploymentDetailsSection>
        ))
      }

    </>
  );
};

ReplicaSetEvents.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default ReplicaSetEvents;
