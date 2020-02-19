import createMuiTheme from '@material-ui/core/styles/createMuiTheme';
import { green } from '@material-ui/core/colors';

export default (type) => createMuiTheme({
  palette: {
    type,
    secondary: {
      main: green[500],
      contrastText: '#ffffff',
    },
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
    MuiTableCell: {
      head: {
        textTransform: 'uppercase'
      }
    },
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
    MuiPickersModal: {
      dialogRoot: {
        '& .MuiDialogActions-root button:last-of-type': {
          color: '#ffffff',
          backgroundColor: green[500],
          borderColor: green[500],
        },
      },
    },
  },
});
