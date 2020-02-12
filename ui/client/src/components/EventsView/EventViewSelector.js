import React, { useMemo } from 'react';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import makeStyles from '@material-ui/core/styles/makeStyles';
import PropTypes from 'prop-types';
import CellStatus from '../Table/Cells/CellStatus';
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
        cell: (row) => row.name,
      },
      {
        name: 'Status',
        header: (name) => <TableCell>{name}</TableCell>,
        cell: (row) => row.status,
      },
    ],
  }), []);
  return (
    <div className={classes.container}>
      <TableStateless data={items} config={tableConfig} tableSize="small" stickyHeader={true}/>
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

export default React.memo(EventsViewSelector);
