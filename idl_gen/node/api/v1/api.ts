/* eslint-disable */
// @generated by protobuf-ts 2.4.0 with parameter generate_dependencies,server_none,optimize_code_size,client_generic,output_typescript,long_type_string,eslint_disable,// @generated from protobuf file "api/v1/api.proto" (package "api.v1", syntax proto3),// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message api.v1.PingRequest
 */
export interface PingRequest {
}
/**
 * @generated from protobuf message api.v1.PingResponse
 */
export interface PingResponse {
    /**
     * @generated from protobuf field: string message = 1;
     */
    message: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class PingRequest$Type extends MessageType<PingRequest> {
    constructor() {
        super("api.v1.PingRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message api.v1.PingRequest
 */
export const PingRequest = new PingRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PingResponse$Type extends MessageType<PingResponse> {
    constructor() {
        super("api.v1.PingResponse", [
            { no: 1, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message api.v1.PingResponse
 */
export const PingResponse = new PingResponse$Type();
/**
 * @generated ServiceType for protobuf service api.v1.API
 */
export const API = new ServiceType("api.v1.API", [
    { name: "Ping", options: {}, I: PingRequest, O: PingResponse }
]);
