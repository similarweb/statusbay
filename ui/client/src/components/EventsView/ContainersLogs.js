import React, { useState } from 'react';
import { Dialog, makeStyles } from '@material-ui/core';
import Box from '@material-ui/core/Box';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import PropTypes from 'prop-types';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import CloseIcon from '@material-ui/icons/Close';

const useStyles = makeStyles((theme) => ({
  dialog: {
    top: "64px !important;",
  },
  logTerminal: {
    background: "black",
    width: "100%",
    color: "white",
    "font-size": "18px"
  }
}));

const ContainersLogs = ({ containers }) => {
  const classes = useStyles();
  const [isOpen, setIsOpen] = useState(false);
  const handleClick = () => {
    setIsOpen(true);
  };
  console.log(containers)
  const handleDialogClose = () => {
    setIsOpen(false);
  };

  const [value, setValue] = useState(0);
  const handleTabChange = (event, newValue) => {
    setValue(newValue);
  };
  return (
    
    <div>
      <span onClick={handleClick} >Click me</span>
      <Dialog className={classes.dialog} open={isOpen} onClose={handleDialogClose} closeAfterTransition={true} onBackdropClick={handleDialogClose} fullScreen>
        <Box p={2} >
        <Toolbar>
            <IconButton edge="start" color="inherit" onClick={handleDialogClose} aria-label="close">
              <CloseIcon />
            </IconButton>
            <Typography variant="h6" className={classes.title}>
              Logs
            </Typography>
          </Toolbar>
        <Tabs value={value} onChange={handleTabChange} >
        {
          containers.map((container, index) => <Tab  label={container.name} value={index} disableRipple />)
        }
        </Tabs>

        <div className={classes.logTerminal}>
        {containers.map((container, index) => 
           value == index &&
            <div key={`tab-content-${container}`} >
                {container.log.map((log, index) => <div key={`log-${container.name}-${index}`}>{}{log}</div>)}
            </div>
          ) 
        }
        </div>
        </Box>
      </Dialog>
      </div>
  );
};

ContainersLogs.propTypes = {
  containers: PropTypes.object,

};

ContainersLogs.defaultProps = {
  containers: {},
};

export default ContainersLogs;
