import React from 'react';
import PropTypes from 'prop-types';
import Tooltip from '@material-ui/core/Tooltip';
import { useTranslation } from 'react-i18next';

const CellDeployBy = ({ children }) => {
  const { t } = useTranslation();
  return <Tooltip title={t('table.columns.deploy.by.tooltip', { name: children })}><span>{children}</span></Tooltip>;
};
CellDeployBy.propTypes = {
  children: PropTypes.string.isRequired,
};

export default CellDeployBy;
