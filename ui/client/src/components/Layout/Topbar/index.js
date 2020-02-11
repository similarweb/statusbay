import Toolbar from '@material-ui/core/Toolbar';
import Box from '@material-ui/core/Box';
import { Link } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import AppBar from '@material-ui/core/AppBar';
import React from 'react';
import PropTypes from 'prop-types';
import GitHubIcon from '@material-ui/icons/GitHub';

const Topbar = ({ onChangeThemeType, isDarkMode, className }) => (
  <AppBar position="fixed" className={className}>
    <Toolbar>
      <Box width="100%" display="flex" alignItems="center" justifyContent="space-between">
        <Link to="/"><Typography variant="h6">StatusBay</Typography></Link>
        <Box display="flex" alignItems="center">
          <FormControlLabel
            control={(
              <Switch
                checked={isDarkMode}
                onChange={onChangeThemeType}
              />
                          )}
            label=""
          />
          <GitHubIcon />
        </Box>
      </Box>
    </Toolbar>
  </AppBar>
);

Topbar.propTypes = {
  onChangeThemeType: PropTypes.func,
  isDarkMode: PropTypes.bool,
  className: PropTypes.string,
};

Topbar.defaultProps = {
  onChangeThemeType: () => null,
  isDarkMode: false,
  className: '',
};
export default Topbar;
