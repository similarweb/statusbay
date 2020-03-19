import React, { useState } from 'react';
import { Dialog, makeStyles } from '@material-ui/core';
import Link from '@material-ui/core/Link';
import PropTypes from 'prop-types';
import ContainersLogsPopup from './ContainersLogsPopup';

const useStyles = makeStyles((theme) => ({
  dialog: {
    top: '64px !important;',
  },
}));

const ContainersLogs = ({ deploymentId, podName }) => {
  const classes = useStyles();
  const [isOpen, setIsOpen] = useState(false);
  const handleClick = () => {
    setIsOpen(true);
  };
  const handleDialogClose = () => {
    setIsOpen(false);
  };
  return (
    <div>
      <Link onClick={handleClick}>Show</Link>
      {
        isOpen && (
        <Dialog className={classes.dialog} open onClose={handleDialogClose} closeAfterTransition={true} onBackdropClick={handleDialogClose} fullScreen>
          <ContainersLogsPopup onClose={handleDialogClose} deploymentId={deploymentId} podName={podName} />
        </Dialog>
        )
      }
    </div>
  );
};

ContainersLogs.propTypes = {
  deploymentId: PropTypes.string,
  podName: PropTypes.string,

};

ContainersLogs.defaultProps = {
  deploymentId: '',
  podName: '',
};

export default ContainersLogs;
