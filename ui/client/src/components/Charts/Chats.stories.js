import React from 'react';
import LineChart from './Line/LineChart';
import mock from './Line/mock';
import SpotChart from './Spot/SpotChart';

export default {
  title: 'UI|Charts',
};

const mock2 = [
  {
    name: 'statuscake1',
    link: '',
    data: [
      {
        from: 1578127665,
        to: 1578127835,
      },
      {
        from: 1578127935,
        to: 1578128000,
      },
    ],
  },
  {
    name: 'statuscake2',
    link: '',
    data: [
      {
        from: 1578127665,
        to: 1578127835,
      },
    ],
  },
  {
    name: 'statuscake3',
    link: '',
    data: [
      {
        from: 1578127665,
        to: 1578127835,
      },
    ],
  },
];

export const LineChartStory = () => (
  <LineChart
    series={mock}
    deploymentTime={1577738820000}
  />
);
LineChartStory.story = 'Line Chart';

export const SpotChartStory = () => (
  <SpotChart
    deploymentTime={1578127835}
    series={mock2}
    currentTime={1578127935}
  />
);
SpotChartStory.story = 'Spot Chart';
