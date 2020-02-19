import React from 'react';
import Box from '@material-ui/core/Box';
import PropTypes from 'prop-types';
import Breadcrumbs from '../Breadcrumbs/index';

const PageContent = ({ children }) => (
  <div>
    {/*<Box m={1}><Breadcrumbs /></Box>*/}
    {children}
  </div>
);
PageContent.propTypes = {
  children: PropTypes.node.isRequired,
};

export default PageContent;
