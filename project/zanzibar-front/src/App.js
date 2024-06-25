import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import theme from './themes/theme';
import AppRoutes from './AppRoutes';
import { ThemeProvider } from '@mui/material';

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Router>
        <AppRoutes />
      </Router>
  </ThemeProvider>
  );
}

export default App;
