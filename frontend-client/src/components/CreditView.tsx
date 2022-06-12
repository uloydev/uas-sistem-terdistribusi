import Image from 'next/image'
import React from 'react'
import Template from './Template'
import WrapperCard from './WrapperCard'

import { BsGithub, BsInstagram } from 'react-icons/bs'

const CreditView = () => {
    return (
        <Template title='Credits'>
            <WrapperCard className=''>
                <div className='space-y-10 my-2'>
                    <div className='flex lg:gap-x-8 ml-5 gap-x-4'>
                        <div className='w-40 h-40 relative rounded-full overflow-hidden'>
                            <Image src="/uloydev.jpg" objectFit='cover' layout='fill' />
                        </div>
                        <div className='font-bold text-slate-300'>
                            <h4 className='text-orange-400 text-lg'>Name : Wahyu Miftahul Aflah</h4>
                            <p>Username : Uloydev</p>
                            <p>Email : uloydev@gmail.com</p>
                            <p>Role : Backend Developer</p>
                            <a href="https://github.com/uloydev" rel='noreferrer' target="_blank" className='mt-2 rounded-md py-2 text-dark px-4 flex items-center bg-orange-400 w-36'>
                                <BsGithub size={20} />
                                <span className='ml-3'>Github</span>
                            </a>
                        </div>
                    </div>
                    <div className='flex lg:gap-x-8 ml-5 gap-x-4'>
                        <div className='w-40 h-40 relative rounded-full brightness-75 overflow-hidden'>
                            <Image src="/ijordev.jpeg" objectFit='cover' layout='fill' />
                        </div>
                        <div className='font-bold text-slate-300'>
                            <h4 className='text-orange-400 text-lg'>Name : Muhammad Fathurrahman</h4>
                            <p>Username : IjorDev</p>
                            <p>Email : gazoon3@gmail.com</p>
                            <p>Role : Frontend Developer</p>
                            <a href="https://github.com/fathrahh" rel='noreferrer' target="_blank" className='mt-2 rounded-md py-2 text-dark px-4 flex items-center bg-orange-400 w-36'>
                                <BsGithub size={20} />
                                <span className='ml-3'>Github</span>
                            </a>
                        </div>
                    </div>
                    <div className='flex lg:gap-x-8 ml-5 gap-x-4 '>
                        <div className='w-40 h-40 relative rounded-full overflow-hidden'>
                            <Image className='scale-150' src="/hafidzdev.jpg" objectFit='cover' layout='fill' />
                        </div>
                        <div className='font-bold text-slate-300'>
                            <h4 className='text-orange-400 text-lg'>Name : Hafidz Ashabi Muhammad</h4>
                            <p>Username : Hafidz Design</p>
                            <p>Email : hafidzashabi@upnvj.ac.id</p>
                            <p>Role : UI/UX | Design Grafis</p>
                            <a href="https://www.instagram.com/hafidzashabi/" rel='noreferrer' target="_blank" className='mt-2 rounded-md py-2 text-dark px-4 flex items-center bg-orange-400 w-36'>
                                <BsInstagram size={20} />
                                <span className='ml-3'>Instagram</span>
                            </a>
                        </div>
                    </div>
                </div>
            </WrapperCard>
        </Template>
    )
}

export default CreditView