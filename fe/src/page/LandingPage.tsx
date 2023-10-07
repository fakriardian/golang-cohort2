import React from 'react';
import { Button, Paper, Typography, TextField } from '@mui/material';
import { useLocation, useNavigate } from 'react-router-dom';

const LandingPage: React.FC = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const userData = location.state && location.state.userData;

    const handleLogout = () => {
        navigate('/');
    };
    return (
        <div className="landing-page-container">
            <Paper elevation={3} className="landing-page-box">
                <Typography variant="h5" component="div">
                    Profile Page
                </Typography>
                <hr />
                {userData.map((item: any, index: any) => (
                    Object.entries(item).map(([key, value]) => (
                        key !== 'ID' && (
                            <TextField
                                key={key}
                                label={key}
                                fullWidth
                                margin="normal"
                                variant="outlined"
                                disabled={true}
                                value={value}
                            />
                        )
                    ))
                ))}

                <Button
                    variant="contained"
                    color="warning"
                    fullWidth
                    onClick={handleLogout}
                >
                    Logout
                </Button>
            </Paper>
        </div>
    );
};

export default LandingPage;
