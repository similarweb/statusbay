import React from 'react';
import PropTypes from 'prop-types';
import Skeleton from '@material-ui/lab/Skeleton';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import { makeStyles, useTheme } from '@material-ui/core';
import ThumbUpOutlinedIcon from '@material-ui/icons/ThumbUpOutlined';
import SpotChart from '../components/Charts/Spot/SpotChart';
import { useAlertsData } from '../Hooks/AlertsHooks';

const useStyles = makeStyles((theme) => ({
  noAlerts: {
    display: 'flex',
    alignItems: 'center',
    '& svg': {
      marginRight: 12
    }
  },
}));

const AlertsChartContainer = ({
  provider, tags, deploymentTime,
}) => {
  const theme = useTheme();
  const classes = useStyles();
  const { data, loading } = useAlertsData(provider, tags, deploymentTime);
  // loading state
  if (loading) {
    return <Skeleton variant="rect" width="auto" height={118} />;
  }
  // no alerts state
  if (data.length === 0) {
    return (
      <Box display="flex" justifyContent="space-around">
        <Typography
          className={classes.noAlerts}
          variant="h5"
        >
          <ThumbUpOutlinedIcon fontSize="medium" style={{ color: theme.palette.success.main }} />
Alerts not found
        </Typography>
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
