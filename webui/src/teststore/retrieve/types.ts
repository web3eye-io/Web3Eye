import { BaseRequest } from '../local'

export interface Retrieve {
  ChainType: string
  ChainID  : string
  Contract : string
  TokenID  : string
  RetrieveState: string
  ProposalID: string
  DealID    : number
  BackupPayloadCID: string
}

export interface StartRetrieveRequest extends BaseRequest {
  ChainType: string;
  ChainID  : string;
  Contract : string;
  TokenID  : string;
}

export interface StartRetrieveResponse {
  Info: Retrieve 
}

export interface StatRetrieveRequest extends BaseRequest {
  ChainType: string;
  ChainID  : string;
  Contract : string;
  TokenID  : string;
}

export interface StatRetrieveResponse {
  Info:  Retrieve
}