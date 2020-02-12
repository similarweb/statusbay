import React, { useMemo, useReducer } from 'react';
import * as PropTypes from 'prop-types';

const initialState = {
  rowsPerPage: 20,
};
const AppSettingsContext = React.createContext({});

const appSettingsReducer = (state, action) => {
  switch (action.type) {
    case 'SET_ROWS_PER_PAGE': {
      return {
        ...state,
        rowsPerPage: action.rowsPerPage,
      };
    }
    case 'SET_TABLE_FILTERS': {
      return {
        ...state,
        filters: {
          ...action.filters,
        },
      };
    }
    default: {
      return state;
    }
  }
};
const AppSettingsProvider = ({ children }) => {
  const [state, dispatch] = useReducer(appSettingsReducer, initialState);
  const contextValue = useMemo(() => ({
    appSettings: state,
    dispatch,
  }), [state, dispatch]);
  return (
    <AppSettingsContext.Provider value={contextValue}>
      {children}
    </AppSettingsContext.Provider>
  );
};

AppSettingsProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export { AppSettingsContext, AppSettingsProvider };
