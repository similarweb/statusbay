import React, { useCallback, useState } from 'react';
import PropTypes from 'prop-types';
import Checkbox from '@material-ui/core/Checkbox';
import ListItemText from '@material-ui/core/ListItemText';
import FormControl from '@material-ui/core/FormControl';
import { makeStyles } from '@material-ui/core/styles';
import Select from '@material-ui/core/Select';
import Input from '@material-ui/core/Input';
import Box from '@material-ui/core/Box';
import MenuItem from '@material-ui/core/MenuItem';
import Chip from '@material-ui/core/Chip';
import InputLabel from '@material-ui/core/InputLabel';

const useStyles = makeStyles((theme) => {
  const light = theme.palette.type === 'light';
  return ({
    root: {
      margin: theme.spacing(1),
    },
    chip: {
      marginLeft: 10,
      height: 16,
    },
    chipLabel: {
      paddingLeft: 12,
      paddingRight: 12,
    },
    menu: {
      minWidth: 200,
      maxHeight: 400,
    },
    placeholder: {
      opacity: light ? 0.42 : 0.5,
      color: theme.palette.text.primary,
    },
    checkbox: {
      marginRight: 4,
      backgroundColor: 'transparent',

    },
    checkboxChecked: {
      '&:hover': {
        backgroundColor: 'transparent',
      },
    },
    menuItem: {
      paddingRight: theme.spacing(3),
      paddingLeft: theme.spacing(1),
    },
  });
});

const MultiSelectValue = ({ selected, name }) => {
  const classes = useStyles();
  if (selected.length > 0) {
    return (
      <Box display="flex" alignItems="center">
        <InputLabel>{`${name}:`}</InputLabel>
        <Chip size="small" label={selected.length} classes={{ root: classes.chip, label: classes.chipLabel }} />
      </Box>
    );
  }
  return <Box><InputLabel classes={{ root: classes.placeholder }}>{name}</InputLabel></Box>;
};

MultiSelectValue.propTypes = {
  name: PropTypes.string.isRequired,
  selected: PropTypes.arrayOf(PropTypes.string),
};

MultiSelectValue.defaultProps = {
  selected: [],
};

const MultiSelect = ({
  name, values, selectedValue, onChange,
}) => {
  const classes = useStyles();

  const [open, setOpen] = useState(false);
  const handleOpen = useCallback(() => {
    setOpen(true);
  }, []);
  const handleClose = useCallback(() => {
    setOpen(false);
  }, []);
  const onChangeInternal = useCallback((event) => {
    setOpen(false);
    onChange(event.target.value);
  }, [onChange]);
  return (
    <FormControl className={classes.root}>
      <Select
        open={open}
        onClose={handleClose}
        onOpen={handleOpen}
        variant="outlined"
        multiple
        displayEmpty
        value={selectedValue}
        onChange={onChangeInternal}
        renderValue={(selected) => <MultiSelectValue selected={selected} name={name} />}
        MenuProps={{
          classes: {
            paper: classes.menu,
          },
        }}
      >
        <MenuItem disabled value="">
          {name}
        </MenuItem>
        {values != undefined && values.map((value) => (
          <MenuItem value={value} key={value} classes={{ gutters: classes.menuItem }}>
            <Box display="flex" alignItems="center">
              <Checkbox checked={selectedValue.indexOf(value) > -1} size="small" classes={{ root: classes.checkbox }} disableRipple />
              <ListItemText primary={value} primaryTypographyProps={{ variant: 'body2' }} />
            </Box>
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};

MultiSelect.propTypes = {
  name: PropTypes.string.isRequired,
  values: PropTypes.arrayOf(PropTypes.string).isRequired,
  selectedValue: PropTypes.arrayOf(PropTypes.string),
  onChange: PropTypes.func.isRequired,
};

MultiSelect.defaultProps = {
  selectedValue: [],
  values: [],
};

export default MultiSelect;
