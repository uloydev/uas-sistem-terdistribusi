import React from 'react'

interface BurgerToggleProps {
    onClick: () => void
}

const BurgerToggle: React.FC<BurgerToggleProps> = ({ onClick }) => {
    return (
        <div className='space-y-1 mr-3 cursor-pointer lg:hidden' onClick={onClick}>
            <div className="w-5 h-0.5 bg-white rounded-lg"></div>
            <div className="w-5 h-0.5 bg-white rounded-lg"></div>
            <div className="w-5 h-0.5 bg-white rounded-lg"></div>
        </div>
    )
}

export default BurgerToggle