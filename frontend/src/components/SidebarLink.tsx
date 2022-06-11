import React from 'react'
import Link from 'next/link'

interface SidebarLinkProps {
    href: string
    children: React.ReactNode
    className?: string
    Icon: 
}

const SidebarLink: React.FC<SidebarLinkProps> = ({ Icon, href, children, className }) => {
    return (
        <li className=''>
            <Icon />
            <Link href={href}>
                <span className={`${className} `}>{children}</span>
            </Link>
        </li>
    )
}

SidebarLink.defaultProps = {
    className: ""
}

export default SidebarLink