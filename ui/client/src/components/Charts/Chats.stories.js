import React from 'react';
import LineChart from './Line/LineChart';
import mockLine from './Line/mock';
import mockSpot from './Spot/mock';
import SpotChart from './Spot/SpotChart';

export default {
  title: 'UI|Charts',
};


export const LineChartStory = () => (
  <LineChart
    series={mockLine}
    deploymentTime={1583223569}
  />
);
LineChartStory.story = 'Line Chart';

export const SpotChartStory = () => (
  <SpotChart
    deploymentTime={1582186417}
    series={mockSpot}
  />
);
SpotChartStory.story = 'Spot Chart';
