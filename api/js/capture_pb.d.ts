import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';


export class Capture extends jspb.Message {
  getIface(): string;
  setIface(value: string): Capture;

  getFilter(): string;
  setFilter(value: string): Capture;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Capture.AsObject;
  static toObject(includeInstance: boolean, msg: Capture): Capture.AsObject;
  static serializeBinaryToWriter(message: Capture, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Capture;
  static deserializeBinaryFromReader(message: Capture, reader: jspb.BinaryReader): Capture;
}

export namespace Capture {
  export type AsObject = {
    iface: string,
    filter: string,
  }
}

export class Packet extends jspb.Message {
  getProtocol(): string;
  setProtocol(value: string): Packet;

  getSrc(): Address | undefined;
  setSrc(value?: Address): Packet;
  hasSrc(): boolean;
  clearSrc(): Packet;

  getDst(): Address | undefined;
  setDst(value?: Address): Packet;
  hasDst(): boolean;
  clearDst(): Packet;

  getLayers(): string;
  setLayers(value: string): Packet;

  getPayload(): Uint8Array | string;
  getPayload_asU8(): Uint8Array;
  getPayload_asB64(): string;
  setPayload(value: Uint8Array | string): Packet;

  getData(): string;
  setData(value: string): Packet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Packet.AsObject;
  static toObject(includeInstance: boolean, msg: Packet): Packet.AsObject;
  static serializeBinaryToWriter(message: Packet, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Packet;
  static deserializeBinaryFromReader(message: Packet, reader: jspb.BinaryReader): Packet;
}

export namespace Packet {
  export type AsObject = {
    protocol: string,
    src?: Address.AsObject,
    dst?: Address.AsObject,
    layers: string,
    payload: Uint8Array | string,
    data: string,
  }
}

export class Address extends jspb.Message {
  getIp(): string;
  setIp(value: string): Address;

  getPort(): number;
  setPort(value: number): Address;

  getMac(): string;
  setMac(value: string): Address;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Address.AsObject;
  static toObject(includeInstance: boolean, msg: Address): Address.AsObject;
  static serializeBinaryToWriter(message: Address, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Address;
  static deserializeBinaryFromReader(message: Address, reader: jspb.BinaryReader): Address;
}

export namespace Address {
  export type AsObject = {
    ip: string,
    port: number,
    mac: string,
  }
}

export class Point extends jspb.Message {
  getLabel(): string;
  setLabel(value: string): Point;

  getTcpCount(): number;
  setTcpCount(value: number): Point;

  getUdpCount(): number;
  setUdpCount(value: number): Point;

  getOtherCount(): number;
  setOtherCount(value: number): Point;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Point.AsObject;
  static toObject(includeInstance: boolean, msg: Point): Point.AsObject;
  static serializeBinaryToWriter(message: Point, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Point;
  static deserializeBinaryFromReader(message: Point, reader: jspb.BinaryReader): Point;
}

export namespace Point {
  export type AsObject = {
    label: string,
    tcpCount: number,
    udpCount: number,
    otherCount: number,
  }
}

export class CopyRequest extends jspb.Message {
  getIface(): string;
  setIface(value: string): CopyRequest;

  getPort(): string;
  setPort(value: string): CopyRequest;

  getDestinationsList(): Array<CopyRequest.Destination>;
  setDestinationsList(value: Array<CopyRequest.Destination>): CopyRequest;
  clearDestinationsList(): CopyRequest;
  addDestinations(value?: CopyRequest.Destination, index?: number): CopyRequest.Destination;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CopyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CopyRequest): CopyRequest.AsObject;
  static serializeBinaryToWriter(message: CopyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CopyRequest;
  static deserializeBinaryFromReader(message: CopyRequest, reader: jspb.BinaryReader): CopyRequest;
}

export namespace CopyRequest {
  export type AsObject = {
    iface: string,
    port: string,
    destinationsList: Array<CopyRequest.Destination.AsObject>,
  }

  export class Destination extends jspb.Message {
    getIp(): string;
    setIp(value: string): Destination;

    getPort(): string;
    setPort(value: string): Destination;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Destination.AsObject;
    static toObject(includeInstance: boolean, msg: Destination): Destination.AsObject;
    static serializeBinaryToWriter(message: Destination, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Destination;
    static deserializeBinaryFromReader(message: Destination, reader: jspb.BinaryReader): Destination;
  }

  export namespace Destination {
    export type AsObject = {
      ip: string,
      port: string,
    }
  }

}

