import type { NextPage } from 'next'
import Head from 'next/head'
import React from 'react'
import LicenseListView from '../../src/components/LicenseListView'

const create: NextPage = () => {
    return (
        <>
            <Head>
                <title>Sistem Terdistribusi</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <LicenseListView />
        </>
    )
}

export default create