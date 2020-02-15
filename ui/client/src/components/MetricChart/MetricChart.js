import React from 'react';
import PropTypes from 'prop-types';
import Widget from '../Widget/Widget';
import LineChart from '../Charts/Line/LineChart';

const MetricChart = ({ metric, deploymentTime, series }) => (
  <Widget title={metric}>
    <LineChart series={series} deploymentTime={deploymentTime} />
  </Widget>
);

MetricChart.propTypes = {
  metric: PropTypes.string.isRequired,
  deploymentTime: PropTypes.number,
  series: PropTypes.arrayOf(PropTypes.any),
};

MetricChart.defaultProps = {
  deploymentTime: null,
  series: [],
};

export default MetricChart;
