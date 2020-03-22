import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import PropTypes from 'prop-types';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import CloseIcon from '@material-ui/icons/Close';
import { usePodLogs } from '../../Hooks/PodLogsHooks';
import Loader from '../Loader/Loader';

const useStyles = makeStyles((theme) => ({
  wrapper: {
    width: '100%', backgroundColor: '#424242', whiteSpace: 'nowrap', padding: '24px 0',
  },
  line: {
    fontFamily: '"Monaco", monospace',
    color: '#ffffff',
    fontSize: '12px',
    lineHeight: '1.8',
  },
  lineNumber: {
    display: 'inline-block',
    width: '50px',
    margin: '0 40px 0 0',
    direction: 'rtl',
  },
}));

const ContainersLogsPopup = ({ deploymentId, podName, onClose }) => {
  const classes = useStyles();
  const { data, loading, error } = usePodLogs(deploymentId, podName);
  const [selectedTab, setSelectedTab] = useState(0);

  const handleTabChange = (event, newValue) => {
    setSelectedTab(newValue);
  };
  const logs = data ? data[selectedTab].logs : [];
  return (
    <>
      <Toolbar>
        <IconButton edge="start" color="inherit" onClick={onClose} aria-label="close">
          <CloseIcon />
        </IconButton>
        <Typography variant="h6" className={classes.title}>
          Logs
        </Typography>
      </Toolbar>
      {
        loading && <Loader interval={100} />
      }
      {
        !loading && data ? (
          <>
            <Tabs value={selectedTab} onChange={handleTabChange}>
              {
                data.map(({ name }, index) => (
                  <Tab
                    key={name}
                    label={name}
                    value={index}
                    disableRipple
                  />
                ))
              }
            </Tabs>
            <div className={classes.wrapper}>
              {
                logs.map((log, index) => (
                  <div className={classes.line}>
                    <span className={classes.lineNumber}>
                      {index + 1}
                    </span>
                    <span>{log}</span>
                  </div>
                ))
              }

            </div>
          </>
        ) : null
      }
    </>
  );
};

ContainersLogsPopup.propTypes = {
  deploymentId: PropTypes.string,
  podName: PropTypes.string,
  onClose: PropTypes.func,

};

ContainersLogsPopup.defaultProps = {
  deploymentId: '',
  podName: '',
  onClose: () => null,
};

export default ContainersLogsPopup;
