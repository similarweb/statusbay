import React, { useMemo } from 'react';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import makeStyles from '@material-ui/core/styles/makeStyles';
import PropTypes from 'prop-types';
import Box from '@material-ui/core/Box';
import TableStateless from '../Table/TableStateless';

const useStyles = makeStyles((theme) => ({
  container: {
    maxHeight: 437,
    overflowY: 'auto',
  },
  selected: {
    borderLeft: `3px solid ${theme.palette.primary[theme.palette.type]}`,
  },
  hover: {
    cursor: 'pointer',
  },
  marker: {
    width: 10,
    height: 10,
    backgroundColor: `${theme.palette.error.main}`,
    borderRadius: '50%',
    marginRight: 12,
  },
}));


const EventsViewSelector = ({ items, selected, onRowClick }) => {
  const classes = useStyles();
  const tableConfig = useMemo(() => ({
    row: {
      render: (row, rowIndex) => ({ children }) => (
        <TableRow
          key={rowIndex}
          hover
          classes={{ hover: classes.hover, selected: classes.selected }}
          onClick={onRowClick(rowIndex)}
          selected={rowIndex === selected}
        >
          {children}
        </TableRow>
      ),
    },
    cells: [
      {
        name: 'Pod',
        header: (name) => <TableCell>{name}</TableCell>,
        cell: (row) => (
          <Box display="flex" alignItems="center">
            <div className={row.hasError ? classes.marker : undefined} />
            {' '}
            {row.name}
          </Box>
        ),
      },
      {
        name: 'Status',
        header: (name) => <TableCell>{name}</TableCell>,
        cell: (row) => row.status,
      },
    ],
  }), [selected]);
  return (
    <div className={classes.container}>
      <TableStateless data={items} config={tableConfig} tableSize="small" stickyHeader={false} />
    </div>
  );
};

EventsViewSelector.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string.isRequired,
    status: PropTypes.oneOf(['successful', 'failed', 'running', 'timeout']).isRequired,
  })),
  selected: PropTypes.number,
  onRowClick: PropTypes.func,
};

EventsViewSelector.defaultProps = {
  items: [],
  selected: 0,
  onRowClick: () => null,
};

export default EventsViewSelector;
