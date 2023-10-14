import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { TextField, Button, Paper, Typography, Link } from '@mui/material';
import { UserData } from '../dto/UserDto';
import UserListModal from '../modal/UserListModal';
import EmailNotRegisteredModal from '../modal/EmailNotRegisteredModal';

const LoginPage: React.FC = () => {
  const navigate = useNavigate();
  const [email, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [userList, setUserList] = useState<UserData[]>([]);
  const [showUserList, setShowUserList] = useState(false);
  const [showModalNotRegistered, setShowModalNotRegistered] = useState(false);

  const handleLogin = async () => {
    if (!email || !password) {
      alert("Email and password are required.");
      return;
    }
    try {
      const response = await axios.post(`${process.env.REACT_APP_BE_URL}/user/validate`, { email, password })
      if (response.status === 200 && response.data.data) {
        navigate('/landing-page', { state: { userData: response.data.data } })
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 500) {
          setShowModalNotRegistered(true)
        }
      }
      console.error('Error fetching validate user:', error);
    }
  };

  const handleShowUserList = async () => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_BE_URL}/users`);
      setUserList(response.data.data);
      setShowUserList(true);
    } catch (error) {
      console.error('Error fetching user data:', error);
    }
  };

  const handleCloseUserList = () => {
    setShowUserList(false);
  };

  const handlerCloseNotRegistered = () => {
    setShowModalNotRegistered(false);
  }

  const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  return (
    <div className="login-container">
      <Paper elevation={3} className="login-box">
        <Typography variant="h5" component="div">
          Login
        </Typography>
        <TextField
          label="Email Address"
          fullWidth
          margin="normal"
          variant="outlined"
          value={email}
          onChange={handleUsernameChange}
        />
        <TextField
          label="Password"
          fullWidth
          margin="normal"
          variant="outlined"
          type="password"
          value={password}
          onChange={handlePasswordChange}
        />
        <Button
          variant="contained"
          color="primary"
          fullWidth
          onClick={handleLogin}
        >
          Login
        </Button>
        <Link
          component="button"
          variant="body2"
          onClick={handleShowUserList}
        >
          List of Available Emails
        </Link>
        <UserListModal
          isOpen={showUserList}
          onRequestClose={handleCloseUserList}
          userList={userList}
        />
        <EmailNotRegisteredModal
          isOpen={showModalNotRegistered}
          onRequestClose={handlerCloseNotRegistered}
        />
      </Paper>
    </div>
  );
};

export default LoginPage;
