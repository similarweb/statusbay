import Grid from '@material-ui/core/Grid';
import React from 'react';
import MetricChart from '../components/MetricChart/MetricChart';
import { useMetricsData } from '../Hooks/MetricsHooks';

const MetricChartContainer = ({
  name, provider, query, deploymentTime,
}) => {
  const { data } = useMetricsData(provider, query, deploymentTime);
  if (!data) {
    return null;
  }
  const series = data.map((item) => ({
    name: item.metric,
    points: item.points,
  }));
  return (
    <Grid key={name} item xs={12} xl={6}>
      <MetricChart metric={name} series={series} deploymentTime={deploymentTime * 1000} />
    </Grid>
  );
};

MetricChartContainer.propTypes = {

};

export default React.memo(MetricChartContainer, (prevProps, nextProps) => {
  // optimization: ignore re-render when query and provider didn't changed
  if (prevProps.provider === nextProps.provider && prevProps.query === nextProps.query) {
    return true;
  }
  return prevProps === nextProps;
});
