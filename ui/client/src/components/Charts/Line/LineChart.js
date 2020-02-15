import React from 'react';
import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import PropTypes from 'prop-types';
import {
  blue, cyan, deepPurple, teal,
} from '@material-ui/core/colors';
import LineChartConfig from './LineChartConfig';

Highcharts.setOptions({
  chart: {
    height: 250,
    style: {
      fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    },
  },
  legend: {
    itemStyle: {
      fontWeight: 'normal',
    },
  },
});

const colors = [
  deepPurple[500],
  blue[500],
  cyan[500],
  teal[500],
];

const fillColor = (color) => ({
  fillColor: {
    linearGradient: {
      x1: 0,
      y1: 0,
      x2: 0,
      y2: 1,
    },
    stops: [
      [0.2, color],
      [1, Highcharts.Color(color).setOpacity(0.6).get('rgba')],
    ],
  },
});

const buildSeries = (series) => series.map((serie, index) => ({
  name: serie.name,
  data: serie.points,
  showInLegend: true,
  color: colors[index],
  lineWidth: 1,
  states: {
    hover: {
      enabled: false,
    },
  },
  ...fillColor(colors[index]),
}));

const LineChart = ({ series, deploymentTime }) => {
  const options = LineChartConfig(buildSeries(series), [deploymentTime]);

  return (
    <HighchartsReact
      highcharts={Highcharts}
      options={options}
    />
  );
};

LineChart.propTypes = {
  series: PropTypes.arrayOf(PropTypes.shape({
    points: PropTypes.array,
    name: PropTypes.string,
  })).isRequired,
  deploymentTime: PropTypes.number,
};

LineChart.defaultProps = {
  deploymentTime: null,
};

export default LineChart;
