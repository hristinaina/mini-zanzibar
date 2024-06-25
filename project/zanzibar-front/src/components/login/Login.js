import React, {useEffect, useState} from 'react';
import {Link, useNavigate} from 'react-router-dom';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import InputAdornment from '@mui/material/InputAdornment';
import TextField from '@mui/material/TextField';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import Snackbar from '@mui/material/Snackbar';
import CloseIcon from '@mui/icons-material/Close';
import './Login.css';
import authService from '../../services/AuthService';
import lightTheme from '../../themes/theme';
import {ThemeProvider} from '@emotion/react';


const Login = () => {
    useEffect(() => {
        const fetchData = async () => {
            try {
                const result = authService.validateUser();

                if (result) {
                    navigate('/home');
                } 
            } catch (error) {
                console.error('Error:', error);
            }
        };

        fetchData();
    }, []);

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const navigate = useNavigate();
    const [open, setOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');

    const handleLogin = async () => {
        try{
            if (username.trim() == "" || password.trim() == "") {
                setSnackbarMessage("Please fill all the fields!");
                handleClick();
                return;
            }
            const result = await authService.loginUser(username, password);
            console.log(result);
            if (result.status === 200) {
                const result = authService.validateUser();
                console.log(result);
                //todo check what result really is
                navigate('/home');
            } else {
                setSnackbarMessage("Invalid input!");
                handleClick();
            }
        }
        catch (error) {
            setSnackbarMessage("Invalid email or password");
            handleClick();
          }
    };

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    //snackbar
    const handleClick = () => {
        setOpen(true);
    };

    const handleClose = (event, reason) => {
        if (reason === 'clickaway') {
            return;
        }
        setOpen(false);
    };

    const action = (
        <React.Fragment>
            <IconButton size="small" aria-label="close" color="inherit" onClick={handleClose}>
                <CloseIcon fontSize="small"/>
            </IconButton>
        </React.Fragment>
    );

    return (
        <ThemeProvider theme={lightTheme}>
            <div className="background">
                <div className="left-side" >
                    <p className="title-login">Login</p>
                    <form>
                        <div className="fields">
                            <div className="label">Username:</div>
                            <TextField
                                value={username}
                                onChange={handleUsernameChange}
                                id="username"
                                sx={{m: 1, width: '30ch'}}
                                placeholder="someone@example.com"
                                type="email"
                            />
                        </div>
                        <div className="fields">
                            <div className="label">Password:</div>
                            <TextField
                                id="password"
                                type={showPassword ? 'text' : 'password'}
                                sx={{m: 1, width: '30ch'}}
                                value={password}
                                onChange={handlePasswordChange}
                                required
                                InputProps={{
                                    endAdornment: (
                                        <InputAdornment position="end">
                                            <IconButton
                                                aria-label="toggle password visibility"
                                                onClick={handleClickShowPassword}
                                                onMouseDown={handleMouseDownPassword}
                                            >
                                                {showPassword ? <VisibilityOff/> : <Visibility/>}
                                            </IconButton>
                                        </InputAdornment>
                                    ),
                                }}
                            />
                        </div>
                        <Button
                            id="login"
                            variant="contained"
                            onClick={handleLogin}
                            style={{marginTop: "50px", textTransform: 'none'}}
                            sx={{m: 1, width: '39ch'}}
                        >
                            Login
                        </Button>
                        <Link to='/signup' style={{textDecoration: "none"}}>
                            <p className="reg" variant="contained" style={{textTransform: 'none'}}>No account yet? SIGN
                                UP</p>
                        </Link>
                        <Snackbar
                            open={open}
                            autoHideDuration={5000}
                            onClose={handleClose}
                            message={snackbarMessage}
                            action={action}
                        />
                    </form>
                </div>
                <div className="right-side">
                    <img src="/sunflower.png" className="edu-image" alt="logo image"/>
                    <p className='title'>Welcome to <span style={{color: "black"}}> SunDrive! </span>
                    </p>
                    <p className='text'>Organize all your documents in one place.</p>
                </div>
            </div>
        </ThemeProvider>
    );
};

export default Login;