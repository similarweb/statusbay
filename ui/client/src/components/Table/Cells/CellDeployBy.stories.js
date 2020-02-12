import React from 'react';
import { text } from '@storybook/addon-knobs';
import CellDeployBy from './CellDeployBy';

export default {
  title: 'Table|Cells/Deploy By',
};
export const CellDeployByStory = () => <CellDeployBy>{text('username', 'username@company.com')}</CellDeployBy>;

CellDeployByStory.story = 'Cell Deploy By';
