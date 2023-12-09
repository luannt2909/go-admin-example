import React from 'react'
import {BooleanInput, Create, SelectInput, SimpleForm, TextInput} from 'react-admin'
import {WebhookTypes} from './webhookType'
import {Box, Grid, Typography} from "@material-ui/core";
import CronScheduleInput from "./CronScheduleInput";

const ReminderCreate = (props) => {
    return (
        <Create title='Create a Reminder' {...props} redirect="list">
            <SimpleForm sx={{maxWidth: 1000}}>
                <Typography variant="h6" gutterBottom>General</Typography>
                <TextInput source='name'
                           required
                           fullWidth
                           helperText="ex: Daily reminder bot"/>

                <BooleanInput source='is_active' label='Active'/>
                <Separator/>

                <Typography variant="h6" gutterBottom>Cron Schedule Specification</Typography>
                <CronScheduleInput/>
                <Separator/>

                <Typography variant="h6" gutterBottom>Webhook information</Typography>
                <Grid container>
                    <Grid item xs={6} md={2}>
                        <Box sx={{width: '10%'}}>
                            <SelectInput choices={WebhookTypes}
                                         required
                                         label="Webhook Type"
                                         source='webhook_type'/>
                        </Box>
                    </Grid>
                    <Grid item xs={12} md={10}>
                        <Box>
                            <TextInput fullWidth required source='webhook' label="Webhook URL"/>
                        </Box>
                    </Grid>
                </Grid>
                <TextInput multiline
                           fullWidth
                           required
                           placeholder="Input your message..."
                           source='message'/>
            </SimpleForm>
        </Create>
    )
}
const Separator = () => <Box pt="1em"/>;

export default ReminderCreate