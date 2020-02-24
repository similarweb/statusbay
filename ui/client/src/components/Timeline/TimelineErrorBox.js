import React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';
import PropTypes from 'prop-types';

const useStyles = makeStyles((theme) => ({
  card: {
    backgroundColor: theme.palette.error.main,
    color: theme.palette.error.contrastText,
    borderRadius: '5px',
    boxShadow: 'none'
  },
}));

const TimelineErrorBox = ({ children }) => {
  const classes = useStyles();
  return (
    <Card className={classes.card}>
      <Box m={1}>
        <Typography variant="body2" dangerouslySetInnerHTML={{ __html: children }} />
      </Box>
    </Card>
  );
};
TimelineErrorBox.propTypes = {
  children: PropTypes.node.isRequired,
};

export default TimelineErrorBox;
