import React, {
  useCallback, useEffect, useState,
} from 'react';
import Box from '@material-ui/core/Box';
import PropTypes from 'prop-types';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { useHistory, useLocation } from 'react-router-dom';
import querystring from 'query-string';
import EventsViewSelector from './EventViewSelector';
import EventsViewLogs from './EventsViewLogs';

const hasError = (events) => {
  let errorResult = false;
  events.forEach((event) => {
    if (event.logs.some(({ error }) => error)) {
      errorResult = true;
    }
  });
  return errorResult;
};

const EventsView = ({ items, deploymentId }) => {
  const location = useLocation();
  const history = useHistory();
  const params = querystring.parse(location.search);
  const [selectedItem, setSelectedItem] = useState(parseInt(params.pod || 0));
  const [selectedTab, setSelectedTab] = useState(0);

  useEffect(() => {
    history.replace({
      pathname: location.pathname,
      search: `?${new URLSearchParams({
        ...params,
        pod: selectedItem,
      })}`,
    });
  }, [selectedItem]);

  const handleClick = useCallback((row) => () => {
    setSelectedItem(row);
  }, []);

  const handleTabChange = useCallback((event, newValue) => {
    setSelectedTab(newValue);
  });

  if (items.length === 0) {
    return null;
  }
  const selectedLogs = items[selectedItem].events[selectedTab].logs;
  const tabs = items[selectedItem].events.map(({ name, logs }) => ({
    name,
    error: logs.some(({ error }) => error),
  }));

  const eventsViewSelectorItems = items.map(({ name, status, events = [] }) => ({
    name,
    status,
    hasError: hasError(events),
  }));
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
                deploymentId={deploymentId}
              />
            </Box>
            <Box flexBasis="50%" flexShrink="0" flexGrow="1">
              <EventsViewLogs
                logs={selectedLogs}
                onTabChange={handleTabChange}
                selectedTab={selectedTab}
                tabs={tabs}
              />
            </Box>
          </Box>
        </Paper>
      </Grid>
    </Grid>

  );
};

EventsView.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    events: PropTypes.arrayOf(PropTypes.shape({
      name: PropTypes.string,
      logs: EventsViewLogs.propTypes.logs,
    })),
    name: PropTypes.string,
  })),
  deploymentId: PropTypes.string.isRequired,
};

EventsView.defaultProps = {
  items: [],
};
export default EventsView;
