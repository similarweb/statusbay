import React from 'react';
import { radios } from '@storybook/addon-knobs';
import Box from '@material-ui/core/Box';
import DeploymentStatusBox from './index';

export default {
  title: 'UI|Deployment status box',
};
export const Story = () => {
  const label = 'Status';
  const options = {
    Successful: 'successful',
    Failed: 'failed',
    Running: 'running',
    Timeout: 'timeout',
  };
  const defaultValue = 'successful';

  const value = radios(label, options, defaultValue);
  return <Box m={2}><DeploymentStatusBox status={value} /></Box>;
};

Story.story = {
  name: 'default',
};
