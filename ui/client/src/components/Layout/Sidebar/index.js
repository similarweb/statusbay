import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import { Link, NavLink } from 'react-router-dom';
import ListItemText from '@material-ui/core/ListItemText';
import Drawer from '@material-ui/core/Drawer';
import React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import { useTranslation } from 'react-i18next';

const drawerWidth = 190;

const useStyles = makeStyles((theme) => ({
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  toolbar: theme.mixins.toolbar,
}));


export default () => {
  const classes = useStyles();
  const { t } = useTranslation();
  return (
    <Drawer
      variant="permanent"
      className={classes.drawer}
      classes={{
        paper: classes.drawerPaper,
      }}
    >
      <div className={classes.toolbar} />
      <List>
        <ListItem button component={NavLink} to="/applications">
          {/* <ListItemIcon></ListItemIcon> */}
          <ListItemText>{t('applications')}</ListItemText>
        </ListItem>
        <ListItem button component={NavLink} to="/">
          {/* <ListItemIcon><AdjustIcon /></ListItemIcon> */}
          <ListItemText>{t('all deployments')}</ListItemText>
        </ListItem>
      </List>
    </Drawer>
  );
};
