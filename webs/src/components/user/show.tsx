import React from 'react'
import {
    BooleanField,
    DateField,
    DeleteButton,
    EditButton,
    Show,
    SimpleShowLayout,
    TextField,
    TopToolbar,
    usePermissions,
    useRecordContext
} from 'react-admin'
import {RoleAdmin} from "./roles";

const ShowActions = () => {
    const {permissions} = usePermissions()
    const record = useRecordContext()
    const canAction = permissions == RoleAdmin && record && !record.current_user
    return canAction && (
        <TopToolbar>
            <EditButton/>
            <DeleteButton/>
        </TopToolbar>
    );
}
const UserShow = (props) => {
    return (
        <Show {...props} actions={<ShowActions/>}>
            <SimpleShowLayout>
                <TextField source='id'/>
                <TextField source='username'/>
                <BooleanField source='is_active'/>
                <TextField source="created_by"/>
                <DateField source="created_at" showTime={true}/>
            </SimpleShowLayout>
        </Show>
    )
}

export default UserShow