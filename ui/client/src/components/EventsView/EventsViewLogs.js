import React from 'react';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import TableBody from '@material-ui/core/TableBody';
import Table from '@material-ui/core/Table';
import PropTypes from 'prop-types';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Box from '@material-ui/core/Box';
import Timeline from '../Timeline/Timeline';
import NoData from '../Table/NoData';

const useStyles = makeStyles((theme) => ({
  cell: {
    borderBottom: 0,
    padding: 0,
  },
  tabs: {
    minHeight: 24,
  },
  tab: {
    minHeight: 36,
  },
  tabWrapper: {
    display: 'flex',
    alignItems: 'center',
    flexDirection: 'row'
  },
  tabsCell: {
    paddingTop: 0,
    paddingBottom: 0,
  },
  marker: {
    width: 10,
    height: 10,
    backgroundColor: theme.palette.error.main,
    borderRadius: '50%',
    marginRight: 12,
  },
}));

const EventsViewLogs = ({
  logs, onTabChange, selectedTab, tabs,
}) => {
  const classes = useStyles();
  const tableContent = logs.length > 0 ? <Timeline items={logs} maxHeight={400} />
    : <NoData imageWidth={120} message="No events available" />;
  return (
    <Table size="small">
      <TableHead>
        <TableRow>
          <TableCell classes={{ root: classes.tabsCell }}>
            <Tabs
              value={selectedTab}
              onChange={onTabChange}
              className={classes.tabs}
              variant="scrollable"
              scrollButtons="auto"
            >
              {
                tabs.map(({ name, error }) => (
                  <Tab
                    disableRipple
                    classes={{ root: classes.tab, wrapper: classes.tabWrapper }}
                    key={name}
                    icon={<div className={error && classes.marker} />}
                    label={name}
                  />
                ))
              }
            </Tabs>
          </TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        <TableRow>
          <TableCell classes={{ root: classes.cell }}>
            {
              tableContent
            }
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>
  );
};

EventsViewLogs.propTypes = {
  logs: PropTypes.arrayOf(PropTypes.shape({
    title: PropTypes.string,
    time: PropTypes.number,
    error: PropTypes.bool,
    content: PropTypes.string,
  })),
  onTabChange: PropTypes.func.isRequired,
  selectedTab: PropTypes.number.isRequired,
  tabs: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string,
    error: PropTypes.bool,
  })),
};

EventsViewLogs.defaultProps = {
  logs: [],
  tabs: [],
};

export default EventsViewLogs;
