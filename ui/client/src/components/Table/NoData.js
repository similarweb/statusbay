import React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import PropTypes from 'prop-types';
import image from './no-data.svg';

const useStyles = makeStyles((theme) => ({
  text: {
    color: '#bdbdbd',
    marginTop: 8,
    fontWeight: 400,
  },
}));
const NoData = ({ message, imageWidth}) => {
  const classes = useStyles();
  return (
    <Box display="flex" flexDirection="column" alignItems="center" mt={4} mb={4}>
      <img src={image} alt="" width={imageWidth} />
      <Typography variant="h6" className={classes.text}>{message}</Typography>
    </Box>
  );
};
NoData.propTypes = {
  message: PropTypes.string,
  imageWidth: PropTypes.number
};
NoData.defaultProps = {
  message: 'No results match your search',
  imageWidth: 210
};

export default NoData;
