import React from 'react';
import EventsView from '../components/EventsView/EventsView';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const PodEvents = ({ kindIndex }) => {
  const {data} = useDeploymentDetailsContext();
  if (!data || (Array.isArray(data.kinds[kindIndex].podEvents) && data.kinds[kindIndex].podEvents.length === 0)) {
    return null;
  }
  return (
    <DeploymentDetailsSection title="Pod Events" defaultExpanded>
      <EventsView
        key={kindIndex}
        items={data.kinds[kindIndex].podEvents}
      />
    </DeploymentDetailsSection>
  );
};

export default PodEvents;
