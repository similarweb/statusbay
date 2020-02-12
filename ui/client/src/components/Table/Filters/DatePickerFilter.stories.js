import React, { useState } from 'react';
import { action } from '@storybook/addon-actions';
import { select } from '@storybook/addon-knobs';
import DatePickerFilter from './DatePickerFilter';

export default {
  title: 'UI|Filters/Date Picker',
};

export const ToStory = () => {
  const label = 'Formats';
  const options = {
    'DD/MM/YYYY': 'DD/MM/YYYY',
    'MM/DD/YYYY': 'MM/DD/YYYY',
  };
  const defaultValue = 'DD/MM/YYYY';
  const groupId = 'format';

  const format = select(label, options, defaultValue, groupId);
  const [value, setValue] = useState(1577535974469);
  const onChange = (date, newValue) => {
    action('Change(Value)')(newValue);
    action('Change(Moment)')(date);
    setValue(parseInt(date.format('x')));
  };
  return <DatePickerFilter value={value} label="Date" format={format} onChange={onChange} />;
};
