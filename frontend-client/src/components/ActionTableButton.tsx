import React from 'react'
import Button from './Button'

interface ActionTableButtonProps {
    disabled?: boolean
    onUpdate?: () => void
    onDelete?: () => void
}

const ActionTableButton: React.FC<ActionTableButtonProps> = ({ onUpdate, onDelete, disabled }) => {
    return (
        <div className='w-36  flex gap-4'>
            <Button className='flex-1 bg-orange-400 p-2 rounded' disabled={disabled} onClick={onUpdate}>Update</Button>
            <Button className='flex-1 bg-red-500 p-2 rounded' disabled={disabled} onClick={onDelete}>Deleted</Button>
        </div>
    )
}

export default ActionTableButton