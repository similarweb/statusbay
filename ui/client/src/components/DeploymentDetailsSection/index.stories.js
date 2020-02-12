import React from 'react';
import Typography from '@material-ui/core/Typography';
import DeploymentDetailsSection from './index';

export default {
  title: 'UI|Deployment Detail Section',
};
export const Story = () => (
  <DeploymentDetailsSection defaultExpanded title="Section Title">
    <Typography variant="h3">Content</Typography>
  </DeploymentDetailsSection>
);

Story.story = {
  name: 'default',
};
