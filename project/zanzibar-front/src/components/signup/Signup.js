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
import {Switch, FormControlLabel} from '@mui/material';
import {ThemeProvider} from "@emotion/react";
import lightTheme from "../../themes/theme";
import './Signup.css';
import authService from '../../services/AuthService';

export function Signup() {

    const navigate = useNavigate();
    const [confirmPassword, setConfirmPassword] = useState('');
    const [showPassword, setShowPassword] = useState(false);
    const [open, setOpen] = useState(false);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [errors, setErrors] = useState({});

    const [professor, setProfessor] = useState({
        name: '',
        surname: '',
        email: '',
        password: ''
    });


    const validatePassword = (password) => {
        // At least 8 characters long, 1 uppercase letter, 1 lowercase letter, 1 number, 1 special character
        const regex = /^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[*.!@#$%^&(){}[\]:;<>,.?/~_+\-=|\\]).{8,20}$/;
        return regex.test(password);
    };


    const validateForm = () => {
        let tempErrors = {};
        let missingFields = [];
        if (!professor.name) {
            tempErrors.name = "Name is required.";
            missingFields.push('Name');
        }
        if (!professor.surname) {
            tempErrors.surname = "Surname is required.";
            missingFields.push('Surname');
        }
        if (!professor.email) {
            tempErrors.email = "Email is required.";
            missingFields.push('Email');
        }
        if (!professor.password) {
            tempErrors.password = "Password is required.";
            missingFields.push('Password');
        } else if (!validatePassword(professor.password)) {
            tempErrors.password = "Password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character.";
        }
        setErrors(tempErrors);
        return missingFields;
    };

    const handleChange = (e) => {
        setProfessor({...professor, [e.target.name]: e.target.value});
    };

    const handleConfirmPasswordChange = (e) => {
        setConfirmPassword(e.target.value);
    };
    const handleClickShowPassword = () => {
        setShowPassword(!showPassword); // Toggle the visibility of the password
    };


    const handleSubmit = async (e) => {
        e.preventDefault();

        const missingFields = validateForm();
        if (missingFields.length > 0) {
            setSnackbarMessage(`Missing fields: ${missingFields.join(', ')}`);
            return;
        }
        if (professor.password !== confirmPassword) {
            setSnackbarMessage("Passwords do not match!"); // Show the Snackbar
            handleClick();
            return;
        }
        try{
            const result = await authService.register(professor);
            if (result)
                if (result.status === 200) {
                    console.log("Registered");
                    navigate('/login'); // Redirect to login page
                }
        }        
        catch (error) {
            setSnackbarMessage('Please fill all the fields');
            handleClick();
        }

        console.log(professor);

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
            <div id='reg-card'>
                <div className="contain-reg" >
                    <p className="title-reg">Sign up</p>
                    
                    <form id='professor-reg'>
                <div className="fields">
                    <TextField sx={{m: 1, width: '30ch'}} className="fields" name="name" label="Name"
                               value={professor.name}
                               onChange={handleChange}
                               error={Boolean(errors.name)}
                               helperText={errors.name}/>
                </div>
                <div className="fields">
                    <TextField sx={{m: 1, width: '30ch'}} className="fields" name="surname" label="Surname"
                               value={professor.surname}
                               onChange={handleChange}
                               error={Boolean(errors.surname)}
                               helperText={errors.surname}/>
                </div>
                <div className="fields">
                    <TextField sx={{m: 1, width: '30ch'}} className="fields" name="email" label="Email"
                               value={professor.email}
                               onChange={handleChange}
                               error={Boolean(errors.email)}
                               helperText={errors.email}/>
                </div>
                <div className="fields">
                    <TextField
                        sx={{m: 1, width: '30ch'}}
                        className="fields"
                        name="password"
                        label="Password"
                        type={showPassword ? 'text' : 'password'} // Change the type based on the visibility state
                        value={professor.password}
                        onChange={handleChange}
                        error={Boolean(errors.password)}
                        helperText={errors.password}
                        InputProps={{
                            endAdornment: (
                                <InputAdornment position="end">
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={handleClickShowPassword}
                                    >
                                        {showPassword ? <VisibilityOff/> : <Visibility/>}
                                    </IconButton>
                                </InputAdornment>
                            ),
                        }}
                    />
                </div>
                
                <div className="fields">
                    <TextField
                        sx={{m: 1, width: '30ch'}}
                        className="fields"
                        name="confirmPassword"
                        label="Confirm Password"
                        type={showPassword ? 'text' : 'password'} // Change the type based on the visibility state
                        value={confirmPassword}
                        onChange={handleConfirmPasswordChange}
                        InputProps={{
                            endAdornment: (
                                <InputAdornment position="end">
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={handleClickShowPassword}
                                    >
                                        {showPassword ? <VisibilityOff/> : <Visibility/>}
                                    </IconButton>
                                </InputAdornment>
                            ),
                        }}
                    />
                </div> 
                <Button onClick={handleSubmit}
                        id="login"
                        variant="contained"
                        style={{marginTop: "50px", textTransform: 'none', width: '80%', marginLeft: 'auto', marginRight: 'auto', display: 'block'}}
                    
                >Sign up</Button>
                <Snackbar
                        open={open}
                        autoHideDuration={4000}
                        onClose={handleClose}
                        message={snackbarMessage}
                        action={action}
                    />
            </form>
                    

                    <Link to="/login" style={{textDecoration: "none"}}>
                        <p className="reg" variant="contained" style={{textTransform: 'none'}}>Already have an acoount?
                        LOG IN</p>
                    </Link>
                </div>
            </div>
        </ThemeProvider>
    );
}