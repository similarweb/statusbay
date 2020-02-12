import React from 'react';
import PropTypes from 'prop-types';
import Tooltip from '@material-ui/core/Tooltip';
import moment from 'moment';
import { useTranslation } from 'react-i18next';

const CellTime = ({ time }) => {
  const timeObj = moment.unix(time);
  const formatted = timeObj.format('DD/MM/YYYY hh:mm:ss');
  const relative = timeObj.fromNow();
  const { t } = useTranslation();
  return <Tooltip title={t('table.columns.time.tooltip', { time: formatted })}><span>{relative}</span></Tooltip>;
};
CellTime.propTypes = {
  time: PropTypes.number.isRequired,
};

export default CellTime;
