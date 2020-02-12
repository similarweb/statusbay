import * as moment from 'moment';
import useTheme from '@material-ui/core/styles/useTheme';

export default (series, plotlines) => {
  const theme = useTheme();
  const isDarkMode = theme.palette.type === 'dark';
  return ({
    chart: {
      zoomType: 'x',
      backgroundColor: isDarkMode ? '#424242' : '#ffffff',
    },
    title: {
      text: null,
    },
    legend: {
      align: 'left',
      layout: 'horizontal',
      useHTML: true,
    },
    xAxis: {
      labels: {
        formatter() {
          return moment.unix(this.value).format('HH:MM:SS');
        },
      },
      plotLines: plotlines.map((line) => ({
        color: theme.palette.error[theme.palette.type],
        dashStyle: 'dash',
        value: line,
        width: 1,
      })),
    },
    yAxis: {
      title: {
        enabled: false,
      },
      min: 0,
      gridLineColor: '#f5f5f5',
    },
    tooltip: {
      shared: true,
      crosshairs: true,
      // backgroundColor: 'red',
      useHTML: true,
      // formatter() {
      //   return moment.unix(this.x).format('HH:MM:SS');
      // },
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
