import { MuiPickersUtilsProvider, DateTimePicker } from '@material-ui/pickers';
import React from 'react';
import PropTypes from 'prop-types';
import MomentUtils from '@date-io/moment';
import { makeStyles } from '@material-ui/core/styles';
import FormControl from '@material-ui/core/FormControl';

const useStyles = makeStyles((theme) => ({
  root: {
    margin: theme.spacing(1),
  },
}));

const DatePickerFilter = ({
  format, label, value, onChange,
}) => {
  const classes = useStyles();
  return (
    <FormControl className={classes.root}>
      <MuiPickersUtilsProvider utils={MomentUtils}>
        <DateTimePicker
          inputVariant="outlined"
          format={format}
          value={value}
          onChange={onChange}
          emptyLabel={label}
        />
      </MuiPickersUtilsProvider>
    </FormControl>
  );
};

DatePickerFilter.propTypes = {
  format: PropTypes.string,
  label: PropTypes.string.isRequired,
  value: PropTypes.number,
  onChange: PropTypes.func.isRequired,
};

DatePickerFilter.defaultProps = {
  format: 'DD/MM/YYYY HH:MM:SS',
  value: null,
};

export default DatePickerFilter;
