import React, { useMemo } from 'react';
import PropTypes from 'prop-types';

const endpoint = process.env.NODE_ENV === 'development' ? 'localhost:5000' : '';


const SocketIOContext = React.createContext({});

const SocketIOProvider = ({ children, io }) => {
  const deploymentDetails = io(`${endpoint}/deployment-details`, {
    path: '/api/socket',
  });

  const applications = io(`${endpoint}/applications`, {
    path: '/api/socket',
  });

  const value = useMemo(() => ({
    deploymentDetails,
    applications,
  }), [deploymentDetails, applications]);

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
