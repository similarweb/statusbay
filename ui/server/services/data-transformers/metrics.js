const metricsTransformer = (data) => {
  return data.map(current => {
    return {
      metric: current.Metric,
      points: current.Points
    }
  })
}

module.exports = metricsTransformer
