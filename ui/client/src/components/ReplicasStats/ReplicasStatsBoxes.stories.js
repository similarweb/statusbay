import React from 'react';
import { number } from '@storybook/addon-knobs';
import ReplicasStatsBoxes from './ReplicasStatsBoxes';

export default {
  title: 'UI|Deployment Details/status boxes',
};


export const story = () => {
  const data = {
    desired: number('Desired', 1, {
      range: true, min: 0, max: 10, step: 1,
    }),
    current: number('Current', 1, {
      range: true, min: 0, max: 10, step: 1,
    }),
    updated: number('Updated', 1, {
      range: true, min: 0, max: 10, step: 1,
    }),
    ready: number('Ready', 1, {
      range: true, min: 0, max: 10, step: 1,
    }),
    available: number('Available', 1, {
      range: true, min: 0, max: 10, step: 1,
    }),
    unavailable: number('Unavailable', 0, {
      range: true, min: 0, max: 10, step: 1,
    }),
  };
  return <ReplicasStatsBoxes data={data} />;
};
