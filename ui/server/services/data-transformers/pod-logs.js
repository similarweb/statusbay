const podLogsTransformer = (data = []) => {
  return data.map(current => {
    return {
      name: current.ContainerName,
      logs: current.Lines
    }
  })
}

module.exports = podLogsTransformer
