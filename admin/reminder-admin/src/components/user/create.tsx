import React from 'react'
import {BooleanInput, Create, SelectInput, SimpleForm, TextInput} from 'react-admin'
import {Role} from './role'

const UserCreate = (props) => {
    return (
        <Create title='Create a User' {...props}>
            <SimpleForm>
                <TextInput source='email' required fullWidth/>
                <TextInput source='password' required fullWidth/>
                <SelectInput choices={Role} source='role'/>
                <BooleanInput source='is_active'/>
            </SimpleForm>
        </Create>
    )
}

export default UserCreate