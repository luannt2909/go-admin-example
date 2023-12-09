import React from 'react'
import {
    BooleanField,
    CreateButton,
    DatagridConfigurable,
    DeleteButton,
    EditButton,
    List,
    TextField,
    TopToolbar, usePermissions,
    WrapperField,
    DateField
} from 'react-admin'
import { Stack } from '@mui/material';
import TestWebhookButton from "./TestWebhookButton";
import {RoleAdmin} from "../user/role";

// Usage
const ListActions = ({props}) => (
    <TopToolbar>
        <TestWebhookButton label="Webhook Test" {...props} />
        <CreateButton/>
    </TopToolbar>
);
const ReminderList = (props) => {
    const { permissions } = usePermissions();
    const isAdminRole = permissions == RoleAdmin;
    return (
        <List {...props}
              actions={<ListActions {...props}/>}
        >
            <DatagridConfigurable>
                {isAdminRole && <TextField source='id'/>}
                {/*<TextField source='id'/>*/}
                <TextField source='name'/>
                <BooleanField source='is_active' label='Active'/>
                <WrapperField label="Schedule" >
                    <Stack>
                        <TextField source="schedule" sx={{ fontWeight: 'bold'}}/>
                        <TextField source="schedule_human" />
                    </Stack>
                </WrapperField>
                <TextField source='next_time'/>
                <TextField source='webhook_type'/>
                {isAdminRole || <TextField source='webhook' label='Webhook URL'/>}

                {isAdminRole &&
                    <WrapperField label="Created By" sortBy={"created_by"}>
                    <Stack>
                        <TextField source="created_by" />
                        <DateField source="updated_at" showTime={true} transform={value => new Date(value*1000)} />
                    </Stack>
                </WrapperField>}
                <>
                    <EditButton/>
                    <TestWebhookButton label="Test" {...props}/>
                    <DeleteButton/>
                </>
            </DatagridConfigurable>
        </List>
    )
}

export default ReminderList