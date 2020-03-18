import {
  useCallback,
  useContext, useEffect, useRef, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';
import { transformTableData } from '../Services/API/TableApi';

const parseSortBy = (sortby = '|') => sortby.split('|');

const prepareFilters = (filters) => {
  const {
    name, exactName, deployBy, cluster, namespace, status, fromDate, toDate, page, rowsPerPage, sortBy, distinct
  } = filters;
  const [sortByFiled, sortDirection] = parseSortBy(sortBy);
  return {
    name,
    exactName,
    deployBy,
    cluster,
    nameSpace: namespace,
    status,
    from: fromDate,
    to: toDate,
    rowsPerPage,
    page,
    sortBy: sortByFiled,
    sortDirection,
    distinct
  };
};

export const useApplicationsData = (filters) => {
  const { applications } = useContext(SocketIOContext);
  const [data, setData] = useState();
  const [loading, setLoading] = useState(true);
  const prevHashValue = useRef('');
  const socket = useRef(null);
  const onNewData = useCallback(({ data: newData, hashValue }) => {
    // check if data changed
    if (hashValue !== prevHashValue.current) {
      prevHashValue.current = hashValue;
      setData(transformTableData(newData.Records, newData.Count));
    }
    setLoading(false);
  }, []);
  useEffect(() => {
    socket.current = applications.on('data', onNewData);
    return () => {
      applications.emit('close');
      socket.current.removeAllListeners('data');
    };
  }, []);
  useEffect(() => {
    setLoading(true);
    applications.emit('filters', prepareFilters(filters));
  }, [filters]);
  return { data, loading };
};
