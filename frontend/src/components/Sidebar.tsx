import React, { FC } from 'react'
import SidebarLink from './SidebarLink'
import styles from "/styles/Sidebar.module.css"

const Sidebar: FC = () => {

    return (
        <aside className={`${styles.sidebar} bg-primary`}>
            <h4 className={`${styles.header} font-semibold`}>Warement</h4>
            <ul>
                <SidebarLink href="/">Product Catalog</SidebarLink>
                <SidebarLink href="/">Product Edit</SidebarLink>
            </ul>
        </aside>
    )
}

export default Sidebar