import Toolbar from '@material-ui/core/Toolbar';
import Box from '@material-ui/core/Box';
import { Link } from 'react-router-dom';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import AppBar from '@material-ui/core/AppBar';
import React from 'react';
import PropTypes from 'prop-types';
import GitHubIcon from '@material-ui/icons/GitHub';
import logo from '../../../images/logo.png';
import StatusBayVersion from './StatusBayVersion';

const Topbar = ({ onChangeThemeType, isDarkMode, className }) => {
  return (
    <AppBar position="fixed" className={className}>
      <Toolbar>
        <Box width="100%" display="flex" alignItems="center" justifyContent="space-between">
          <Link to="/"><img src={logo} height={59} alt="" /></Link>
          <StatusBayVersion />
          <Box display="flex" alignItems="center" width={180} justifyContent="space-between">
            <FormControlLabel
              control={(
                <Switch
                  checked={isDarkMode}
                  onChange={onChangeThemeType}
                />
              )}
              label="Dark Mode"
            />
            <a href="https://github.com/similarweb/statusbay"><GitHubIcon /></a>
          </Box>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

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
