import React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import ErrorOutlineIcon from '@material-ui/icons/ErrorOutline';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme) => ({
  root: {
    borderLeft: `2px solid ${theme.palette.warning.dark}`,
    height: 60,
    display: 'flex',
    alignItems: 'center',
    backgroundColor: theme.palette.warning.light,
    padding: theme.spacing(2)
  },
  icon: {
    marginRight: theme.spacing(1)
  }
}));
const NoData = () => {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <ErrorOutlineIcon className={classes.icon}/> <Typography variant="subtitle1">No results match your search criteria</Typography>
    </div>
  );
};

export default NoData;
