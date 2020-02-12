import React from 'react';
import TimelineErrorBox from './TimelineErrorBox';
import mock from './mock';

export default {
  title: 'UI|Time line/Timeline error box',
};


export const story = () => (
  <TimelineErrorBox>
    {
    mock[0].content
  }
  </TimelineErrorBox>
);
