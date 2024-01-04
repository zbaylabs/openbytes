import * as jspb from 'google-protobuf'



export class Capture extends jspb.Message {
  getIface(): string;
  setIface(value: string): Capture;

  getProtocal(): string;
  setProtocal(value: string): Capture;

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
    protocal: string,
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

