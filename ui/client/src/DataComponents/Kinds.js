import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';
import React from 'react';
import { useDeploymentDetailsContext } from '../context/DeploymentDetailsContext';


const Kinds = ({ selectedTab, onTabChange }) => {
  const data = useDeploymentDetailsContext();
  if (!data) {
    return null;
  }
  return (
    <AppBar position="static" color="primary">
      <Tabs value={selectedTab} onChange={onTabChange}>
        {
          data.kinds.map((kind, index) => <Tab label={kind.name} value={index} disableRipple />)
      }
      </Tabs>
    </AppBar>
  );
};

export default Kinds;
