import type { NextPage } from 'next'
import Head from 'next/head'
import CreditView from '../src/components/CreditView'

const Home: NextPage = () => {
    return (
        <>
            <Head>
                <title>Sistem Terdistribusi</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <CreditView />
        </>
    )
}

export default Home