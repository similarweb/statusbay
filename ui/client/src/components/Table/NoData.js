import React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import image from './no-data.svg';

const useStyles = makeStyles((theme) => ({
  text: {
    color: '#bdbdbd',
    marginTop: 8,
    fontWeight: 400,
  },
}));
const NoData = () => {
  const classes = useStyles();
  return (
    <Box display="flex" flexDirection="column" alignItems="center" mt={4} mb={4}>
      <img src={image} alt="" width={210} />
      <Typography variant="h6" className={classes.text}>No results match your search</Typography>
    </Box>
  );
};

export default NoData;
