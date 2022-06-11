import React from 'react'

interface TextArea {
    id: string
    value: string
    placeholder?: string
    onChange?: (e: React.ChangeEvent<HTMLTextAreaElement>) => void
}

const TextArea: React.FC<TextArea> = ({ id, value, placeholder, onChange }) => {
    return (
        <textarea
            id={id}
            className="bg-transparent w-full min-h-[100px] placeholder:text-slate-500 text-sm outline-none transition-colors duration-300 focus:border-[#D94FD2] border-gray-700 px-6 py-2 rounded-md border-[1.5px]"
            placeholder={placeholder}
            value={value}
            onChange={onChange} />
    )
}

TextArea.defaultProps = {
    placeholder: "Placeholder text area"
}

export default TextArea