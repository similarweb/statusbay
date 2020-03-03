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
import Box from '@material-ui/core/Box';
import Loader from '../Loader/Loader';

const renderRows = (config, data, page) => data.map((row, rowIndex) => {
  const RowComponent = config.row.render(row, rowIndex);
  // eslint-disable-next-line max-len
  const rowCells = config.cells.map((cellConfig) => <TableCell style={cellConfig.cellStyle}>{cellConfig.cell(row, ((page) * data.length) + rowIndex)}</TableCell>);
  return (
    <RowComponent>
      {rowCells}
    </RowComponent>
  );
});

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
  if (loading) {
    return <Box display="flex" justifyContent="space-around"><Loader inline={true} interval={100} /></Box>;
  }
  return (
    <Table size={tableSize} stickyHeader={stickyHeader}>
      <TableHead>
        <TableRow>
          {
            config.cells.map(({
              id, name, sortable, width
            }) => {
              let content = name;
              if (sortable) {
                content = <TableSortLabel active={sortBy === id} direction={sortDirection} onClick={onSortClick(id)}>{name}</TableSortLabel>;
              }
              return <TableCell style={{width}} key={`table-header-cell-${name}`}>{content}</TableCell>;
            })
          }
        </TableRow>
      </TableHead>
      <TableBody>
        {renderRows(config, data, page)}
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
