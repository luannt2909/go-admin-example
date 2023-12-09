import React from 'react'
import {
    List,
    Datagrid,
    TextField,
    DateField,
    EditButton,
    DeleteButton,
    BooleanField
} from 'react-admin'

const UserList = (props) => {
    return (
        <List {...props}>
            <Datagrid>
                <TextField source='id' />
                <TextField source='email' />
                <BooleanField source='is_active' />
                <TextField source='role_text' />
                <EditButton />
                <DeleteButton/>
            </Datagrid>
        </List>
    )
}

export default UserList