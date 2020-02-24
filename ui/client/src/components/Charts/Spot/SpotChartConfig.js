import * as moment from 'moment';
import useTheme from '@material-ui/core/styles/useTheme';
import numeral from 'numeral';

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

export default (series, deploymentTime) => {
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
          return moment.unix(this.value).utc().format('HH:mm:ss');
        },
        style: {
          color: isDarkMode ? '#ffffff' : '#666666',
        },
      },
      min: moment.unix(deploymentTime).utc().subtract(30, 'm').valueOf() / 1000,
      max: moment.unix(deploymentTime).utc().add(30, 'm').valueOf() / 1000,
      startOnTick: true,
      endOnTick: true,
    },
    yAxis: {
      title: null,
      categories: [],
      labels: {
        formatter() {
          return `<a style="color: ${theme.palette.primary.light}" href="${series[this.value].link}" target="_blank">${series[this.value].name}</a>`;
        },
      },
    },
    tooltip: {
      useHTML: true,
      formatter() {
        const downTime = numeral(moment.unix(this.x2).utc().diff(moment.unix(this.x).utc(), 'minutes')).format('0,0');
        return `<b>${series[this.y].name}</b><br /><b>Downtime:</b> ${downTime} minutes
<br /><b>From:</b> ${moment.unix(this.x).utc().format('DD/MM/YYYY HH:mm:ss')}
<br /><b>To:</b> ${moment.unix(this.x2).utc().format('DD/MM/YYYY HH:mm:ss')}`
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
      data: buildData(series, theme.palette.error.main),
      dataLabels: {
        enabled: true,
      },
    }],
  });
};
