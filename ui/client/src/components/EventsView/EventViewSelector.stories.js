import React, { useState } from 'react';
import EventsViewSelector from './EventViewSelector';

export default {
  title: 'UI|Event View Selector',
};

const items = [
  { name: 'pod1', status: 'failed' },
  { name: 'pod2', status: 'running' },
];

const DefaultStory = () => {
  const [selected, setSelected] = useState(0);
  const handleChange = (index) => () => {
    setSelected(index);
  };
  return <EventsViewSelector items={items} selected={selected} onRowClick={handleChange} />;
};


export const story = () => <DefaultStory />;
story.story = {
  name: 'default',
};
