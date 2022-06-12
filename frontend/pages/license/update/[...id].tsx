import type { NextPage } from 'next'
import Head from 'next/head'
import React from 'react'
import LicenseEditView from '../../../src/components/LicenseEditView'

const Update: NextPage = () => {
    return (
        <>
            <Head>
                <title>Sistem Terdistribusi</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <LicenseEditView />
        </>
    )
}

export default Update