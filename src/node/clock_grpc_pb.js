// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var clock_pb = require('./clock_pb.js');

function serialize_ClockRequest(arg) {
  if (!(arg instanceof clock_pb.ClockRequest)) {
    throw new Error('Expected argument of type ClockRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_ClockRequest(buffer_arg) {
  return clock_pb.ClockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ClockResponse(arg) {
  if (!(arg instanceof clock_pb.ClockResponse)) {
    throw new Error('Expected argument of type ClockResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_ClockResponse(buffer_arg) {
  return clock_pb.ClockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ClockService = exports.ClockService = {
  whatTimeIsIt: {
    path: '/clock.Clock/WhatTimeIsIt',
    requestStream: false,
    responseStream: true,
    requestType: clock_pb.ClockRequest,
    responseType: clock_pb.ClockResponse,
    requestSerialize: serialize_ClockRequest,
    requestDeserialize: deserialize_ClockRequest,
    responseSerialize: serialize_ClockResponse,
    responseDeserialize: deserialize_ClockResponse,
  },
};

exports.ClockClient = grpc.makeGenericClientConstructor(ClockService);
