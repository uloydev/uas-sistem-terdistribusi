import React from 'react'

interface InputProps {
    type?: string
    id?: string
    value?: string | number
    placeholder?: string
    onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void
}

const Input: React.FC<InputProps> = ({ type, onChange, id, placeholder, value }) => {

    return (
        <input
            id={id}
            placeholder={placeholder}
            className="bg-transparent w-full placeholder:text-slate-500 text-sm outline-none transition-colors duration-300 focus:border-[#D94FD2] border-gray-700 px-6 py-1 rounded-md border-[1.5px]"
            type={type}
            value={value}
            onChange={onChange}
        />
    )
}
Input.defaultProps = {
    type: "text"
}

export default Input