import {
  useCallback,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';

export const useAlertsData = (provider, tags, deploymentTime) => {
  const { alerts } = useContext(SocketIOContext);
  const [data, setData] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();
  const [tagsWarning, setTagsWarning] = useState();

  const onNewData = useCallback(({ data: newData, config }) => {
    // validate we use the relevant data
    setData(newData);
    if (loading) {
      setLoading(false);
    }
  }, [provider, tags, loading]);

  const onAlertsError = useCallback(({ error, config }) => {
    setError(error);
    if (loading) {
      setLoading(false);
    }
  }, [provider, tags, loading]);

  const onTagsWarning = useCallback(({ tags }) => {
    setTagsWarning(tags);
    if (loading) {
      setLoading(false);
    }
  }, [provider, tags, loading]);

  useEffect(() => {
    alerts.on('data', onNewData);
    alerts.on('alerts-error', onAlertsError);
    alerts.on('alerts-tags', onTagsWarning);
    return () => {
      alerts.emit('close');
    };
  }, []);

  useEffect(() => {
    // setLoading(true);
    alerts.emit('init', tags, provider, deploymentTime);
  }, [tags, provider, deploymentTime]);

  return { data, loading, error, tagsWarning };
};
