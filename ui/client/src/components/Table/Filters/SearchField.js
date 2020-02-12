import InputAdornment from '@material-ui/core/InputAdornment';
import SearchIcon from '@material-ui/icons/Search';
import TextField from '@material-ui/core/TextField';
import React, { useRef } from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import FormControl from '@material-ui/core/FormControl';

const useStyles = makeStyles((theme) => ({
  root: {
    margin: theme.spacing(1),
  },
}));

const SearchField = ({
  label, defaultValue, onChange, delay,
}) => {
  const classes = useStyles();
  const timer = useRef();
  const handleChange = (event) => {
    const { value } = event.target;
    if (delay) {
      if (timer.current) {
        clearTimeout(timer.current);
      }
      timer.current = setTimeout(() => {
        onChange(value);
        timer.current = null;
      }, delay);
    } else {
      onChange(value);
    }
  };
  return (
    <FormControl className={classes.root}>
      <TextField
        variant="outlined"
        placeholder={label}
        defaultValue={defaultValue}
        onChange={handleChange}
        classes={{ root: classes.textField }}
        InputProps={{
          endAdornment: (
            <InputAdornment position="end">
              <SearchIcon />
            </InputAdornment>
          ),
        }}
      />
    </FormControl>
  );
};
SearchField.propTypes = {
  label: PropTypes.string.isRequired,
  defaultValue: PropTypes.string,
  onChange: PropTypes.func.isRequired,
  delay: PropTypes.number,
};

SearchField.defaultProps = {
  defaultValue: null,
};

export default SearchField;
