import React, { useState } from "react";
import { Navbar, NavItem, NavLink } from "reactstrap";
import { Link, useLocation } from "react-router-dom";
import "./Navigation.css";
import authService from "../../services/AuthService";

const Navigation = () => {
    const location = useLocation();

    const handleLogout = () => {
        authService.logout();
    };

    return (
        <header>
            <Navbar className="navbar">
                    <ul>
                        <img src="/sunflower.png" alt="sunflower" width="50" height="50" style={{marginRight: "20px"}}></img>
                        <span className="logo">SunDrive</span>
                        <NavItem>
                            <NavLink tag={Link} className="text-light" to="/real-estates">Home</NavLink>
                        </NavItem>
                        <NavItem>
                            <NavLink tag={Link} className="text-light" to="/consumption">Namespaces</NavLink>
                        </NavItem>
                        <NavItem className="logout">
                            <NavLink tag={Link} className="text-light" to="/" onClick={handleLogout}>Log out</NavLink>
                        </NavItem>
                    </ul>
            </Navbar>
        </header>
    );

};

export default Navigation;