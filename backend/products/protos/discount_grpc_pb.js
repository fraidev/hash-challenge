// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var discount_pb = require('./discount_pb.js');

function serialize_discount_DiscountReply(arg) {
  if (!(arg instanceof discount_pb.DiscountReply)) {
    throw new Error('Expected argument of type discount.DiscountReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_discount_DiscountReply(buffer_arg) {
  return discount_pb.DiscountReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_discount_DiscountRequest(arg) {
  if (!(arg instanceof discount_pb.DiscountRequest)) {
    throw new Error('Expected argument of type discount.DiscountRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_discount_DiscountRequest(buffer_arg) {
  return discount_pb.DiscountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var DiscountServiceService = exports.DiscountServiceService = {
  getDiscount: {
    path: '/discount.DiscountService/GetDiscount',
    requestStream: false,
    responseStream: false,
    requestType: discount_pb.DiscountRequest,
    responseType: discount_pb.DiscountReply,
    requestSerialize: serialize_discount_DiscountRequest,
    requestDeserialize: deserialize_discount_DiscountRequest,
    responseSerialize: serialize_discount_DiscountReply,
    responseDeserialize: deserialize_discount_DiscountReply,
  },
};

exports.DiscountServiceClient = grpc.makeGenericClientConstructor(DiscountServiceService);
