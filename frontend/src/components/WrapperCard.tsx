import React from 'react'

interface WrapperCard {
    children: React.ReactNode
}

const WrapperCard: React.FC<WrapperCard> = ({ children }) => {
    return (
        <div className='bg-dark-light rounded-md p-4'>
            {children}
        </div>
    )
}

export default WrapperCard