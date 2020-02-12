import React from 'react';
import TimelineItem from './TimelineItem';
import mock from "./mock";

export default {
  title: 'UI|Time line/Timeline item',
};



export const story = () => <TimelineItem title={mock[0].title} time={mock[0].time} error={mock[0].error} content={mock[0].content}  />;
