import axios from 'axios'
import { useRouter } from 'next/router'
import React from 'react'
import { BASE_URL_LICENSE_API } from '../../utils/const'
import { unixToDate } from '../../utils/helpers'
import ActionTableButton from './ActionTableButton'
import CustomTable from './CustomTable'
import Template from './Template'
import WrapperCard from './WrapperCard'

interface LicenseProp {
    id: number
    created_at: number
    updated_at: number
    username: string
    key: string
    is_active: string
}

const LicenseListView = () => {
    const fetchDataLicense = async () => {
        try {
            const { data: result } = await axios.get(BASE_URL_LICENSE_API)
            const { data } = result
            if (data === null) {
                setIsLoading(false)
                setLicense([])
                return
            }
            const filtererData = data.map((item: any) => {
                return {
                    ...item,
                    created_at: unixToDate(item.created_at),
                    updated_at: unixToDate(item.updated_at),
                    action: <ActionTableButton onUpdate={() => router.push(`/license/update/${item.id}`)} onDelete={() => handleDeleteLicense(item.id)} />
                }
            })
            setLicense([...filtererData])
            setIsLoading(false)
        } catch (err) {
            console.log(err)
        }
    }
    const router = useRouter()
    const [license, setLicense] = React.useState<LicenseProp[] | []>([])
    const [isLoading, setIsLoading] = React.useState<boolean>(true)

    const fetcher = React.useRef(fetchDataLicense);

    const handleDeleteLicense = async (id: number) => {
        try {
            const response = await axios.delete(BASE_URL_LICENSE_API + `/${id}`)
            fetchDataLicense()
        } catch (err) {
            console.log(err)
        }
    }

    React.useEffect(() => {
        fetcher.current()
    }, [fetcher])

    const columns = [
        {
            name: 'Id',
            selector: (row: any) => row.id
        },
        {
            name: 'Created at',
            selector: (row: any) => row.created_at
        },
        {
            name: 'Update at',
            selector: (row: any) => row.updated_at
        },
        {
            name: 'Username',
            selector: (row: any) => row.username
        },
        {
            name: 'Key',
            selector: (row: any) => row.key
        },
        {
            name: 'Action',
            selector: (row: any) => row.action
        },
    ]

    return (

        <Template title="List License">
            <WrapperCard>
                {isLoading ? (<div>loading</div>) : (<CustomTable columns={columns} data={license} />)}
            </WrapperCard>
        </Template>
    )
}

export default LicenseListView