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
  const [newVersion, setNewVersion] = useState(null);
  useEffect(() => {
    const getVersionData = async () => {
      const { data, error } = await API('/api/version');
      if (data && data.outdated) {
        setNewVersion(data);
      }
    };
    getVersionData();
  }, []);
  return newVersion && (
    <Alert severity="success" classes={{ root: classes.alertRoot, message: classes.alertMessage }}>
      New version of StatusBay is available: {newVersion.current_version}.
      <a style={{ marginLeft: 8 }} href={newVersion.current_download_url} target="_blank">Get it
        now!</a>
    </Alert>
  );
};

export default StatusBayVersion;
