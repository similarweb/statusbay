import React, { useState } from 'react';
import { Button, Dialog, makeStyles } from '@material-ui/core';
import WarningIcon from '@material-ui/icons/Warning';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import image from './metrics.png';

const useStyles = makeStyles((theme) => ({
  code: {
    whiteSpace: 'pre',
    fontFamily: 'monospace',
  },
}));

const MetricIntegrationModal = () => {
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
      <Button variant="contained" color="primary" onClick={handleClick} startIcon={<WarningIcon />}>Click To Integration</Button>
      <Dialog open={isOpen} onClose={handleDialogClose} closeAfterTransition={true} onBackdropClick={handleDialogClose}>
        <Box p={2}>
          <Typography variant="h4">Kubernetes:</Typography>
          <Typography variant="body1">
            Add
            {' '}
            <Link href="https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/" target="_blank">Annotation</Link>
            {' '}
to deployment:
          </Typography>
          <Typography className={classes.code}>apiVersion: apps/v1</Typography>
          <Typography className={classes.code}>kind: Deployment</Typography>
          <Typography className={classes.code}>metadata:</Typography>
          <Typography className={classes.code}>  annotations:</Typography>
          <Typography className={classes.code}>    statusbay.io/alerts-statuscake-tags: nginx</Typography>
          <img width={560} src={image} />
        </Box>
      </Dialog>
    </div>
  );
};

export default MetricIntegrationModal;
