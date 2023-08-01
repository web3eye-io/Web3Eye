import { ChainTokenType, ChainType } from '../basetypes/const';
import { BaseRequest } from '../local';

export interface Contract {
  ID: string;
  ChainType: ChainType;
  ChainID: string;
  Address: string;
  Name: string;
  Symbol: string;
  Creator: string;
  BlockNum: string;
  TxHash: string;
  TxTime: number;
  ProfileURL: string;
  BaseURL: string;
  BannerURL: string;
  Description: string;
  Remark: string;
}

export interface ShotToken {
  ID: string;
  ChainType: ChainType;
  TokenType: ChainTokenType;
  TokenID: string;
  Owner: string;
  ImageURL: string;
  Name: string;
  IPFSImageURL: string;
  ImageSnapshotID: string;
  TransfersNum: number;
}

export interface GetContractAndTokensRequest extends BaseRequest {
  Contract: string;
  Offset: number;
  Limit: number;
}

export interface GetContractAndTokensResponse {
  Contract: Contract;
  Tokens: ShotToken[];
  TotalTokens: number;
}