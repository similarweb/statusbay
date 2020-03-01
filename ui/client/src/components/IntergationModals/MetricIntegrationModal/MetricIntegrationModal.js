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
  image:{
    width: "100%",
  }
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
      <Dialog fullWidth={true} maxWidth="lg" open={isOpen} onClose={handleDialogClose} closeAfterTransition={true} onBackdropClick={handleDialogClose}>
        <Box p={2}> 
          <Typography variant="h4">Metric Integration:</Typography>
          <Typography variant="body1">
          StatusBay has the ability to show different metrics for different set of queries. More info at 
            {' '}
            <Link href="https://github.com/similarweb/statusbay/blob/master/docs/integrations.md" target="_blank">StatusBay README</Link>
            {' '}
           <br/><br/>
          Annotation example:
          </Typography>
          <Typography className={classes.code}>apiVersion: apps/v1</Typography>
          <Typography className={classes.code}>kind: Deployment</Typography>
          <Typography className={classes.code}>metadata:</Typography>
          <Typography className={classes.code}>  annotations:</Typography>
          <Typography className={classes.code}>    statusbay.io/metrics-datadog-2xx: "sum:web.http.2xxֿ{`{*}`}.as_count()"</Typography>
          <br/><br/>
          <img className={classes.image} src={image} />
        </Box>
      </Dialog>
    </div>
  );
};

export default MetricIntegrationModal;
