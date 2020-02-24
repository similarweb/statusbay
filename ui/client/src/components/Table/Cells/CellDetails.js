import Tooltip from '@material-ui/core/Tooltip';
import Button from '@material-ui/core/Button';
import { Link } from 'react-router-dom';
import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    opacity: 0,
  },
}));

const CellDetails = ({ row, className }) => {
  const classes = useStyles();
  const content = row.status === 'deleted' ? (
    <Tooltip title="Details are unavailable for deleted deployments">
      <span>
        <Button variant="contained" color="primary" disabled={true}>Details</Button>
      </span>
    </Tooltip>
  )
    : (
      <Link
        to={`/application/${row.id}`}
      >
        {' '}
        <Button variant="contained" color="primary">Details</Button>

      </Link>
    );

  return <div className={`${className} ${classes.root}`}>{content}</div>;
};

CellDetails.propTypes = {
  row: PropTypes.shape({
    status: PropTypes.string.isRequired,
  }).isRequired,
};


export default CellDetails;
