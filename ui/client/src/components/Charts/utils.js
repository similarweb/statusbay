/* eslint-disable import/prefer-default-export */
export const addPlotLine = (plotLineConfig, config) => {
  const {
    dashStyle = 'dash', width = 1, ...rest
  } = plotLineConfig;
  const plotLines = config.xAxis.plotLines || [];
  return {
    ...config,
    xAxis: {
      ...config.xAxis,
      plotLines: [
        ...plotLines,
        {
          dashStyle,
          width,
          ...rest,
        },
      ],
    },
  };
};
