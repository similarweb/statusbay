import React from 'react';
import PropTypes from 'prop-types';
import DoneIcon from '@material-ui/icons/Done';
import { cyan, amber } from '@material-ui/core/colors';
import { makeStyles } from '@material-ui/core/styles';
import Chip from '@material-ui/core/Chip';
import ClearIcon from '@material-ui/icons/Clear';
import AlarmOffIcon from '@material-ui/icons/AlarmOff';
import CircularProgress from '@material-ui/core/CircularProgress';
import { deploymentStatuses } from '../../constants';

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: ({ status }) => {
      switch (status) {
        case 'running':
          return cyan[500];
        case 'successful':
          return theme.palette.success.main;
        case 'timeout':
          return amber[500];
        case 'deleted':
          return theme.palette.grey[500];
        case 'failed':
          return theme.palette.error.main;
        case 'cancelled':
          return theme.palette.grey[500];
      }
    },
    color: theme.palette.primary.contrastText,
    textTransform: 'uppercase'
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
  deleted: 'Deployment deleted',
  cancelled: 'Deployment cancelled',
};

const DeploymentStatusBox = ({ status }) => {
  const classes = useStyles({ status });
  if (status) {
    return (
      <Chip
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
