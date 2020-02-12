import { useHistory, useLocation } from 'react-router-dom';
import querystring from 'query-string';
import { useEffect, useMemo, useState } from 'react';

const useTableFilter = (name, valueTransformer, defaultValue) => {
  const location = useLocation();
  const history = useHistory();
  const params = querystring.parse(location.search);
  const [value, setValue] = useState(valueTransformer(params[name] || defaultValue));
  const handleChange = (newValue) => {
    setValue(newValue);
  };

  useEffect(() => {
    history.push({
      pathname: location.pathname,
      search: `?${new URLSearchParams({
        ...params,
        [name]: value,
      })}`,
    });
  }, [value]);

  return [value, handleChange];
};

const paramToArray = (param) => (param ? param.split(',') : []);


export const useTableFilterArray = (name) => useTableFilter(name, (value) => paramToArray(value));
// eslint-disable-next-line max-len
export const useTableFilterString = (name, defaultValue) => useTableFilter(name,
  (value) => value,
  defaultValue);
export const useTableFilterNumber = (name, defaultValue) => useTableFilter(name, (value) => {
  if (value || value === 0) {
    return parseInt(value);
  }
  return null;
}, defaultValue);

export const useTableFilters = (filtersConfig = {}) => {
  const location = useLocation();
  const history = useHistory();
  const params = querystring.parse(location.search);
  const defaultState = useMemo(() => {
    const result = {};
    Object.entries(filtersConfig).forEach(([name, config]) => {
      const { transformValue, defaultValue } = config;
      // try to get the initial value from the url
      const value = defaultValue;
      if (value !== null && value !== undefined && value !== 'null' && value !== 'undefined') {
        result[name] = value;
      }
    });
    return result;
  }, []);
  const initialState = useMemo(() => {
    const result = {};
    Object.entries(filtersConfig).forEach(([name, config]) => {
      const { transformValue, defaultValue } = config;
      // try to get the initial value from the url
      let value;
      if (typeof transformValue === 'function') {
        value = transformValue(params[name]) || defaultValue;
      } else {
        value = defaultValue;
      }
      if (value !== null && value !== undefined && value !== 'null' && value !== 'undefined') {
        result[name] = value;
      }
    });
    return result;
  }, []);
  const [state, setState] = useState(initialState);

  // update url according to the filters
  useEffect(() => {
    const getUrlParams = (urlParams) => {
      // filter empty params
      const filteredParams = Object.fromEntries(Object.entries(urlParams).filter(([name, value]) =>
        // const isDefaultValue = value === filtersConfig[name].defaultValue;
        value !== null && value !== undefined && (Array.isArray(value) ? value.length > 0 : true)));
      return new URLSearchParams(filteredParams);
    };

    const urlParams = {
      ...state,
    };
    history.push({
      pathname: location.pathname,
      search: `?${getUrlParams(urlParams)}`,
    });
  }, [state]);


  const handleChange = (name, value) => {
    setState({
      ...state,
      [name]: value,
    });
  };

  const reset = () => {
    setState(defaultState);
  };
  return [state, handleChange, reset];
};
