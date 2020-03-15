import React from 'react';
import PropTypes from 'prop-types';
import StepLabel from '@material-ui/core/StepLabel';
import StepContent from '@material-ui/core/StepContent';
import Step from '@material-ui/core/Step';
import Typography from '@material-ui/core/Typography';
import StepIcon from '@material-ui/core/StepIcon';
import makeStyles from '@material-ui/core/styles/makeStyles';
import Box from '@material-ui/core/Box';
import moment from 'moment';
import WarningIcon from '@material-ui/icons/Warning';
import FiberManualRecord from '@material-ui/icons/FiberManualRecord';
import TimelineErrorBox from './TimelineErrorBox';

const useStyles = makeStyles((theme) => ({
  root: {
    '& svg': {
      transform: 'scale(0.8)',
    },
  },
  label: {
    '&, &.MuiStepLabel-active': {
      ...theme.typography.body2,
    },
  },
  error: {
    color: theme.palette.error.main,
  },
  completed: {
    color: theme.palette.primary.main,
  },
  date: {
    marginRight: 20,
  },
}));
const TimelineItem = ({
  error, content, title, time,
}) => {
  const classes = useStyles();
  return (
    <Step active>
      <StepLabel
        error={error}
        classes={{
          root: classes.root, label: classes.label, error: classes.error, completed: classes.completed,
        }}
        StepIconComponent={StepIcon}
        StepIconProps={{
          icon: error ? <WarningIcon /> : <FiberManualRecord className={classes.completed} />,
        }}
      >
        <Box display="flex" justifyContent="space-between">
          <Typography variant="body2">{title}</Typography>
          <Box display="flex" alignItems="center">
            <Typography classes={{ root: classes.date }} variant="body2">{moment(time / 1000000).format('DD/MM/YYYY')}</Typography>
            <Typography variant="body2">{moment(time / 1000000).format('HH:MM:SS')}</Typography>
          </Box>
        </Box>
      </StepLabel>
      {
            content && (
            <StepContent orientation="vertical" icon="">
              <TimelineErrorBox>{content}</TimelineErrorBox>
            </StepContent>
            )
        }
    </Step>
  );
};

TimelineItem.propTypes = {
  error: PropTypes.bool,
  content: PropTypes.oneOfType([PropTypes.bool, PropTypes.string]),
  title: PropTypes.string.isRequired,
  time: PropTypes.number.isRequired,
};

TimelineItem.defaultProps = {
  error: false,
  content: null,
};

export default TimelineItem;
