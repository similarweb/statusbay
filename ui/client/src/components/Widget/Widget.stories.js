import React from 'react';
import { text } from '@storybook/addon-knobs';
import Widget from './Widget';

export default {
  title: 'UI|Widget',
};

export const story = () => <Widget title={text('Title', 'Widget Title')}><div>Content</div></Widget>;
