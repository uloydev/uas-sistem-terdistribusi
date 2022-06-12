import React from 'react'
import BurgerToggle from './BurgerToggle'

interface HeaderProps {
    title: string
    sidebarToggle: () => void
}

const Header: React.FC<HeaderProps> = ({ title, sidebarToggle }) => {
    return (
        <nav className={`flex items-center px-5 py-4`}>
            <BurgerToggle onClick={sidebarToggle} />
            <h3 className='font-bold'>{title}</h3>
        </nav>
    )
}

export default Header