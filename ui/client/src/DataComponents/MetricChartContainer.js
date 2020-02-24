import Grid from '@material-ui/core/Grid';
import React from 'react';
import Skeleton from '@material-ui/lab/Skeleton';
import { useMetricsData } from '../Hooks/MetricsHooks';
import LineChart from '../components/Charts/Line/LineChart';
import Widget from '../components/Widget/Widget';
import PropTypes from 'prop-types';

const MetricChartContainer = ({
  name, provider, query, deploymentTime,
}) => {
  const { data, loading } = useMetricsData(provider, query, deploymentTime);
  if (loading) {
    return (
      <Grid key={name} item xs={12} xl={6}>
        <Widget title={name}>
          <Skeleton variant="rect" width="auto" height={118} />
        </Widget>
      </Grid>
    );
  }
  const series = data.map((item) => ({
    name: item.metric,
    points: item.points,
  }));
  return (
    <Grid key={name} item xs={12} xl={6}>
      <Widget title={name}>
        <LineChart series={series} deploymentTime={deploymentTime * 1000} />
      </Widget>
    </Grid>
  );
};

export default React.memo(MetricChartContainer, (prevProps, nextProps) => {
  // optimization: ignore re-render when query and provider didn't changed
  if (prevProps.provider === nextProps.provider && prevProps.query === nextProps.query) {
    return true;
  }
  return prevProps === nextProps;
});
