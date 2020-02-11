import React from 'react';
import PropTypes from 'prop-types';
import DoneIcon from '@material-ui/icons/Done';
import { green, red, cyan, yellow } from '@material-ui/core/colors';
import { makeStyles } from '@material-ui/core/styles';
import Chip from '@material-ui/core/Chip';
import ClearIcon from '@material-ui/icons/Clear';
import AlarmOffIcon from '@material-ui/icons/AlarmOff';
import CircularProgress from '@material-ui/core/CircularProgress';
import { deploymentStatuses } from '../../constants';

const color = (status) => {
  switch (status) {
    case 'running':
      return cyan[500];
    case 'successful':
      return green[500];
    case 'timeout':
      return yellow[500];
    default:
      return red[500];
  }
};

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: ({ status }) => color(status),
    color: theme.palette.primary.contrastText,
  },
  icon: {
    color: theme.palette.primary.contrastText,
  },
}));


const messages = {
  running: 'Deployment is running',
  successful: 'Deployment completed successfully',
  failed: 'Deployment failed',
  timeout: 'Deployment timeout',
  deleted: 'Deployment deleted'
};

const icons = {
  running: <CircularProgress size={16} color="secondary" />,
  successful: <DoneIcon color="primary" />,
  failed: <ClearIcon color="primary" />,
  timeout: <AlarmOffIcon color="primary" />,
};

const DeploymentStatusBox = ({ status }) => {
  const classes = useStyles({ status });
  if (status) {
    return (
      <Chip
        icon={icons[status]}
        label={messages[status]}
        classes={{ root: classes.root, icon: classes.icon }}
      />
    );
  }
  return null;
};

DeploymentStatusBox.propTypes = {
  status: PropTypes.oneOf(deploymentStatuses),
};

DeploymentStatusBox.defaultProps = {
  status: null,
};

export default DeploymentStatusBox;
