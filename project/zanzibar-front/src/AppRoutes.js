import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/login/Login';
import HomePage from './components/home/Home';

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" />} />
      <Route path="/login" element={<Login/>} />
      <Route path="/home" element={<HomePage />} />
      {/* Add more routes if needed */}
    </Routes>
  );
};  

export default AppRoutes;