import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb'; // proto import: "google/protobuf/struct.proto"
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb'; // proto import: "google/protobuf/wrappers.proto"
import * as capture_pb from './capture_pb'; // proto import: "capture.proto"


export class CapturesClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  device(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_struct_pb.ListValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_struct_pb.ListValue>;

  list(
    request: capture_pb.Capture,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<capture_pb.Packet>;

  traffic(
    request: capture_pb.Capture,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<capture_pb.Point>;

  copy(
    request: capture_pb.CopyRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_wrappers_pb.StringValue) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_wrappers_pb.StringValue>;

}

export class CapturesPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  device(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_struct_pb.ListValue>;

  list(
    request: capture_pb.Capture,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<capture_pb.Packet>;

  traffic(
    request: capture_pb.Capture,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<capture_pb.Point>;

  copy(
    request: capture_pb.CopyRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_wrappers_pb.StringValue>;

}

