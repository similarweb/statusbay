import React, {
  useCallback,
  useContext,
  useEffect, useState,
} from 'react';
import TableContainer from '@material-ui/core/TableContainer';
import Paper from '@material-ui/core/Paper';
import {
  Link, useHistory, useLocation,
} from 'react-router-dom';
import Toolbar from '@material-ui/core/Toolbar';
import Box from '@material-ui/core/Box';
import TableRow from '@material-ui/core/TableRow';
import TablePagination from '@material-ui/core/TablePagination';
import { useTranslation } from 'react-i18next';
import * as PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import MultiSelect from './Filters/MultiSelect';
import TableStateless from './TableStateless';
import CellStatus from './Cells/CellStatus';
import CellDeployBy from './Cells/CellDeployBy';
import CellTime from './Cells/CellTime';
import {
  useTableFilters,
} from '../../Hooks/TableHooks';
import { AppSettingsContext } from '../../context/AppSettingsContext';
import DatePickerFilter from './Filters/DatePickerFilter';
import SearchField from './Filters/SearchField';
import NoData from './NoData';
import { useApplicationsData } from '../../Hooks/ApplicationsHooks';

const parseSortBy = (sortby = '|') => sortby.split('|');
const paramToArray = (param = '') => (param ? param.split(',') : []);
const paramToNumber = (value) => {
  if (value || value === 0) {
    return parseInt(value);
  }
  return null;
};

