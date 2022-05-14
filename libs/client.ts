import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport"
import { APIClient, IAPIClient } from "../idl_gen/node/api/v1/api.client"
import useSWR from "swr"
import { RpcOptions, UnaryCall } from "@protobuf-ts/runtime-rpc"
import { KeyedMutator } from "swr/dist/types"

function baseURL() {
  if (typeof window !== "undefined") {
    return window.location.origin
  }
  return "http://localhost:3000"
}

export const API = new APIClient(
  new GrpcWebFetchTransport({
    baseUrl: baseURL(),
    format: "text",
  })
)

type UnaryCallData<T extends UnaryCall> = T extends UnaryCall<infer R, infer P>
  ? P
  : never

export function useAPI<
  KEY extends keyof IAPIClient,
  FUN extends IAPIClient[KEY],
  REQ extends Parameters<FUN>[0],
  RESP extends UnaryCallData<ReturnType<FUN>>
>(
  method: KEY,
  input: REQ,
  opts?: RpcOptions
): {
  data?: RESP
  isValidating: boolean
  error?: any
  mutate: KeyedMutator<RESP>
} {
  return useSWR<
    RESP,
    any,
    {
      method: KEY
      input: REQ
      opts?: RpcOptions
    }
  >({ method, input, opts }, async ({ method, input, opts }): Promise<RESP> => {
    const call = API[method](input, opts)
    return call.response as any
  })
}
