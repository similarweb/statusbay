import React from 'react';

import DeploymentStatusBox from '../components/DeploymentStatusBox';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const DeploymentStatus = () => {
  const data = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  return <DeploymentStatusBox status={data.status}/>;
};

export default DeploymentStatus;
