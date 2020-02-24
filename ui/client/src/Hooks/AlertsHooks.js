import {
  useCallback,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';

export const useAlertsData = (provider, tags, deploymentTime) => {
  const { alerts } = useContext(SocketIOContext);
  const [data, setData] = useState();

  const onNewData = useCallback(({ data: newData, config }) => {
    // validate we use the relevant data
    // if (config.query === query && config.provider === provider) {
    setData(newData);
    // }
  }, [provider, tags]);

  useEffect(() => {
    alerts.on('data', onNewData);
    return () => {
      alerts.emit('close');
    };
  }, []);

  useEffect(() => {
    // setLoading(true);
    alerts.emit('init', tags, provider, deploymentTime);
  }, [tags, provider, deploymentTime]);

  return { data };
};
