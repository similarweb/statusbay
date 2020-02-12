import { Link, useParams } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';
import React from 'react';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles((theme) => ({
  link: {
    ...theme.typography.body2,
  },
  selected: {
    fontWeight: theme.typography.fontWeightMedium,
  },
}));

export default () => {
  const { appId, deploymentId } = useParams();
  const classes = useStyles();
  return (
    <Breadcrumbs aria-label="breadcrumb" classes={classes.typography}>
      {
            appId && (
            <Link to="/" className={classes.link}>
                    Applications
            </Link>
            )
        }

      {
            appId && (
            <Link to={`/application/${appId}`} className={classes.link}>
              {appId}
            </Link>
            )
        }

      {
            deploymentId && <Typography classes={{ root: classes.selected }} variant="body2">{deploymentId}</Typography>
        }

    </Breadcrumbs>
  );
};
