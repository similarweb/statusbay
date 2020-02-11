import React, {
  createContext,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from './SocketIOContext';

const DeploymentDetailsContext = createContext(null);
export const DeploymentDetailsContextProvider = ({ id, children }) => {
  const { deploymentDetails } = useContext(SocketIOContext);
  const [data, setData] = useState();
  useEffect(() => {
    deploymentDetails.emit('init', id);
    deploymentDetails.on('data', ({ data: newData }) => {
      setData(newData);
    });
    return () => {
      deploymentDetails.emit('close');
    };
  }, [id]);
  return (
    <DeploymentDetailsContext.Provider value={data}>
      {children}
    </DeploymentDetailsContext.Provider>
  );
};
export const useDeploymentDetailsContext = () => useContext(DeploymentDetailsContext);
