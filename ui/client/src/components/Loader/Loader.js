import makeStyles from '@material-ui/core/styles/makeStyles';
import React from 'react';
import PropTypes from 'prop-types';
import LoaderSvg from './LoaderSvg';

const useStyles = makeStyles((theme) => ({
  root: {
    position: 'fixed',
    top: '50%',
    left: '50%',
    transform: 'translateX(-50%)translateY(-50%)',
  },
}));

const Loader = ({ inline, interval }) => {
  const classes = useStyles();
  if (inline) {
    return <LoaderSvg interval={interval} />;
  } return (
    <div className={classes.root}>
      <LoaderSvg interval={interval} />
    </div>
  );
};

Loader.propTypes = {
  inline: PropTypes.bool,
  interval: PropTypes.number,
};

Loader.defaultProps = {
  inline: false,
  interval: 250,
};

export default Loader;
