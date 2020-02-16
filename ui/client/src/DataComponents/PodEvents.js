import React from 'react';
import EventsView from '../components/EventsView/EventsView';
import DeploymentDetailsSection from '../components/DeploymentDetailsSection';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const PodEvents = ({ kindIndex }) => {
  const {data} = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  return (
    <DeploymentDetailsSection title="Pod Events" defaultExpanded>
      <EventsView
        items={data.kinds[kindIndex].podEvents}
      />
    </DeploymentDetailsSection>
  );
};

export default PodEvents;
