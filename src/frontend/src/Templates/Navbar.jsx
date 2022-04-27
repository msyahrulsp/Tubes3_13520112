import React from 'react';
import './Navbar.scss';

import { NavLink } from 'react-router-dom';

import Logo from "../images/dna.png"

export const Navbar = () => {
    return (
        <div className="navbar">
            <div className="nav-header">
                <img src={ Logo } alt="logo" />
                <h1 className="nav-logo">Finding (D)N(A)emo</h1>
            </div>
            <div className="nav-links">
                <NavLink 
                    exact
                    to="/"
                    className="nav-button">
                    <h2>Home</h2>
                </NavLink>
                <NavLink 
                    exact
                    to="/add-dna"
                    className="nav-button">
                    <h2>Add DNA</h2>
                </NavLink>
                <NavLink 
                    exact
                    to="/check-dna"
                    className="nav-button">
                    <h2>Check DNA</h2>
                </NavLink>
            </div>
        </div>
    )
}