import React from 'react';
import PropTypes from 'prop-types';
import DoneIcon from '@material-ui/icons/Done';
import ClearIcon from '@material-ui/icons/Clear';
import DeleteSweepIcon from '@material-ui/icons/DeleteSweep';
import makeStyles from '@material-ui/core/styles/makeStyles';
import AlarmOffIcon from '@material-ui/icons/AlarmOff';
import Tooltip from '@material-ui/core/Tooltip';
import Box from '@material-ui/core/Box';
import { cyan, amber } from '@material-ui/core/colors';
import CircularProgress from '@material-ui/core/CircularProgress';
import { HighlightOff } from '@material-ui/icons';
import { deploymentStatuses } from '../../../constants';

const useStyles = makeStyles((theme) => ({
  running: {
    color: cyan[500],
  },
  failed: {
    color: theme.palette.error.main,
  },
  successful: {
    color: theme.palette.success.main,
  },
  timeout: {
    color: amber[500],
  },
  deleted: {
    color: theme.palette.grey[500],
  },
  iconWrapper: {
    display: 'flex',
    alignItems: 'center',
  },
}));

const CellStatus = ({ status }) => {
  let icon;
  const classes = useStyles();
  switch (status) {
    case 'running':
      icon = <CircularProgress classes={{ root: classes.running }} size={16} color="secondary" />;
      break;
    case 'successful':
      icon = <DoneIcon className={classes.successful} />;
      break;
    case 'failed':
      icon = <ClearIcon className={classes.failed} />;
      break;
    case 'timeout':
      icon = <AlarmOffIcon className={classes.timeout} />;
      break;
    case 'deleted':
      icon = <DeleteSweepIcon className={classes.deleted} />;
      break;
    case 'cancelled':
      icon = <HighlightOff className={classes.deleted} />;
      break;
    default:
      break;
  }
  if (!icon) {
    return null;
  }
  return (
    <Box display="flex" alignItems="center">
      <Tooltip title={status}>
        <span className={classes.iconWrapper}>
          {icon}
        </span>
      </Tooltip>
    </Box>
  );
};

CellStatus.propTypes = {
  status: PropTypes.oneOf(deploymentStatuses).isRequired,
};

export default CellStatus;
