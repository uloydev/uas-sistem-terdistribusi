import React from 'react'

interface ButtonProps {
    disabled?: boolean
    className?: string
    onClick?: () => void
    children: React.ReactNode
}

const Button: React.FC<ButtonProps> = ({ className, onClick, children, disabled }) => {
    return (
        <button disabled={disabled} onClick={onClick} className={`${className} w-full disabled:opacity-70 disabled:cursor-not-allowed`}>
            {children}
        </button>
    )
}

Button.defaultProps = {
    className: "",
    disabled: false,
}

export default Button