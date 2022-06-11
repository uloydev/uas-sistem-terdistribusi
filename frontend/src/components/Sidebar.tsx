import React, { FC } from 'react'
import { BiTable, BiPencil } from 'react-icons/bi'
import { TbLicense } from 'react-icons/tb'
import { MdOutlineCreateNewFolder } from 'react-icons/md'
import { CgEditMarkup } from 'react-icons/cg'

import SidebarLink from './SidebarLink'

interface SidebarProps {
    open: boolean
}

const Sidebar: FC<SidebarProps> = ({ open }) => {
    return (
        <aside className={`bg-primary absolute lg:relative h-screen w-[285px] pb-10 rounded-r-lg lg:translate-x-0 ${open ? "" : "-translate-x-full"}`}>
            <div className='px-3 mb-3'>
                <h4 className={`after:content-[''] after:h-0.5 after:rounded-full after:w-full after:bg-slate-200 after:absolute after:bottom-0 after:left-0 relative text-center text-slate-200 text-xl font-bold py-5`}>
                    Warement
                </h4>
            </div>
            <ul>
                <SidebarLink Icon={BiTable} href="/">Product Catalog</SidebarLink>
                <SidebarLink Icon={BiPencil} href="/product/create">Product Create</SidebarLink>
                <SidebarLink Icon={BiPencil} href="/product/edit">Product Edit</SidebarLink>
                <SidebarLink Icon={TbLicense} href="/product/edit">List License</SidebarLink>
                <SidebarLink Icon={MdOutlineCreateNewFolder} href="/license/create">Create License</SidebarLink>
                <SidebarLink Icon={CgEditMarkup} href="/product/edit">Edit License</SidebarLink>
            </ul>
        </aside>
    )
}

export default Sidebar