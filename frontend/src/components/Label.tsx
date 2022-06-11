import React from 'react'

interface LabelProps {
    id: string
    children: React.ReactNode
}

const Label: React.FC<LabelProps> = ({ id, children }) => {
    return (
        <label className='block text-[#98A9AE] mb-2' htmlFor={`${id}`}>{children}</label>
    )
}

export default Label