import React from 'react';
import * as PropTypes from 'prop-types';
import ReplicasStatsBoxes from '../components/ReplicasStats/ReplicasStatsBoxes';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const ReplicasStats = ({ kindIndex }) => {
  const data = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  return <ReplicasStatsBoxes data={data.kinds[kindIndex].stats} />;
};

ReplicasStats.propTypes = {
  kindIndex: PropTypes.number.isRequired,
};

export default ReplicasStats;
