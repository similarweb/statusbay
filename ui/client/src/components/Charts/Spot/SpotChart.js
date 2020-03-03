import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import Xrange from 'highcharts/modules/xrange';
import Box from '@material-ui/core/Box';
import PropTypes from 'prop-types';
import useTheme from '@material-ui/core/styles/useTheme';
import { addPlotLine } from '../utils';
import SpotChartConfig from './SpotChartConfig';
import * as moment from 'moment';

Xrange(Highcharts);

const SpotChart = ({ series, deploymentTime, currentTime }) => {
  const theme = useTheme();
  const isDarkMode = theme.palette.type === 'dark';
  let options = SpotChartConfig(series, deploymentTime);

  options = addPlotLine({
    color: theme.palette.error.main,
    value: deploymentTime,
    label: {
      align: 'center',
      rotation: 0,
      x: 0,
      y: -10,
      style: {
        color: isDarkMode ? '#ffffff' : '#999999',
      },
      formatter() {
        return `Deployment time: ${moment.unix(deploymentTime).format('HH:mm:ss')}`;
      },
    },
  }, options);

  options = addPlotLine({
    color: theme.palette.primary.main,
    value: currentTime,
    dashStyle: 'solid',
    width: 2,
    zIndex: 3,
    label: {
      align: 'center',
      rotation: 0,
      x: 0,
      y: -10,
      text: 'NOW',
    },
  }, options);
  return (
    <Box>
      <HighchartsReact
        highcharts={Highcharts}
        options={options}
      />
    </Box>
  );
};

SpotChart.propTypes = {
  deploymentTime: PropTypes.number,
  currentTime: PropTypes.number,
  series: PropTypes.arrayOf(PropTypes.shape({
    data: PropTypes.arrayOf(PropTypes.shape({
      form: PropTypes.number,
      to: PropTypes.number,
    })),
    name: PropTypes.string,
    link: PropTypes.string,
  })),
};

SpotChart.defaultProps = {
  deploymentTime: null,
  currentTime: null,
  series: [],
};

export default SpotChart;
