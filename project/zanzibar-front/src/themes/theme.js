import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2', // Customize your primary color here
    },
    secondary: {
      main: '#dc004e', // Customize your secondary color here
    },
    background: {
        default: '#3a0820',
        paper: '#3a0820',
    },
    text: {
        primary: '#000000',
        secondary:'#000000',
    },
  },
  typography: {
    fontFamily: 'Roboto, sans-serif',
  },
});

export default theme;
