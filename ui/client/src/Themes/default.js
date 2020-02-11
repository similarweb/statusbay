import createMuiTheme from '@material-ui/core/styles/createMuiTheme';

export default (type) => createMuiTheme({
  palette: {
    type,
  },
  typography: {
    fontWeightRegular: 300,
    fontWeightMedium: 400,
    fontWeightBold: 500,
  },
  props: {
    MuiButton: {
      variant: 'outlined',
    },
    MuiTypography: {
      variantMapping: {
        h1: 'h3',
      },
    },
    MuiTable: {
      component: 'div',
    },
    MuiTableHead: {
      component: 'div',
    },
    MuiTableBody: {
      component: 'div',
    },
    MuiTableRow: {
      component: 'div',
    },
    MuiTableCell: {
      component: 'div',
    },
    MuiTablePagination: {
      component: 'div',
    },
  },
  overrides: {
    MuiTableRow: {
      hover: {
        cursor: 'pointer',
      },
    },
    MuiOutlinedInput: {
      input: {
        padding: '7.5px 10px',
      },
    },
  },
});
