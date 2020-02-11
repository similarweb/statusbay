import React from 'react';
import PropTypes from 'prop-types';
import Stepper from '@material-ui/core/Stepper';
import StepConnector from '@material-ui/core/StepConnector';
import makeStyles from '@material-ui/core/styles/makeStyles';
import TimelineItem from './TimelineItem';

const useStyles = makeStyles(() => ({
  connector: {
    paddingBottom: 0,
  },
  wrapper: {
    maxHeight: (props) => props.maxHeight || 'none',
    overflowY: 'auto',
  },
}));

const Timeline = ({ items, maxHeight }) => {
  const classes = useStyles({ maxHeight });
  return (
    <div className={classes.wrapper}>
      <Stepper activeStep={0} orientation="vertical" connector={<StepConnector classes={{ vertical: classes.connector }} />}>
        {items.map((item) => <TimelineItem key={`timeline-item-${item.time}`} {...item} />)}
      </Stepper>
    </div>
  );
};

Timeline.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape(TimelineItem.propTypes)),
  maxHeight: PropTypes.number,
};

Timeline.defaultProps = {
  items: [],
  maxHeight: null,
};

export default Timeline;
