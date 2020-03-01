import React from 'react';
import PropTypes from 'prop-types';
import Skeleton from '@material-ui/lab/Skeleton';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import { makeStyles, useTheme } from '@material-ui/core';
import ThumbUpOutlinedIcon from '@material-ui/icons/ThumbUpOutlined';
import Alert from '@material-ui/lab/Alert';
import AlertTitle from '@material-ui/lab/AlertTitle';
import SpotChart from '../components/Charts/Spot/SpotChart';
import { useAlertsData } from '../Hooks/AlertsHooks';

const useStyles = makeStyles((theme) => ({
  noAlerts: {
    fontSize: 16,
  },
}));

const AlertsChartContainer = ({
  provider, tags, deploymentTime,
}) => {
  const classes = useStyles();
  const { data, loading, error, tagsWarning } = useAlertsData(provider, tags, deploymentTime);
  // loading state
  if (loading) {
    return <Skeleton variant="rect" width="auto" height={118} />;
  }
  if (tagsWarning) {
    return (
      <Alert severity="warning">
        <AlertTitle>Check tags not found:</AlertTitle>
        <code>{JSON.stringify(tagsWarning, undefined, 4)}</code>
      </Alert>
    );
  }
  if (error) {
    return (
      <Alert severity="error">
        <AlertTitle>Alerts error:</AlertTitle>
        <code>{JSON.stringify(error, undefined, 4)}</code>
      </Alert>
    );
  }
  // no alerts state
  if (data.length === 0) {
    return (
      <Box display="flex" justifyContent="space-around">
        <Alert severity="success" classes={{ message: classes.noAlerts }} icon={<ThumbUpOutlinedIcon fontSize="medium" />}>Alerts not found</Alert>
      </Box>
    );
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
