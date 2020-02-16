import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';
import React from 'react';
import makeStyles from '@material-ui/styles/makeStyles';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';

const useStyles = makeStyles((theme) => ({
  root: {
    boxShadow: 'none',
  },
}));

const Kinds = ({ selectedTab, onTabChange }) => {
  const classes = useStyles();
  const {data} = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  return (
    <AppBar position="static" color="primary" classes={{ root: classes.root }}>
      <Tabs value={selectedTab} onChange={onTabChange}>
        {
          data.kinds.map((kind, index) => <Tab label={kind.name} value={index} disableRipple />)
      }
      </Tabs>
    </AppBar>
  );
};

export default Kinds;
