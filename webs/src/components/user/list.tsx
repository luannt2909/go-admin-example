import React from 'react'
import {
    BooleanField,
    Datagrid,
    DateField,
    DeleteButton,
    EditButton,
    FunctionField,
    List,
    SelectInput,
    TextField,
    TextInput,
    usePermissions,
    WrapperField
} from 'react-admin'
import {GetRole, RoleAdmin, Roles} from "./roles";

const userFilters = [
    <TextInput label="username" source="username" alwaysOn/>,
    <SelectInput choices={Roles} source='role' alwaysOn/>,
];

const UserList = (props) => {
    const {permissions} = usePermissions()
    const isAdminRole = permissions == RoleAdmin
    return (
        <List {...props} filters={userFilters} hasCreate={isAdminRole}>
            <Datagrid rowClick="show">
                <TextField source='id'/>
                <TextField source='username'/>
                <BooleanField source='is_active'/>
                <FunctionField
                    label="Role"
                    render={record => GetRole(record.role)}
                />;
                <WrapperField label="Created By">
                    <TextField source='created_by'/>
                    <br/>
                    <DateField source='created_at' showTime={true}/>
                </WrapperField>
                <FunctionField
                    render={record => {
                        return (!record.current_user && isAdminRole) &&
                            (<>
                                <EditButton/>
                                <DeleteButton/>
                            </>)
                    }}
                />;
            </Datagrid>
        </List>
    )
}

export default UserList