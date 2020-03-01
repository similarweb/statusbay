import React from 'react';
import LineChart from './Line/LineChart';
import mock from './Line/mock';
import SpotChart from './Spot/SpotChart';

export default {
  title: 'UI|Charts',
};

const mock2 = [
  {
    name: 'Pro: Leading Folders (mobile) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=3588189',
    data: [
      {
        from: 1581006979,
        to: 1581008480,
      },
    ],
  },
  {
    name: 'Pro: Leading Folders (desktop) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=3588188',
    data: [
      {
        from: 1581006948,
        to: 1581008738,
      },
    ],
  },
  {
    name: 'Pro: Leading Folders Table Window (HC) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=5319887',
    data: [
      {
        from: 1581007132,
        to: 1582186417,
      },
      {
        from: 1581006929,
        to: 1581007056,
      },
    ],
  },
  {
    name: 'Pro: Leading Folders Table 3Month (HC) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=3588146',
    data: [
      {
        from: 1581006978,
        to: 1582186417,
      },
    ],
  },
  {
    name: 'Pro: Popular Pages (desktop) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=3588157',
    data: [
      {
        from: 1581006982,
        to: 1581008397,
      },
    ],
  },
  {
    name: 'Pro: Popular Pages (mobile) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=3588183',
    data: [
      {
        from: 1581007006,
        to: 1581008477,
      },
    ],
  },
];

const mock3 = [
  {
    name: 'Pro: Leading Folders Table Window (HC) (us-east-1)',
    link: 'https://app.statuscake.com/UptimeStatus.php?tid=5319887',
    data: [
      {
        from: 1582185817,
        to: 1582186117,
      },
      {
        from: 1582186717,
        to: 1582187017,
      },
    ],
  }
];

export const LineChartStory = () => (
  <LineChart
    series={mock}
    deploymentTime={1581929522 * 1000}
  />
);
LineChartStory.story = 'Line Chart';

export const SpotChartStory = () => (
  <SpotChart
    deploymentTime={1582186417}
    series={mock3}
  />
);
SpotChartStory.story = 'Spot Chart';
