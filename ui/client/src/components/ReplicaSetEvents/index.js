import PropTypes from 'prop-types';
import React from 'react';
import Timeline from '../Timeline/Timeline';

const ReplicaSetEvents = ({ logs }) => <Timeline items={logs} maxHeight={400} />;

ReplicaSetEvents.propTypes = {
  logs: PropTypes.arrayOf(PropTypes.shape({
    title: PropTypes.string,
    time: PropTypes.number,
    error: PropTypes.bool,
    content: PropTypes.string,
  })),
};
ReplicaSetEvents.defaultProps = {
  logs: [],
};

export default ReplicaSetEvents;
