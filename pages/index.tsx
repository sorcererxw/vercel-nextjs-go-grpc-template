import React from "react"
import { useAPI } from "../libs/client"

const Page = () => {
  const { data, isValidating } = useAPI("ping", {})
  return <div>{isValidating ? "loading..." : data?.message}</div>
}

export default Page
