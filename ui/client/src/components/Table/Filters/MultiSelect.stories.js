import React, { useState } from 'react';
import { text } from '@storybook/addon-knobs';
import MultiSelect from './MultiSelect';

export default {
  title: 'Table|Filters/Multi Select',
};
export const ToStory = () => {
  const [selectedValues, setSelectedValues] = useState([]);
  const values = [
    'option1',
    'option2',
    'option3',
    'option4',
    'option5',
  ];
  const handleChange = (value) => {
    setSelectedValues(value);
  };

  return <MultiSelect name={text('Name', 'Tag')} onChange={handleChange} selectedValue={selectedValues} values={values} />;
};

ToStory.story = 'Multi select stor';
