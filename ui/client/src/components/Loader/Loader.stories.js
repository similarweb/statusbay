import React from 'react';
import { boolean } from '@storybook/addon-knobs';
import Loader from './Loader';

export default {
  title: 'UI|Loader',
};


export const Story = () => <Loader inline={boolean('inline', false)} />;
Story.story = { name: 'Loader' };
