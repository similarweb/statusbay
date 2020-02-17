import React from 'react';
import Typography from '@material-ui/core/Typography';
import PropTypes from 'prop-types';

const PageTitle = ({ children }) => (
  <Typography variant="h3">
    {children}
  </Typography>
);
PageTitle.propTypes = {
  children: PropTypes.node.isRequired,
};
export default PageTitle;
