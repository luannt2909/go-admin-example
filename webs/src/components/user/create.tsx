import React from 'react'
import {BooleanInput, Create, SelectInput, SimpleForm, TextInput} from 'react-admin'
import {Roles} from './roles'

const UserCreate = (props) => {
    return (
        <Create title='Create a User' {...props}>
            <SimpleForm>
                <TextInput source='username' required fullWidth/>
                <SelectInput choices={Roles} source='role'/>
                <BooleanInput source='is_active'/>
            </SimpleForm>
        </Create>
    )
}

export default UserCreate