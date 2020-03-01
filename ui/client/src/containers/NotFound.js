import React from 'react';
import NoData from '../components/Table/NoData.js'
import makeStyles from '@material-ui/core/styles/makeStyles';
import { Link } from 'react-router-dom';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme) => ({
  main: {
    position: 'fixed',
    top: '50%',
    left: '50%',
    transform: 'translateX(-50%)translateY(-50%)',
  },
  goto: {
    "text-align": "center"
  },
}));


const NotFound = () => {
  const classes = useStyles();
  return (
    <div className={classes.main}>
      <NoData imageWidth={250} message="404 PAGE NOT FOUND" />
      <div className={classes.goto}>
        <Link to={`/`} >
            <Button variant="contained" color="primary">HOME</Button>
        </Link>
      </div>

    </div>
  );
};

export default NotFound;
