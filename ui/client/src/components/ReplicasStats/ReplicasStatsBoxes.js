import React from 'react';
import PropTypes from 'prop-types';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import Box from '@material-ui/core/Box';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme) => ({
  typography: {
    textTransform: 'uppercase',
  },
  paper: {
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  cardError: {
    backgroundColor: theme.palette.error.main,
  },
}));

const SingleBox = ({ name, data, error }) => {
  const classes = useStyles({ error });
  return (
    <Card classes={{ root: error ? classes.cardError : undefined }}>
      <Box display="flex" flexDirection="column" alignItems="center">
        <Box m={1}><Typography classes={{ root: classes.typography }} variant="h6">{name}</Typography></Box>
        <Box m={1}><Typography classes={{ root: classes.typography }} variant="h4">{data}</Typography></Box>
      </Box>
    </Card>
  );
};
SingleBox.propTypes = {
  name: PropTypes.string.isRequired,
  data: PropTypes.number,
  error: PropTypes.bool,
};
SingleBox.defaultProps = {
  data: null,
  error: false,
};

const ReplicasStatsBoxes = ({ data }) => {
  const classes = useStyles;
  const {
    desired, current, updated, ready, available, unavailable,
  } = data;
  if (!data) {
    return null;
  }
  return (
    <Box classes={{ root: classes.root }}>
      <Grid container spacing={3}>
        <Grid item xs={2}><SingleBox name="desired" data={desired} className={classes.paper} /></Grid>
        <Grid item xs={2}><SingleBox name="current" data={current} className={classes.paper} /></Grid>
        <Grid item xs={2}><SingleBox name="updated" data={updated} className={classes.paper} /></Grid>
        <Grid item xs={2}><SingleBox name="ready" data={ready} className={classes.paper} /></Grid>
        <Grid item xs={2}><SingleBox name="available" data={available} className={classes.paper} /></Grid>
        <Grid item xs={2}><SingleBox name="unavailable" data={unavailable} className={classes.paper} error={unavailable > 0} /></Grid>
      </Grid>
    </Box>

  );
};

ReplicasStatsBoxes.propTypes = {
  data: PropTypes.shape({
    desired: PropTypes.number,
    current: PropTypes.number,
    updated: PropTypes.number,
    ready: PropTypes.number,
    available: PropTypes.number,
    unavailable: PropTypes.number,
  }),
};

ReplicasStatsBoxes.defaultProps = {
  data: {
    desired: 0,
    current: 0,
    updated: 0,
    ready: 0,
    available: 0,
    unavailable: 0,
  },
};

export default ReplicasStatsBoxes;
