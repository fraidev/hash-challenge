// package: discount
// file: discount.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as discount_pb from "./discount_pb";

interface IDiscountServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getDiscount: IDiscountServiceService_IGetDiscount;
}

interface IDiscountServiceService_IGetDiscount extends grpc.MethodDefinition<discount_pb.DiscountRequest, discount_pb.DiscountReply> {
    path: "/discount.DiscountService/GetDiscount";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<discount_pb.DiscountRequest>;
    requestDeserialize: grpc.deserialize<discount_pb.DiscountRequest>;
    responseSerialize: grpc.serialize<discount_pb.DiscountReply>;
    responseDeserialize: grpc.deserialize<discount_pb.DiscountReply>;
}

export const DiscountServiceService: IDiscountServiceService;

export interface IDiscountServiceServer extends grpc.UntypedServiceImplementation {
    getDiscount: grpc.handleUnaryCall<discount_pb.DiscountRequest, discount_pb.DiscountReply>;
}

export interface IDiscountServiceClient {
    getDiscount(request: discount_pb.DiscountRequest, callback: (error: grpc.ServiceError | null, response: discount_pb.DiscountReply) => void): grpc.ClientUnaryCall;
    getDiscount(request: discount_pb.DiscountRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: discount_pb.DiscountReply) => void): grpc.ClientUnaryCall;
    getDiscount(request: discount_pb.DiscountRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: discount_pb.DiscountReply) => void): grpc.ClientUnaryCall;
}

export class DiscountServiceClient extends grpc.Client implements IDiscountServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getDiscount(request: discount_pb.DiscountRequest, callback: (error: grpc.ServiceError | null, response: discount_pb.DiscountReply) => void): grpc.ClientUnaryCall;
    public getDiscount(request: discount_pb.DiscountRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: discount_pb.DiscountReply) => void): grpc.ClientUnaryCall;
    public getDiscount(request: discount_pb.DiscountRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: discount_pb.DiscountReply) => void): grpc.ClientUnaryCall;
}
