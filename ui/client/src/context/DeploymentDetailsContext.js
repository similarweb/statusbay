import React, {
  createContext,
  useContext, useEffect, useMemo, useRef, useState,
} from 'react';
import { SocketIOContext } from './SocketIOContext';

const DeploymentDetailsContext = createContext(null);
export const DeploymentDetailsContextProvider = ({ id, children }) => {
  const { deploymentDetails } = useContext(SocketIOContext);
  const [data, setData] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const prevHashValue = useRef('');
  const socket = useRef(null);
  useEffect(() => {
    deploymentDetails.emit('init', id);
    socket.current = deploymentDetails.on('data', ({ data: newData, hashValue }) => {
      // check if data changed
      if (hashValue !== prevHashValue.current) {
        prevHashValue.current = hashValue;
        setData(newData);
      }
      if (loading) {
        setLoading(false);
      }
    });
    deploymentDetails.on('not-found', ({ error }) => {
      setError(error);
    });
    return () => {
      deploymentDetails.emit('close');
      socket.current.removeAllListeners('data');
      socket.current.removeAllListeners('not-found');
    };
  }, []);
  const value = useMemo(() => ({ data, loading, error }), [data, loading, error]);
  return (
    <DeploymentDetailsContext.Provider value={value}>
      {children({ loading, data, error })}
    </DeploymentDetailsContext.Provider>
  );
};
export const useDeploymentDetailsContext = () => useContext(DeploymentDetailsContext);
