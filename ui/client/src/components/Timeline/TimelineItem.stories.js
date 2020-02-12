import React from 'react';
import TimelineItem from './TimelineItem';

export default {
  title: 'UI|Time line/Timeline item',
};

const data = {
  desired: 1,
  current: 2,
  updated: 3,
  ready: 4,
  available: 5,
  unavailable: 6,
};
export const story = () => <TimelineItem data={data} />;
