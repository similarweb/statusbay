import * as moment from 'moment';
import useTheme from '@material-ui/core/styles/useTheme';
import numeral from 'numeral';
import {
  amber,
  blue,
  brown,
  cyan,
  deepPurple,
  lightGreen,
  lime,
  teal
} from '@material-ui/core/colors';

const createTooltipContent = (points) => points.map(({ color, y, series: { name } }) => `<br><span style="color:${color}">●</span> ${name}: <b>${numeral(y).format('0,0')}</b>`);

const colors = [
  deepPurple[500],
  blue[500],
  cyan[500],
  teal[500],
  lightGreen[500],
  lime[500],
  amber[500],
  brown[500],
];

const darkColors = [
  deepPurple[300],
  blue[300],
  cyan[300],
  teal[300],
  lightGreen[300],
  lime[300],
  amber[300],
  brown[300],
];

export default (series, plotlines) => {
  const theme = useTheme();
  const isDarkMode = theme.palette.type === 'dark';
  return ({
    chart: {
      marginTop: 20,
      zoomType: 'x',
      backgroundColor: isDarkMode ? '#424242' : '#ffffff',
    },
    colors: isDarkMode ? darkColors : colors,
    title: {
      text: null,
    },
    legend: {
      align: 'left',
      layout: 'horizontal',
      useHTML: true,
      itemStyle: {
        color: isDarkMode ? '#ffffff' : '#333333',
      }
    },
    xAxis: {
      min: parseInt(moment.unix(plotlines[0]).subtract(30, 'm').valueOf()),
      max: parseInt(moment.unix(plotlines[0]).add(30, 'm').valueOf()),
      labels: {
        formatter() {
          return moment(this.value).format('HH:mm:ss');
        },
        style: {
          color: isDarkMode ? '#ffffff' : '#666666',
        },
        startOnTick: true,
        endOnTick: true
      },
      plotLines: plotlines.map((line) => ({
        color: theme.palette.error.main,
        dashStyle: 'dash',
        value: line * 1000,
        width: 1,
        label: {
          align: 'center',
          rotation: 0,
          x: 0,
          y: -10,
          style: {
            color: isDarkMode ? '#ffffff' : '#999999',
          },
          formatter() {
            return `Deployment time: ${moment(line).format('HH:mm:ss')}`;
          },
        },
      })),
    },
    yAxis: {
      title: {
        enabled: false,
      },
      min: 0,
      gridLineColor: '#f5f5f5',
      labels: {
        style: {
          color: isDarkMode ? '#ffffff' : '#666666',
        }
      }
    },
    tooltip: {
      shared: true,
      crosshairs: true,
      useHTML: true,
      formatter() {
        return `<span style="font-size: 10px">${moment(this.x).format('DD/MM/YYYY HH:mm:ss')}</span>${createTooltipContent(this.points)}`;
      },
    },
    plotOptions: {
      line: {
        marker: {
          enabled: false,
        },
        lineWidth: 1,
        states: {
          hover: {
            lineWidth: 1,
            enabled: false,
          },
        },
        threshold: null,
      },
      series: {
        marker: {
          symbol: 'circle',
          states: {
            hover: {
              enabled: false,
            },
          },
        },

      },
    },
    series,
    defs: {
      gradient0: {
        tagName: 'linearGradient',
        id: 'gradient-0',
        x1: 0,
        y1: 0,
        x2: 0,
        y2: 1,
        children: [{
          tagName: 'stop',
          offset: 0,
        }, {
          tagName: 'stop',
          offset: 1,
        }],
      },
      gradient1: {
        tagName: 'linearGradient',
        id: 'gradient-1',
        x1: 0,
        y1: 0,
        x2: 0,
        y2: 1,
        children: [{
          tagName: 'stop',
          offset: 0,
        }, {
          tagName: 'stop',
          offset: 1,
        }],
      },
    },
  });
};
