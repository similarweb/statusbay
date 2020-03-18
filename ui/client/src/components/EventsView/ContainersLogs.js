import React, { useState } from 'react';
import { Dialog, makeStyles } from '@material-ui/core';
import Link from '@material-ui/core/Link';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import PropTypes from 'prop-types';
import Toolbar from '@material-ui/core/Toolbar';
import { LazyLog } from 'react-lazylog';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import CloseIcon from '@material-ui/icons/Close';
import { useAlertsData } from '../../Hooks/AlertsHooks';
import { usePodLogs } from '../../Hooks/PodLogsHooks';

const useStyles = makeStyles((theme) => ({
  dialog: {
    top: '64px !important;',
  },
}));

const ContainersLogs = ({ deploymentId, podName }) => {
  const classes = useStyles();
  const { data, loading, error } = usePodLogs(deploymentId, podName);
  const [isOpen, setIsOpen] = useState(false);
  const [selectedTab, setSelectedTab] = useState(0);
  const handleClick = () => {
    setIsOpen(true);
  };
  const handleDialogClose = () => {
    setIsOpen(false);
  };
  const handleTabChange = (event, newValue) => {
    setSelectedTab(newValue);
  };
  return (
    <div>
      <Link onClick={handleClick}>Show</Link>
      <Dialog className={classes.dialog} open={data && isOpen} onClose={handleDialogClose} closeAfterTransition={true} onBackdropClick={handleDialogClose} fullScreen>
        <Toolbar>
          <IconButton edge="start" color="inherit" onClick={handleDialogClose} aria-label="close">
            <CloseIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
              Logs
          </Typography>
        </Toolbar>
        {
          data ? (
            <>
              <Tabs value={selectedTab} onChange={handleTabChange}>
                {
                data.map(({ name }, index) => <Tab key={name} label={name} value={index} disableRipple />)
              }
              </Tabs>
              <div style={{ height: '100%', width: '100%' }}>
                <LazyLog key={selectedTab} extraLines={1} enableSearch text={data[selectedTab].logs.join('\n')} caseInsensitive />
              </div>
            </>
          ) : null
        }
      </Dialog>
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
