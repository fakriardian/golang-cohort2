import React from 'react';
import './App.css';
import LoginPage from './page/LoginPage';
import LandingPage from './page/LandingPage';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
// import { UserData } from './dto/UserDto';

function App() {
  return (
    <Router>
      <Routes >
        <Route path="/" element={<LoginPage />} />
        <Route path="/landing-page" element={<LandingPage />} />
      </Routes >
    </Router>
  );
}

export default App;
