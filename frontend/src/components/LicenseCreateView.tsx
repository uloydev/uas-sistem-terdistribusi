import React from 'react'
import { useRouter } from 'next/router';
import axios from 'axios';

import { createKey } from '../../utils/helpers';
import buttonStyles from './../../styles/Button.module.css'

import Button from './Button';
import Input from './Input';
import Template from './Template'
import WrapperCard from './WrapperCard';
import Label from './Label';
import { BASE_URL_LICENSE_API } from '../../utils/const';

const LicenseCreateView = () => {
    const router = useRouter()
    const [key, setKey] = React.useState('')
    const [username, setUsername] = React.useState('')

    const handleChangeUserName = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(e.target.value)
    }

    const handleGenerate = () => {
        setKey(createKey())
    }

    const requestSaveLicense = async () => {
        const payload = {
            "is_active": true,
            key: key,
            username: username
        }
        try {
            await axios.post(BASE_URL_LICENSE_API, payload)
            router.push('/license/show')
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <Template title='Create License'>
            <WrapperCard>
                <h4>License Generator</h4>
                <div className='w-96 my-4 flex gap-x-4 items-center'>
                    <Label id="key">Key</Label>
                    <Input id="key" placeholder='click button generate' value={key} disabled />
                    <div className='w-48'>
                        <Button
                            className={`${buttonStyles['btn-lavender']} hover:translate-x-0.5 duration-200 py-2 rounded`}
                            onClick={handleGenerate}>
                            Generate
                        </Button>
                    </div>
                </div>
                <div className='w-96 my-4 flex gap-x-4 items-center'>
                    <Label id='username'>Username</Label>
                    <Input id='username' placeholder='username' value={username} onChange={handleChangeUserName} />
                </div>
                <div className='w-96 mt-10 mb-5'>
                    <Button
                        className={`${buttonStyles['btn-lavender']} hover:-translate-y-0.5 duration-200 py-2 rounded`}
                        onClick={requestSaveLicense}>
                        Save
                    </Button>
                </div>
            </WrapperCard>
        </Template>
    )
}

export default LicenseCreateView