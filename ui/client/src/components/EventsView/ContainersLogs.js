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

const useStyles = makeStyles((theme) => ({
  dialog: {
    top: "64px !important;",
  },
}));

const ContainersLogs = ({ podName }) => {
  const classes = useStyles();
  const [isOpen, setIsOpen] = useState(false);
  const handleClick = () => {
    setIsOpen(true);
  };

  const handleDialogClose = () => {
    setIsOpen(false);
  };

  const [value, setValue] = useState(0);
  const handleTabChange = (event, newValue) => {
    setValue(newValue);
  };
  const text = `
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
TEST message
`;
  return (

    <div>
      <Link onClick={handleClick}>Show</Link>
      <Dialog className={classes.dialog} open={isOpen} onClose={handleDialogClose} closeAfterTransition={true} onBackdropClick={handleDialogClose} fullScreen>
        <Toolbar>
            <IconButton edge="start" color="inherit" onClick={handleDialogClose} aria-label="close">
              <CloseIcon />
            </IconButton>
            <Typography variant="h6" className={classes.title}>
              Logs
            </Typography>
          </Toolbar>
        <Tabs value={value} onChange={handleTabChange} >
          <Tab  label={"List of pods"} value={"index"} disableRipple />
        </Tabs>


        <div style={{ height: "100%", width: "100%" }}>
          <LazyLog extraLines={1} enableSearch text={text} caseInsensitive />
        </div>

      </Dialog>
      </div>
  );
};

ContainersLogs.propTypes = {
  podName: PropTypes.string,

};

ContainersLogs.defaultProps = {
};

export default ContainersLogs; 