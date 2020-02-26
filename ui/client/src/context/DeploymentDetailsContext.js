import React, {
  createContext,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from './SocketIOContext';

const DeploymentDetailsContext = createContext(null);
export const DeploymentDetailsContextProvider = ({ id, children }) => {
  const { deploymentDetails } = useContext(SocketIOContext);
  const [data, setData] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    deploymentDetails.emit('init', id);
    deploymentDetails.on('data', ({ data: newData }) => {
      setData(newData);
      if (loading) {
        setLoading(false);
      }
    });
    deploymentDetails.on('not-found', ({ error }) => {
      setError(error);
    });
    return () => {
      deploymentDetails.emit('close');
    };
  }, [id, loading]);
  return (
    <DeploymentDetailsContext.Provider value={{data, loading, error}}>
      {children({loading, data, error})}
    </DeploymentDetailsContext.Provider>
  );
};
export const useDeploymentDetailsContext = () => useContext(DeploymentDetailsContext);
