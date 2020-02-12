import React from 'react';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import Typography from '@material-ui/core/Typography';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import PropTypes from 'prop-types';
import makeStyles from '@material-ui/core/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  root: {
    boxShadow: 'none',
    backgroundColor: 'transparent',
  },
  details: {
    padding: 0,
  },
  summary: {
    flexDirection: 'row-reverse',
    '&:not(.Mui-expanded)': {
      borderBottom: `1px solid ${theme.palette.divider}`,
    },
  },
  summaryContent: {
    '&.Mui-expanded': {
    },
  },
  card: {
    width: '100%',
  },
  expandButton: {
    padding: 0,
    marginRight: theme.spacing(1),
  },
  typography: {
    textTransform: 'uppercase',
  },
}));

const DeploymentDetailsSection = ({
  defaultExpanded, title, children,
}) => {
  const classes = useStyles();
  return (
    <ExpansionPanel className={classes.root} defaultExpanded={defaultExpanded}>
      <ExpansionPanelSummary
        className={classes.summary}
        classes={{ content: classes.summaryContent }}
        expandIcon={<ExpandMoreIcon />}
        IconButtonProps={{
          disableFocusRipple: true,
          disableRipple: true,
          edge: 'start',
          classes: {
            root: classes.expandButton,
          },
        }}
      >
        <Typography classes={{ root: classes.typography }}>{title}</Typography>
      </ExpansionPanelSummary>
      <ExpansionPanelDetails className={classes.details}>
        {children}
      </ExpansionPanelDetails>
    </ExpansionPanel>
  );
};

DeploymentDetailsSection.propTypes = {
  defaultExpanded: PropTypes.bool,
  title: PropTypes.string.isRequired,
  children: PropTypes.node,
};

DeploymentDetailsSection.defaultProps = {
  defaultExpanded: false,
  children: null,
};

export default DeploymentDetailsSection;
