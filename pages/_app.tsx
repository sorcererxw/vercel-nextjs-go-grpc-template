import React from "react"
import { AppProps } from "next/app"
import Head from "next/head"
import "../styles/global.css"

const MyApp = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Head>
        <link rel="icon" type="image/png" sizes="16x16" href="/favicon.png" />
        <title>Template</title>
      </Head>
      <Component {...pageProps} />
    </>
  )
}

export default MyApp
