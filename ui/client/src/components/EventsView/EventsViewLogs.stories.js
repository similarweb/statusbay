import React from 'react';
import EventsViewLogs from './EventsViewLogs';
import mock from './logsMock';

export default {
  title: 'UI|Event View Logs',
};

export const Story = () => <EventsViewLogs logs={mock} />;

Story.story = {
  name: 'default',
};
