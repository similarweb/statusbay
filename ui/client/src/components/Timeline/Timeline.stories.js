import React from 'react';
import Timeline from './Timeline';
import mock from './mock';

export default {
  title: 'UI|Time line/Timeline',
};


export const story = () => {return (<Timeline items={mock} />);}
