import React, { useCallback, useState } from 'react';
import Box from '@material-ui/core/Box';
import PropTypes from 'prop-types';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import EventsViewSelector from './EventViewSelector';
import EventsViewLogs from './EventsViewLogs';

const EventsView = ({ items }) => {
  const [selectedItem, setSelectedItem] = useState(0);
  const handleClick = useCallback((row) => () => {
    setSelectedItem(row);
  }, []);
  if (items.length === 0) {
    return null;
  }
  const selectedLogs = items[selectedItem].logs;
  const eventsViewSelectorItems = items.map(({ name, status, logs = [] }) => ({ name, status, hasError: logs.some(({ error }) => error) }));
  return (

    <Grid container>
      <Grid item xs={12}>
        <Paper>
          <Box display="flex">
            <Box flexBasis="50%" flexShrink="0" flexGrow="1">
              <EventsViewSelector
                selected={selectedItem}
                onRowClick={handleClick}
                items={eventsViewSelectorItems}
              />
            </Box>
            <Box flexBasis="50%" flexShrink="0" flexGrow="1">
              <EventsViewLogs logs={selectedLogs} />
            </Box>
          </Box>
        </Paper>
      </Grid>
    </Grid>

  );
};

EventsView.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    logs: PropTypes.arrayOf(PropTypes.shape(EventsViewLogs.propTypes)),
    ...EventsViewSelector.propTypes,
  })),
};

EventsView.defaultProps = {
  items: [],
};
export default EventsView;
