import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"


export class ReadProviderRequest extends jspb.Message {
  getId(): number;
  setId(value: number): ReadProviderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadProviderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadProviderRequest): ReadProviderRequest.AsObject;
  static serializeBinaryToWriter(message: ReadProviderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadProviderRequest;
  static deserializeBinaryFromReader(message: ReadProviderRequest, reader: jspb.BinaryReader): ReadProviderRequest;
}

export namespace ReadProviderRequest {
  export type AsObject = {
    id: number,
  }
}

export class ReadProviderResponse extends jspb.Message {
  getId(): number;
  setId(value: number): ReadProviderResponse;

  getName(): string;
  setName(value: string): ReadProviderResponse;

  getCity(): string;
  setCity(value: string): ReadProviderResponse;

  getRegion(): string;
  setRegion(value: string): ReadProviderResponse;

  getLineOfBusiness(): string;
  setLineOfBusiness(value: string): ReadProviderResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadProviderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReadProviderResponse): ReadProviderResponse.AsObject;
  static serializeBinaryToWriter(message: ReadProviderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadProviderResponse;
  static deserializeBinaryFromReader(message: ReadProviderResponse, reader: jspb.BinaryReader): ReadProviderResponse;
}

export namespace ReadProviderResponse {
  export type AsObject = {
    id: number,
    name: string,
    city: string,
    region: string,
    lineOfBusiness: string,
  }
}

export class CreateProviderRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateProviderRequest;

  getCity(): string;
  setCity(value: string): CreateProviderRequest;

  getRegion(): string;
  setRegion(value: string): CreateProviderRequest;

  getLineOfBusiness(): string;
  setLineOfBusiness(value: string): CreateProviderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateProviderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateProviderRequest): CreateProviderRequest.AsObject;
  static serializeBinaryToWriter(message: CreateProviderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateProviderRequest;
  static deserializeBinaryFromReader(message: CreateProviderRequest, reader: jspb.BinaryReader): CreateProviderRequest;
}

export namespace CreateProviderRequest {
  export type AsObject = {
    name: string,
    city: string,
    region: string,
    lineOfBusiness: string,
  }
}

export class CreateProviderResponse extends jspb.Message {
  getId(): number;
  setId(value: number): CreateProviderResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateProviderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateProviderResponse): CreateProviderResponse.AsObject;
  static serializeBinaryToWriter(message: CreateProviderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateProviderResponse;
  static deserializeBinaryFromReader(message: CreateProviderResponse, reader: jspb.BinaryReader): CreateProviderResponse;
}

export namespace CreateProviderResponse {
  export type AsObject = {
    id: number,
  }
}

export class ListProviderResponse extends jspb.Message {
  getProvidersList(): Array<ReadProviderResponse>;
  setProvidersList(value: Array<ReadProviderResponse>): ListProviderResponse;
  clearProvidersList(): ListProviderResponse;
  addProviders(value?: ReadProviderResponse, index?: number): ReadProviderResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProviderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListProviderResponse): ListProviderResponse.AsObject;
  static serializeBinaryToWriter(message: ListProviderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProviderResponse;
  static deserializeBinaryFromReader(message: ListProviderResponse, reader: jspb.BinaryReader): ListProviderResponse;
}

export namespace ListProviderResponse {
  export type AsObject = {
    providersList: Array<ReadProviderResponse.AsObject>,
  }
}

export class RefreshTokenRequest extends jspb.Message {
  getRefreshToken(): string;
  setRefreshToken(value: string): RefreshTokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshTokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshTokenRequest): RefreshTokenRequest.AsObject;
  static serializeBinaryToWriter(message: RefreshTokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshTokenRequest;
  static deserializeBinaryFromReader(message: RefreshTokenRequest, reader: jspb.BinaryReader): RefreshTokenRequest;
}

export namespace RefreshTokenRequest {
  export type AsObject = {
    refreshToken: string,
  }
}

export class RegistrationRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): RegistrationRequest;

  getPassword(): string;
  setPassword(value: string): RegistrationRequest;

  getUsername(): string;
  setUsername(value: string): RegistrationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegistrationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegistrationRequest): RegistrationRequest.AsObject;
  static serializeBinaryToWriter(message: RegistrationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegistrationRequest;
  static deserializeBinaryFromReader(message: RegistrationRequest, reader: jspb.BinaryReader): RegistrationRequest;
}

export namespace RegistrationRequest {
  export type AsObject = {
    email: string,
    password: string,
    username: string,
  }
}

export class LoginRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    email: string,
    password: string,
  }
}

export class TokenResponse extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): TokenResponse;

  getExp(): number;
  setExp(value: number): TokenResponse;

  getTokenType(): string;
  setTokenType(value: string): TokenResponse;

  getRefreshToken(): string;
  setRefreshToken(value: string): TokenResponse;
  hasRefreshToken(): boolean;
  clearRefreshToken(): TokenResponse;

  getScope(): string;
  setScope(value: string): TokenResponse;
  hasScope(): boolean;
  clearScope(): TokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TokenResponse): TokenResponse.AsObject;
  static serializeBinaryToWriter(message: TokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenResponse;
  static deserializeBinaryFromReader(message: TokenResponse, reader: jspb.BinaryReader): TokenResponse;
}

export namespace TokenResponse {
  export type AsObject = {
    accessToken: string,
    exp: number,
    tokenType: string,
    refreshToken?: string,
    scope?: string,
  }

  export enum RefreshTokenCase { 
    _REFRESH_TOKEN_NOT_SET = 0,
    REFRESH_TOKEN = 4,
  }

  export enum ScopeCase { 
    _SCOPE_NOT_SET = 0,
    SCOPE = 5,
  }
}

