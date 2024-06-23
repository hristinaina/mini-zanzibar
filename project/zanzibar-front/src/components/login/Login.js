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
                const result = await authService.validateUser();
                console.log(result);

                if (result.name === "student") {
                    console.log("student");
                    navigate('/programs');
                } else if (result.name === "professor") {
                    console.log("professor");
                    navigate('/professor');
                } else {
                    console.log("login");
                    navigate("/login");
                }
            } catch (error) {
                console.error('Error:', error);
                handleClick();
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
            const result = await authService.loginUser(username, password);
            console.log(result);
            if (result) {
                if (result.status === 200) {
                    const result = await authService.validateUser();
                    console.log(result);
                    if (result === 'student') {
                        navigate('/programs');
                    } else if (result === 'professor') {
                        navigate('/prof-mentorship');
                    } else {
                        navigate('/login');
                    }
                } else {
                    setSnackbarMessage(result.error);
                    handleClick();
                }
            }
            else {
                setSnackbarMessage("Invalid email or password");
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
                <img src="/logo-white.png" className="top-right-logo" alt="Logo"/>
                <div className="left-side" >
                    <p className="title-login">Login</p>
                    <form>
                        <div className="fields">
                            <div className="label">Email:</div>
                            <TextField
                                value={username}
                                onChange={handleUsernameChange}
                                id="username"
                                sx={{m: 1, width: '30ch'}}
                                placeholder="someone@example.com"
                                helperText="Required"
                                type="email"
                            />
                        </div>
                        <div className="fields">
                            <div className="label">Password:</div>
                            <TextField
                                id="password"
                                type={showPassword ? 'text' : 'password'}
                                sx={{m: 1, width: '30ch'}}
                                helperText="Required. Min 8 characters, special character, capital letter"
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
                        <Link to="/reg" style={{textDecoration: "none"}}>
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
                    <img src="/edu.gif" className="edu-image" alt="education image"/>
                    <p className='title'>Welcome to <span style={{color: "var(--background-blue)"}}> RefLetter! </span>
                    </p>
                    <p className='text'>Your Compass to Higher Education</p>
                </div>
            </div>
        </ThemeProvider>
    );
};

export default Login;