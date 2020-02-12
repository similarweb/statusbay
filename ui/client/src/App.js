import React, {
  useContext, useEffect, useState,
} from 'react';
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import {
  BrowserRouter as Router,
  Switch as RouterSwitch,
  Route,
} from 'react-router-dom';
import makeStyles from '@material-ui/core/styles/makeStyles';
import LinearProgress from '@material-ui/core/LinearProgress';
import Box from '@material-ui/core/Box';
import Topbar from './components/Layout/Topbar';
import Sidebar from './components/Layout/Sidebar';
import Deployments from './containers/Deployments';
import Applications from './containers/Applications';
import DeploymentDetails from './containers/DeploymentDetails';
import ApplicationDeployments from './containers/ApplicationDeployments';
import { getMetaData } from './Services/API/TableApi';
import { AppSettingsContext } from './context/AppSettingsContext';
import useDarkMode from './Hooks/DarkMode';

const drawerWidth = 190;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    '@global': {
      a: {
        textDecoration: 'none',
        color: 'inherit',
      },
      'a:visited': {
        color: 'inherit',
      },
    },
  },
  main: {
    flexGrow: 1,
    minHeight: '100vh',
    display: 'flex',
    flexDirection: 'column',
  },
  gridItem: {
    padding: theme.spacing(2),
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
  toolbar: theme.mixins.toolbar,
}));

function App() {
  const [theme, toggleDarkMode] = useDarkMode();
  const { dispatch } = useContext(AppSettingsContext);
  const [loadSettings, setLoadSettings] = useState(true);
  useEffect(() => {
    const getAppSettings = async () => {
      const { clusters, namespaces, statuses } = await getMetaData();
      dispatch({ type: 'SET_TABLE_FILTERS', filters: { clusters, namespaces, statuses } });
      setLoadSettings(false);
    };
    getAppSettings();
  }, []);
  const onChangeThemeType = (event) => {
    toggleDarkMode(event.target.checked);
  };
  const classes = useStyles();
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <div
          className={classes.root}
          style={{ backgrounColor: theme.palette.background.default }}
        >
          <Topbar
            className={classes.appBar}
            isDarkMode={theme.palette.type === 'dark'}
            onChangeThemeType={onChangeThemeType}
          />
          <Sidebar />
          <main className={classes.main}>
            <div className={classes.toolbar} />
            {
              loadSettings ? <Box m={2} flexGrow={1} justifyContent="space-around" display="flex" flexDirection="column"><LinearProgress /></Box> : (
                <RouterSwitch>
                  <Route exact path="/">
                    <Deployments />
                  </Route>
                  <Route exact path="/applications">
                    <Applications />
                  </Route>
                  <Route exact path="/application-deployments/:appId">
                    <ApplicationDeployments />
                  </Route>
                  <Route path="/application/:deploymentId">
                    <DeploymentDetails />
                  </Route>
                </RouterSwitch>
              )
            }
          </main>
        </div>
      </Router>
    </ThemeProvider>
  );
}

export default App;
