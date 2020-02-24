import React from 'react';
import PropTypes from 'prop-types';
import { useAlertsData } from '../Hooks/AlertsHooks';
import SpotChart from '../components/Charts/Spot/SpotChart';
import Skeleton from '@material-ui/lab/Skeleton';

const AlertsChartContainer = ({
  provider, tags, deploymentTime,
}) => {
  const { data, loading } = useAlertsData(provider, tags, deploymentTime);
  if (loading) {
    return <Skeleton variant="rect" width="auto" height={118} />;
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
