import React from 'react';
import { Button, Dialog, DialogContent, DialogActions, Typography } from '@mui/material';

type EmailNotRegisteredModalProps = {
    isOpen: boolean;
    onRequestClose: () => void;
}

const EmailNotRegisteredModal: React.FC<EmailNotRegisteredModalProps> = ({ isOpen, onRequestClose }) => {
    return (
        <Dialog open={isOpen} onClose={onRequestClose}>
            <DialogContent sx={{ width: '400px', maxWidth: '90vw', maxHeight: '80vh' }}>
                <Typography variant="h5" component="div">
                    Email is not registered!
                </Typography>
            </DialogContent>
            <DialogActions>
                <Button onClick={onRequestClose}>Close</Button>
            </DialogActions>
        </Dialog>
    );
}

export default EmailNotRegisteredModal;
