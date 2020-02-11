import React from 'react';
import ReplicasStatsBoxes from './ReplicasStatsBoxes';

export default {
  title: 'UI|Deployment Details/status boxes',
};

const data = {
  desired: 1,
  current: 2,
  updated: 3,
  ready: 4,
  available: 5,
  unavailable: 6,
};
export const story = () => <ReplicasStatsBoxes data={data} />;
