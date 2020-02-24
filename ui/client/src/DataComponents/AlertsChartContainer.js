import React from 'react';
import PropTypes from 'prop-types';
import { useAlertsData } from '../Hooks/AlertsHooks';
import SpotChart from '../components/Charts/Spot/SpotChart';

const AlertsChartContainer = ({
  provider, tags, deploymentTime,
}) => {
  const { data } = useAlertsData(provider, tags, deploymentTime);
  if (!data) {
    return null;
  }
  const series = data.map(({ name, url, periods }) => ({
    name,
    link: url,
    data: periods.map((p) => ({
      from: p.start,
      to: p.end,
    })),
  }));
  return (
    <SpotChart series={series} deploymentTime={deploymentTime} />
  );
};

AlertsChartContainer.propTypes = {
  provider: PropTypes.string.isRequired,
  tags: PropTypes.string.isRequired,
  deploymentTime: PropTypes.number.isRequired,
};

export default AlertsChartContainer;
