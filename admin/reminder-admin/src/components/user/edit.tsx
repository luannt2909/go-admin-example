import React from 'react'
import {BooleanInput, Edit, SelectInput, SimpleForm, TextInput, useInput} from 'react-admin'
import {WebhookTypes} from "./role";

const UserEdit = (props) => {
    return (
        <Edit title='Edit User' {...props}>
            <SimpleForm>
                <TextInput disabled source='id'/>
                <TextInput source='email' required fullWidth/>
                <TextInput source='password' required fullWidth/>
                <BooleanInput source='is_active'/>
            </SimpleForm>
        </Edit>
    )
}

export default UserEdit