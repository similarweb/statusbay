import {
  useCallback,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';

export const useMetricsData = (provider, query, deploymentTime) => {
  const { metrics } = useContext(SocketIOContext);
  const [data, setData] = useState();

  const onNewData = useCallback(({ data: newData, config }) => {
    // validate we use the relevant data
    if (config.query === query && config.provider === provider) {
      setData(newData);
    }
  }, [provider, query]);

  useEffect(() => {
    metrics.on('data', onNewData);
    return () => {
      metrics.emit('close');
    };
  }, []);

  useEffect(() => {
    // setLoading(true);
    metrics.emit('init', query, provider, deploymentTime);
  }, [query, provider, deploymentTime]);

  return { data };
};