const Table = ({ hideNameFilter, onRowClick, filters }) => {
  const { appSettings, dispatch } = useContext(AppSettingsContext);
  const [tableFilters, setTableFilters, resetTableFilters] = useTableFilters({
    cluster: {
      transformValue: paramToArray,
    },
    namespace: {
      transformValue: paramToArray,
    },
    status: {
      transformValue: paramToArray,
    },
    deployBy: {
      transformValue: (value) => value,
    },
    fromDate: {
      transformValue: paramToNumber,
      defaultValue: null,
    },
    toDate: {
      transformValue: paramToNumber,
      defaultValue: null,
    },
    name: {
      transformValue: (value) => value,
    },
    sortBy: {
      transformValue: (value) => value,
    },
    page: {
      transformValue: paramToNumber,
      defaultValue: 0,
    },
    rowsPerPage: {
      transformValue: paramToNumber,
      defaultValue: appSettings.rowsPerPage,
    },
    ...Object.fromEntries(Object.entries(filters).map(([name, value]) => {
      return [name, {
        defaultValue: value,
      }];
    })),
  });
  const { data: tableData, loading } = useApplicationsData(tableFilters);
  const { t } = useTranslation();
  const [sortByFiled, sortDirection] = parseSortBy(tableFilters.sortBy);

  const handleDateChange = (cb) => (date) => {
    const unix = date.format('x');
    cb(parseInt(unix));
  };

  const handleRowsPerPageChange = (event) => {
    dispatch({ type: 'SET_ROWS_PER_PAGE', rowsPerPage: event.target.value });
    setTableFilters('rowsPerPage', event.target.value);
  };

  const handlePageChange = (event, nextPage) => {
    setTableFilters('page', nextPage);
  };

  const resetFilters = () => {
    resetTableFilters();
  };

  const onSort = (id, direction) => {
    setTableFilters('sortBy', `${id}|${direction}`);
  };
  const tableConfig = {
    row: {
      render: (row, index) => ({ children }) => (
        <TableRow
          onClick={onRowClick(row)}
          hover
          key={`${row.name}-${index}`}
        >
          {children}
        </TableRow>
      ),
    },
    cells: [
      {
        id: 'name',
        name: t('table.filters.name'),
        cell: (row) => row.name,
        sortable: true,
      },
      {
        id: 'status',
        name: t('table.filters.status'),
        cell: (row) => <CellStatus status={row.status} />,
        sortable: true,
      },
      {
        id: 'cluster',
        name: t('table.filters.cluster'),
        cell: (row) => row.cluster,
        sortable: true,
      },
      {
        id: 'namepsace',
        name: t('table.filters.namespace'),
        cell: (row) => row.namespace,
        sortable: true,
      },
      {
        id: 'deployBy',
        name: t('table.filters.deploy.by'),
        cell: (row) => <CellDeployBy>{row.deployBy}</CellDeployBy>,
        sortable: true,
      },
      {
        id: 'time',
        name: t('table.columns.time'),
        cell: (row) => <CellTime time={row.time} />,
        sortable: true,
      },
      {
        id: 'details',
        name: '',
        cell: (row) => (
          <Link
            to={`/application/${row.name}/${row.time}`}
          >
            <Box display="flex" alignItems="center">
              <Button variant="outlined" color="primary">Details</Button>
            </Box>
          </Link>
        ),
        sortable: false,
      },
    ],
  };
  return (
    <Box m={2}>
      <Paper>
        <TableContainer component={Paper}>
          <Toolbar>
            <Box m={1} display="flex" flexDirection="column">
              <Box display="flex" alignItems="center">
                {
                  !hideNameFilter && (
                    <SearchField
                      label={t('table.filters.name')}
                      onChange={setTableFilters.bind(null, 'name')}
                      defaultValue={tableFilters.name}
                      delay={250}
                    />
                  )
                }
                <SearchField
                  label={t('table.filters.deploy.by')}
                  onChange={setTableFilters.bind(null, 'deployBy')}
                  defaultValue={tableFilters.deployBy}
                  delay={250}
                />
              </Box>
              <Box display="flex" alignItems="center">
                <MultiSelect
                  name={t('table.filters.cluster')}
                  onChange={setTableFilters.bind(null, 'cluster')}
                  selectedValue={tableFilters.cluster}
                  values={appSettings.filters.clusters}
                />
                <MultiSelect
                  name={t('table.filters.namespace')}
                  onChange={setTableFilters.bind(null, 'namespace')}
                  selectedValue={tableFilters.namespace}
                  values={appSettings.filters.namespaces}
                />
                <MultiSelect
                  name={t('table.filters.status')}
                  onChange={setTableFilters.bind(null, 'status')}
                  selectedValue={tableFilters.status}
                  values={appSettings.filters.statuses}
                />
                <DatePickerFilter
                  label={t('table.filters.from')}
                  value={tableFilters.from}
                  onChange={handleDateChange(setTableFilters.bind(null, 'from'))}
                />
                <DatePickerFilter
                  label={t('table.filters.to')}
                  value={tableFilters.to}
                  onChange={handleDateChange(setTableFilters.bind(null, 'to'))}
                />
                <Button variant="contained" color="secondary" onClick={resetFilters}>Reset</Button>
              </Box>
            </Box>
          </Toolbar>
          <TableStateless
            data={tableData && tableData.rows}
            config={tableConfig}
            page={parseInt(tableFilters.page)}
            loading={loading}
            sortBy={sortByFiled}
            sortDirection={sortDirection}
            onSort={onSort}
          />
          {
            !(loading || !tableData || !tableData.rows.length) && (
              <TablePagination
                rowsPerPageOptions={[20, 50, 100]}
                rowsPerPage={tableFilters.rowsPerPage}
                onChangeRowsPerPage={handleRowsPerPageChange}
                count={tableData && tableData.totalCount}
                page={tableFilters.page}
                onChangePage={handlePageChange}
              />
            )
          }
          {
            !loading && !tableData && <Box m={2}><NoData /></Box>
          }
        </TableContainer>
      </Paper>
    </Box>
  );
};
Table.propTypes = {
  hideNameFilter: PropTypes.bool,
  onRowClick: PropTypes.func,
  filters: PropTypes.object,
};
Table.defaultProps = {
  hideNameFilter: false,
  onRowClick: () => () => null,
  filters: {},
};
export default Table;
