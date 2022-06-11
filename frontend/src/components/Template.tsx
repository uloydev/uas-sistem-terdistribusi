import React from 'react'
import { useState } from 'react'
import Container from './Container'
import Footer from './Footer'
import Header from './Header'
import Sidebar from './Sidebar'

interface TemplateProps {
    title: string
    children: React.ReactNode
}

const Template: React.FC<TemplateProps> = ({ title, children }) => {
    const [sidebarActive, setSidebarActive] = useState(false)

    return (
        <div className="bg-dark h-screen overflow-hidden relative w-screen">
            <div className='flex h-full'>
                <Sidebar open={sidebarActive} />
                {/* content here */}
                <div className={`flex-1 flex flex-col overflow-y-scroll h-full pt-4 text-white ${sidebarActive ? "translate-x-72" : ""}`}>
                    <Header sidebarToggle={() => setSidebarActive(!sidebarActive)} title={title} />
                    <Container className="flex-1 mb-10">
                        {children}
                    </Container>
                    <Footer />
                </div>
            </div>
        </div>
    )
}

export default Template