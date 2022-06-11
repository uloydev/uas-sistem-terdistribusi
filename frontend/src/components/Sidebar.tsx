import React, { FC } from 'react'
import SidebarLink from './SidebarLink'
import { BiTable, BiPencil } from 'react-icons/bi'

const Sidebar: FC = () => {

    return (
        <aside className={`bg-primary  h-screen w-[285px] pb-10 rounded-r-lg`}>
            <div className='px-3 mb-3'>
                <h4 className={`after:content-[''] after:h-0.5 after:rounded-full after:w-full after:bg-slate-200 after:absolute after:bottom-0 after:left-0 relative text-center text-slate-200 text-xl font-bold py-5`}>
                    Warement
                </h4>
            </div>
            <ul>
                <SidebarLink Icon={BiTable} href="/">Product Catalog</SidebarLink>
                <SidebarLink Icon={BiPencil} href="/">Product Edit</SidebarLink>
            </ul>
        </aside>
    )
}

export default Sidebar