import React from 'react';
import { radios } from '@storybook/addon-knobs';
import CellStatus from './CellStatus';

export default {
  title: 'Table|Cells/Status',
};

export const CellStatusStory = () => {
  const label = 'Status';
  const options = {
    Successful: 'successful',
    Failed: 'failed',
    Running: 'running',
    Timeout: 'timeout',
    Deleted: 'deleted',
    Cancelled: 'cancelled',
  };
  const defaultValue = 'successful';

  const value = radios(label, options, defaultValue);
  return <CellStatus status={value} />;
};

CellStatusStory.story = 'Status Cell';
