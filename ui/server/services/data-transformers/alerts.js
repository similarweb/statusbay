const alertsTransformer = (data) => {
  const result = [];
  data.forEach(({ID, URL, Name, Periods = []}) => {
    // ignore items with empty periods
    if (Periods.lenth === 0) {
      return;
    }
    const periodsWithDownStatus = Periods.filter(period => period.Status === 'Down');
    // we need only period with status = down
    if (periodsWithDownStatus.length === 0) {
      return;
    }
    result.push({
      name: Name,
      url: URL,
      periods: periodsWithDownStatus.map(({Status, StartUnix, EndUnix}) => {
        return {
          status: Status,
          start: StartUnix,
          end: EndUnix
        }
      })
    })
  });
  return result;
}

module.exports = alertsTransformer
