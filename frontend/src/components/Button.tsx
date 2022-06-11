import React from 'react'

interface ButtonProps {
    className?: string
    onClick?: () => void
    children: React.ReactNode
}

const Button: React.FC<ButtonProps> = ({ className, onClick, children }) => {
    return (
        <button onClick={onClick} className={`${className} w-full`}>
            {children}
        </button>
    )
}

Button.defaultProps = {
    className: "",
}

export default Button