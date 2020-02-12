import React from 'react';
import TimelineErrorBox from './TimelineErrorBox';
/* eslint-disable react/jsx-no-target-blank */

export default {
  title: 'UI|Time line/Timeline error box',
};


export const story = () => (
  <TimelineErrorBox>
        The pod not pass the hatcheck:
    <ul>
      <li>
Indicates whether the Container is ready to
                service requests. If the readiness probe fails,
                the endpoints controller removes the Pod’s IP address
                from the endpoints of all Services that match the Pod.
      </li>
      <li>
        <a
          href="https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate"
          target="_blank"
        >
Read more on Readiness details
        </a>
      </li>
    </ul>
  </TimelineErrorBox>
);
