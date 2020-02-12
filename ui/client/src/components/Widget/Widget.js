import React from 'react';
import PropTypes from 'prop-types';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardContent from '@material-ui/core/CardContent';
import makeStyles from '@material-ui/core/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  header: {
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
}));

const Widget = ({ title, children }) => {
  const classes = useStyles();
  const t = title;
  return (
    <Card>
      <CardHeader title={t} titleTypographyProps={{ variant: 'h6' }} classes={{ root: classes.header }} />
      <CardContent>{children}</CardContent>
    </Card>
  );
};
Widget.propTypes = {
  title: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
};

export default Widget;
