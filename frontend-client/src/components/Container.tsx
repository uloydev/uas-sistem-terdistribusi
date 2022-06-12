import React from 'react'

interface ContainerProps {
    children: React.ReactNode
    className?: string
}

const Container: React.FC<ContainerProps> = ({ children, className }) => {
    return (
        <div className={`${className} px-5`}>
            {children}
        </div>
    )
}

export default Container