/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./word_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.WordAPIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.WordAPIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Empty,
 *   !proto.WordCounts>}
 */
const methodDescriptor_WordAPI_GetWords = new grpc.web.MethodDescriptor(
  '/WordAPI/GetWords',
  grpc.web.MethodType.UNARY,
  proto.Empty,
  proto.WordCounts,
  /**
   * @param {!proto.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.WordCounts.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Empty,
 *   !proto.WordCounts>}
 */
const methodInfo_WordAPI_GetWords = new grpc.web.AbstractClientBase.MethodInfo(
  proto.WordCounts,
  /**
   * @param {!proto.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.WordCounts.deserializeBinary
);


/**
 * @param {!proto.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.WordCounts)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.WordCounts>|undefined}
 *     The XHR Node Readable Stream
 */
proto.WordAPIClient.prototype.getWords =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/WordAPI/GetWords',
      request,
      metadata || {},
      methodDescriptor_WordAPI_GetWords,
      callback);
};


/**
 * @param {!proto.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.WordCounts>}
 *     Promise that resolves to the response
 */
proto.WordAPIPromiseClient.prototype.getWords =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/WordAPI/GetWords',
      request,
      metadata || {},
      methodDescriptor_WordAPI_GetWords);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Word,
 *   !proto.Empty>}
 */
const methodDescriptor_WordAPI_UpdateWord = new grpc.web.MethodDescriptor(
  '/WordAPI/UpdateWord',
  grpc.web.MethodType.UNARY,
  proto.Word,
  proto.Empty,
  /**
   * @param {!proto.Word} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Word,
 *   !proto.Empty>}
 */
const methodInfo_WordAPI_UpdateWord = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Empty,
  /**
   * @param {!proto.Word} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Empty.deserializeBinary
);


/**
 * @param {!proto.Word} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.WordAPIClient.prototype.updateWord =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/WordAPI/UpdateWord',
      request,
      metadata || {},
      methodDescriptor_WordAPI_UpdateWord,
      callback);
};


/**
 * @param {!proto.Word} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Empty>}
 *     Promise that resolves to the response
 */
proto.WordAPIPromiseClient.prototype.updateWord =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/WordAPI/UpdateWord',
      request,
      metadata || {},
      methodDescriptor_WordAPI_UpdateWord);
};


module.exports = proto;

