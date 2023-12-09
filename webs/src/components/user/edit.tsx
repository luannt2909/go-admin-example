import React from 'react'
import {BooleanInput, Edit, SelectInput, SimpleForm, TextInput, usePermissions, useRecordContext, WithRecord} from 'react-admin'
import {RoleAdmin, Roles} from "./roles";

const UserEdit = (props) => {
    const {permissions} = usePermissions()
    return (
        <Edit title='Edit User' {...props}>
            <SimpleForm >
                <TextInput disabled source='id'/>
                <TextInput source='username' disabled fullWidth/>
                <WithRecord render={record => {
                    const isDisabled = permissions !== RoleAdmin || record.current_user == true
                    return (
                        <>
                            <BooleanInput source='is_active' disabled={isDisabled} />
                            <SelectInput source='role' choices={Roles} disabled={isDisabled}/>
                        </>
                    )
                }}/>
            </SimpleForm>
        </Edit>
    )
}

export default UserEdit