import React from 'react';
import { Button, Dialog, DialogContent, DialogActions, DialogTitle } from '@mui/material';

import { UserData } from '../dto/UserDto';

type UserListModalProps = {
    isOpen: boolean;
    onRequestClose: () => void;
    userList: UserData[];
}

const UserListModal: React.FC<UserListModalProps> = ({ isOpen, onRequestClose, userList }) => {
    return (
        <Dialog open={isOpen} onClose={onRequestClose}>
            <DialogTitle>Email List</DialogTitle>
            <DialogContent sx={{ width: '400px', maxWidth: '90vw', maxHeight: '80vh' }}>
                <ul>
                    {userList.map((user) => (
                        <li key={user.ID}>{user.email}</li>
                    ))}
                </ul>
            </DialogContent>
            <DialogActions>
                <Button onClick={onRequestClose}>Close</Button>
            </DialogActions>
        </Dialog>
    );
}

export default UserListModal;
