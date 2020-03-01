import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableBody from '@material-ui/core/TableBody';
import Table from '@material-ui/core/Table';
import React from 'react';
import PropTypes from 'prop-types';
import TableCell from '@material-ui/core/TableCell';
import Skeleton from '@material-ui/lab/Skeleton';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import makeStyles from '@material-ui/core/styles/makeStyles';

const renderRows = (config, data, page) => data.map((row, rowIndex) => {
  const RowComponent = config.row.render(row, rowIndex);
  // eslint-disable-next-line max-len
  const rowCells = config.cells.map((cellConfig) => <TableCell>{cellConfig.cell(row, ((page) * data.length) + rowIndex)}</TableCell>);
  return (
    <RowComponent>
      {rowCells}
    </RowComponent>
  );
});

const renderLoadingState = (config) => [...Array(10).keys()].map((index) => (
  <TableRow>
    {
        config.cells.map((cell, cellIndex, cells) => {
          const isLast = cellIndex === cells.length - 1;
          if (!isLast) {
            return (
              <TableCell key={cell.id}>
                <Skeleton variant="rect" width="auto" height={27} />
              </TableCell>
            );
          }
          return null;
        })
      }
  </TableRow>
));

const TableStateless = (props) => {
  const {
    config, data, page, tableSize, loading, sortBy, sortDirection, onSort, stickyHeader,
  } = props;
  const onSortClick = (id) => () => {
    if (sortBy === id) {
      onSort(sortBy, sortDirection === 'asc' ? 'desc' : 'asc');
    } else {
      onSort(id, 'desc');
    }
  };
  return (
    <Table size={tableSize} stickyHeader={stickyHeader}>
      <TableHead>
        <TableRow>
          {
            config.cells.map(({
              id, name, sortable,
            }) => {
              if (sortable) {
                return <TableCell key={`table-header-cell-${name}`}><TableSortLabel active={sortBy === id} direction={sortDirection} onClick={onSortClick(id)}>{name}</TableSortLabel></TableCell>;
              }
              return <TableCell key={`table-header-cell-${name}`}>{name}</TableCell>;
            })
          }
        </TableRow>
      </TableHead>
      <TableBody>
        {loading ? renderLoadingState(config) : renderRows(config, data, page)}
      </TableBody>
    </Table>
  );
};

const cellPropTypes = {
  id: PropTypes.is,
  name: PropTypes.string,
  cell: PropTypes.func,
  sortable: PropTypes.bool,
};

TableStateless.propTypes = {
  data: PropTypes.arrayOf(PropTypes.any),
  config: PropTypes.shape({
    cells: PropTypes.arrayOf(PropTypes.shape(cellPropTypes)),
    row: PropTypes.shape({
      render: PropTypes.func,
    }),
  }),
  page: PropTypes.number,
  tableSize: PropTypes.string,
  loading: PropTypes.bool,
  sortBy: PropTypes.string,
  sortDirection: PropTypes.oneOf(['asc', 'desc']),
  onSort: PropTypes.func,
  stickyHeader: PropTypes.bool,
};

TableStateless.defaultProps = {
  data: [],
  config: {
    cells: [],
    row: { render: () => null },
  },
  page: 0,
  tableSize: 'small',
  loading: false,
  sortBy: null,
  sortDirection: 'desc',
  onSort: () => null,
  stickyHeader: false,
};

export default React.memo(TableStateless);
