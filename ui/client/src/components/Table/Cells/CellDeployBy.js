import React from 'react';
import PropTypes from 'prop-types';
import Tooltip from '@material-ui/core/Tooltip';
import { useTranslation } from 'react-i18next';
import Typography from '@material-ui/core/Typography';

const CellDeployBy = ({ children }) => {
  const { t } = useTranslation();
  return <Tooltip title={t('table.columns.deploy.by.tooltip', { name: children })}><Typography variant="body2" component="span">{children}</Typography></Tooltip>;
};
CellDeployBy.propTypes = {
  children: PropTypes.string.isRequired,
};

export default CellDeployBy;
