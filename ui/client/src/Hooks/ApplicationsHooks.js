import {
  useCallback,
  useContext, useEffect, useState,
} from 'react';
import { SocketIOContext } from '../context/SocketIOContext';
import { transformTableData } from '../Services/API/TableApi';

const parseSortBy = (sortby = '|') => sortby.split('|');

const prepareFilters = (filters) => {
  const {
    name, deployBy, cluster, namespace, status, fromDate, toDate, page, rowsPerPage, sortBy, distinct
  } = filters;
  const [sortByFiled, sortDirection] = parseSortBy(sortBy);
  return {
    name,
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
  const onNewData = useCallback(({ data: newData }) => {
    setData(transformTableData(newData.Records, newData.Count));
    setLoading(false);
  }, []);
  useEffect(() => {
    applications.on('data', onNewData);
    return () => {
      applications.emit('close');
    };
  }, []);
  useEffect(() => {
    setLoading(true);
    applications.emit('filters', prepareFilters(filters));
  }, [filters]);
  return { data, loading };
};
