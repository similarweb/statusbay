import React from 'react';
import EventsView from '../components/EventsView/EventsView';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';
import PropTypes from 'prop-types';

const PodEvents = ({ kindIndex, deploymentId }) => {
  const {data} = useDeploymentDetailsContext();
  if (!data || (Array.isArray(data.kinds[kindIndex].podEvents) && data.kinds[kindIndex].podEvents.length === 0)) {
    return null;
  }
  return (
    <DeploymentDetailsSection title="Pod Events" defaultExpanded>
      <EventsView
        key={kindIndex}
        items={data.kinds[kindIndex].podEvents}
        deploymentId={deploymentId}
      />
    </DeploymentDetailsSection>
  );
};

export default PodEvents;
