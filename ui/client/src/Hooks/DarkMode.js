import useLocalStorage from 'react-use/lib/useLocalStorage';
import { useState } from 'react';
import Theme from '../Themes/default';

export default () => {
  const [muiThemeType, setMuiThemeType] = useLocalStorage('mui-theme-type', 'light');
  const [theme, setTheme] = useState(Theme(muiThemeType));

  const toggleDarkMode = () => {
    const type = muiThemeType === 'light' ? 'dark' : 'light';
    setMuiThemeType(type);
    setTheme(Theme(type));
  };

  return [theme, toggleDarkMode];
};
