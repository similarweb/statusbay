import * as moment from 'moment';
import useTheme from '@material-ui/core/styles/useTheme';

const buildData = (series, color) => series.reduce((result, serie, index) => {
  serie.data.forEach((point) => {
    result.push({
      color,
      x: point.from,
      x2: point.to,
      y: index,
    });
  });
  return result;
}, []);

export default (series) => {
  const theme = useTheme();
  const isDarkMode = theme.palette.type === 'dark';
  return ({
    chart: {
      marginTop: 20,
      type: 'xrange',
      zoomType: 'x',
      backgroundColor: isDarkMode ? '#424242' : '#ffffff',
    },
    title: {
      text: null,
    },
    xAxis: {
      labels: {
        formatter() {
          return moment.unix(this.value).format('HH:MM:SS');
        },
      },
    },
    yAxis: {
      title: null,
      categories: [],
      reversed: true,
      labels: {
        useHTML: true,
        formatter() {
          return `<a href="${series[this.value].link}">${series[this.value].name}</a>`;
        },
      },
    },
    tooltip: {
      useHTML: true,
      formatter() {
        return `<b>${series[this.y].name}</b><br /><b>Downtime:</b> ${moment.unix(this.x2).diff(moment.unix(this.x), 'minutes')} minutes`;
      },
    },
    series: [{
      showInLegend: false,
      pointPadding: 1,
      groupPadding: 1,
      pointWidth: 20,
      borderWidth: 0,
      borderColor: 'transparent',
      borderRadius: 5,
      states: {
        hover: {
          enabled: false,
        },
      },
      label: {
        enabled: false,
      },
      data: buildData(series, theme.palette.error[theme.palette.type]),
      dataLabels: {
        enabled: true,
      },
    }],
  });
};
