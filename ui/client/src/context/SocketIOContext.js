import React, { useMemo } from 'react';
import PropTypes from 'prop-types';

const endpoint = process.env.NODE_ENV === 'development' ? 'localhost:5000' : '';


const SocketIOContext = React.createContext({});
const baseSocketConfig = {
  transports: ['websocket'],
  secure: true,
};
const SocketIOProvider = ({ children, io }) => {
  const deploymentDetails = io(`${endpoint}/deployment-details`, {
    ...baseSocketConfig,
    path: '/api/socket',
  });

  const applications = io(`${endpoint}/applications`, {
    ...baseSocketConfig,
    path: '/api/socket',
  });

  const metrics = io(`${endpoint}/metrics`, {
    ...baseSocketConfig,
    path: '/api/socket',
  });

  const alerts = io(`${endpoint}/alerts`, {
    ...baseSocketConfig,
    path: '/api/socket',
  });

  const value = useMemo(() => ({
    deploymentDetails,
    applications,
    metrics,
    alerts
  }), [deploymentDetails, applications, metrics, alerts]);

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
