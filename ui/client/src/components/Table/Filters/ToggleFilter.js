import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import React from 'react';
import PropTypes from 'prop-types';

const ToggleFilter = ({ label, checked, onChange }) => (
  <FormControlLabel
    control={
      <Switch checked={checked} onChange={onChange} value={label} />
    }
    label={label}
  />
);

ToggleFilter.propTypes = {
  label: PropTypes.string.isRequired,
  checked: PropTypes.bool.isRequired,
  onChange: PropTypes.func.isRequired,
};

export default ToggleFilter;
