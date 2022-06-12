import React from 'react'

interface WrapperCard {
    className?: string
    children: React.ReactNode
}

const WrapperCard: React.FC<WrapperCard> = ({ children, className }) => {
    return (
        <div className={`${className} bg-dark-light rounded-md p-4`}>
            {children}
        </div>
    )
}

WrapperCard.defaultProps = {
    className: ''
}

export default WrapperCard