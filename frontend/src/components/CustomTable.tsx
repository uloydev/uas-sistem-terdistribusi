import React from 'react'
import DataTable, { createTheme } from 'react-data-table-component'

interface CustomTableProps {
    columns: Array<object>
    data: Array<object>
}

const CustomTable: React.FC<CustomTableProps> = ({ columns, data }) => {
    createTheme('solarize', {
        text: {
            primary: '#fffff',
        },
        background: {
            default: 'transparent'
        },
        context: {
            background: '#FFFFF',
            text: '#FFFFFF',
        },
        divider: {
            default: 'rgb(30,30,47)',
        },
    }, 'dark')
    return (
        <DataTable
            theme='solarize'
            columns={columns}
            data={data}
            pagination
        />
    )
}

export default CustomTable