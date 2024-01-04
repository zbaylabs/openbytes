import * as grpcWeb from 'grpc-web';

import * as capture_pb from './capture_pb';


export class CapturesClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  list(
    request: capture_pb.Capture,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<capture_pb.Packet>;

}

export class CapturesPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  list(
    request: capture_pb.Capture,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<capture_pb.Packet>;

}

