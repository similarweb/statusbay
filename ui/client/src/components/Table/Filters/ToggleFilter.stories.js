import React from 'react';
import { text, boolean } from '@storybook/addon-knobs';
import { action } from '@storybook/addon-actions';
import ToggleFilter from './ToggleFilter';

export default {
  title: 'UI|Filters/Toggle Filter',
};
export const ToStory = () => {
  const [state, setState] = React.useState(false);
  const handleChange = () => {
    setState(!state);
  }
  return <ToggleFilter label={text('Label', 'example')} onChange={handleChange} checked={state} />;
}

ToStory.story = 'Toggle filter story';
