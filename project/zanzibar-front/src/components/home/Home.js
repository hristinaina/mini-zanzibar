import React from 'react';
import theme from '../../themes/theme';
import Navigation from '../navigation/Navigation';
import { ThemeProvider } from '@mui/material';
import './Home.css';

const Home = () =>  {

    
    return (
        <ThemeProvider theme={theme}>
            <Navigation/>
            <div className='home-background'>
                <p>Welcome to the Home Page!</p>
            </div>
        </ThemeProvider>
    );
}

export default Home;