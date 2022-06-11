import React from 'react'
import Link from 'next/link'
import { IconType } from 'react-icons'

interface SidebarLinkProps {
    href: string
    children: React.ReactNode
    className?: string
    Icon?: IconType
}

const SidebarLink: React.FC<SidebarLinkProps> = ({ Icon, href, children, className }) => {
    return (
        <Link href={href}>
            <li className='cursor-pointer p-5 hover:brightness-110 bg-primary flex items-center gap-x-4 text-slate-300'>
                {Icon && <Icon size={20} />}
                <span className={`${className} `}>{children}</span>
            </li>
        </Link>
    )
}

SidebarLink.defaultProps = {
    className: ""
}

export default SidebarLink