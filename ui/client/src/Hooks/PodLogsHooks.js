import {
  useCallback,
  useContext, useEffect, useRef, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';

export const usePodLogs = (deploymentId, podName) => {
  const { podLogs } = useContext(SocketIOContext);
  const [data, setData] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  const onNewData = useCallback(({ data: newData, config }) => {
    // validate we use the relevant data
    if (config.deploymentId === deploymentId && config.podName === podName) {
      setData(newData);
      if (loading) {
        setLoading(false);
      }
    }
  }, []);

  const onPodLogsError = useCallback(({ error, config }) => {
    // validate we use the relevant data
    if (config.deploymentId === deploymentId && config.podName === podName) {
      setError(error);
      if (loading) {
        setLoading(false);
      }
    }
  }, []);

  useEffect(() => {
    podLogs.on('data', onNewData);
    podLogs.on('pod-logs-error', onPodLogsError);
    return () => {
      podLogs.emit('close');
    };
  }, []);

  useEffect(() => {
    // setLoading(true);
    podLogs.emit('init', deploymentId, podName);
  }, []);

  return { data, loading, error };
};
