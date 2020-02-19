import React from 'react';
import { text } from '@storybook/addon-knobs';
import { action } from '@storybook/addon-actions';
import SearchField from './SearchField';

export default {
  title: 'Table|Filters/Search Field',
};
export const ToStory = () => <SearchField label={text('Label', 'Search')} onChange={action('Change')} defaultValue="" />;

ToStory.story = 'Multi select stor';
