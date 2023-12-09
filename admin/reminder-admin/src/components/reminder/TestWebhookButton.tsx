import React, {useState} from 'react'
import IconContentSend from "@material-ui/icons/Send";
import IconCancel from '@material-ui/icons/Cancel';
import {Dialog, DialogContent, Toolbar} from "@material-ui/core";
import DialogTitle from "@material-ui/core/DialogTitle";
import {Button, fetchUtils, SaveButton, SelectInput, SimpleForm, TextInput} from "react-admin";
import {useNotify} from "ra-core";
import {WebhookTypes} from "./webhookType";
import {retrieveToken} from "../../authProvider";

const apiURL = import.meta.env.VITE_SIMPLE_REST_URL

const WebhookTestButton = ({...props}) => {
    const {label} = props
    const [showDialog, setShowDialog] = useState(false);
    const notify = useNotify();
    const handleClick = () => {
        setShowDialog(true);
    };

    const handleCloseClick = () => {
        setShowDialog(false);
    };

    const handleSubmit = async values => {
        const token = retrieveToken()
        let options = {
            method: 'POST',
            body: JSON.stringify(values),
            user: {
                authenticated: true,
                token: `Bearer ${token}`
            }}
        fetchUtils.fetchJson(`${apiURL}/webhook/send`, options)
            .then(() => {
                notify('Send message successful');
            })
            .catch((e) => {
                notify(`Error: Send message failed: ${e.message}`, {type: 'error'})
            })
            .finally(() => {
                // setLoading(false);
            });
    };

    return (
        <>
            <Button onClick={handleClick} label={label}>
                <IconContentSend/>
            </Button>
            <Dialog
                fullWidth
                open={showDialog}
                onClose={handleCloseClick}
                aria-label="Test Webhook"
            >
                <DialogTitle>Send Message To Webhook</DialogTitle>
                <DialogContent>
                    <SimpleForm
                        resource="webhook"
                        // We override the redux-form onSubmit prop to handle the submission ourselves
                        onSubmit={handleSubmit}
                        // We want no toolbar at all as we have our modal actions
                        toolbar={<TestWebhookButtonToolbar onCancel={handleCloseClick}/>}
                    >
                        <SelectInput choices={WebhookTypes} required source='webhook_type' label='Webhook Type'/>
                        <TextInput type="url" multiline
                                   fullWidth required
                                   source='webhook' label='Webhook URL'/>
                        <TextInput multiline fullWidth required source='message'/>
                    </SimpleForm>

                </DialogContent>
            </Dialog>
        </>
    );
}

function TestWebhookButtonToolbar({onCancel, ...props}) {
    return (
        <Toolbar {...props}>
            <SaveButton icon={<IconContentSend/>} alwaysEnable label="Send" submitOnEnter={true}/>
            <CloseButton size='large' color='secondary' sx={{ml: 1}} onClick={onCancel}/>
        </Toolbar>
    );
}

function CloseButton(props) {
    return (
        <Button label="Close" {...props}>
            <IconCancel/>
        </Button>
    );
}

export default WebhookTestButton