import React, { useMemo } from 'react';
import PropTypes from 'prop-types';

const endpoint = process.env.NODE_ENV === 'development' ? 'localhost:5000' : '';


const SocketIOContext = React.createContext({});
const baseSocketConfig = {
  transports: ['websocket'],
  secure: true,
  path: '/api/socket',
};
const SocketIOProvider = ({ children, io }) => {
  const deploymentDetails = io(`${endpoint}/deployment-details`, {
    ...baseSocketConfig,
  });

  const applications = io(`${endpoint}/applications`, {
    ...baseSocketConfig,
  });

  const metrics = io(`${endpoint}/metrics`, {
    ...baseSocketConfig,
  });

  const alerts = io(`${endpoint}/alerts`, {
    ...baseSocketConfig,
  });

  const podLogs = io(`${endpoint}/pod-logs`, {
    ...baseSocketConfig,
  });

  const value = useMemo(() => ({
    deploymentDetails,
    applications,
    metrics,
    alerts,
    podLogs
  }), [deploymentDetails, applications, metrics, alerts, podLogs]);

  return (
    <SocketIOContext.Provider value={value}>
      {children}
    </SocketIOContext.Provider>
  );
};

SocketIOProvider.propTypes = {
  children: PropTypes.node.isRequired,
  io: PropTypes.func.isRequired,
};


export { SocketIOContext, SocketIOProvider };
