import Alert from '@material-ui/lab/Alert';
import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core';
import API from '../../../Services/API/index';

const useStyles = makeStyles((theme) => ({
  alertRoot: {
    paddingTop: 4,
    paddingBottom: 4,
  },
  alertMessage: {
    flexDirection: 'row',
  },
}));

const StatusBayVersion = () => {
  const classes = useStyles();
  const [showNewVersion, setShowNewVersion] = useState(false);
  useEffect(() => {
    const getVersionData = async () => {
      const { data, error } = await API('/api/version');
      if (data && data.Outdated) {
        setShowNewVersion(true);
      }
    };
    getVersionData();
  }, []);
  return showNewVersion && (
  <Alert severity="success" classes={{ root: classes.alertRoot, message: classes.alertMessage }}>
    New version of StatusBay is available.
    <a style={{ marginLeft: 8 }} href="https://github.com/similarweb/statusbay" target="_blank">Get it now!</a>
  </Alert>
  );
};

export default StatusBayVersion;
