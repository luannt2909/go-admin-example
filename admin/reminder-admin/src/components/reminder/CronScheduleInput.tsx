import React, {useEffect, useState} from 'react'
import {TextInput, useRecordContext} from "react-admin";
import {Cron} from 'react-js-cron'
import 'react-js-cron/dist/styles.css'
import {useFormContext} from 'react-hook-form';

const CronScheduleInput = props => {
    const record = useRecordContext();
    const initValue = record ? `${record.schedule}` : '30 5 * * 1,6'
    const [schedule, setSchedule] = useState(initValue)
    const {setValue} = useFormContext();
    useEffect(() => {
        setValue('schedule', schedule,{ shouldDirty: true })
    });
    const onChange = (event) => {
        setSchedule(event.target.value)
    }
    return (
        <>
            <TextInput source='schedule'
                       sx={{width: "50%"}}
                       required
                       onChange={onChange}
                       helperText="ex: '* * * * *', '@every 5m',... "/>
            <Cron value={schedule} setValue={setSchedule}/>
        </>
    )
}
export default CronScheduleInput;