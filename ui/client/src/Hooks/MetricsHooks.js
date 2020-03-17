import {
  useCallback,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';

export const useMetricsData = (provider, query, deploymentTime) => {
  const { metrics } = useContext(SocketIOContext);
  const [data, setData] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  const onNewData = useCallback(({ data: newData, config }) => {
    // validate we use the relevant data
    if (config.query === query && config.provider === provider) {
      setData(newData);
      if (loading) {
        setLoading(false);
      }
    }
  }, [provider, query, loading]);

  const onMetricError = useCallback(({ error, config }) => {
    // validate we use the relevant data
    if (config.query === query && config.provider === provider) {
      setError(error);
      if (loading) {
        setLoading(false);
      }
    }
  }, [provider, query, loading]);

  useEffect(() => {
    metrics.on('data', onNewData);
    metrics.on('metric-error', onMetricError);
    return () => {
      metrics.emit('close');
    };
  }, []);

  useEffect(() => {
    // setLoading(true);
    metrics.emit('init', query, provider, deploymentTime);
  }, [query, provider, deploymentTime]);

  return { data, loading, error };
};
