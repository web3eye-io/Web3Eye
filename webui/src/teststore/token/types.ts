import { ChainType, TokenType } from '../basetypes/const'
import { BaseRequest } from '../local'

export enum ConvertState {
    Default = 'Default',
    Waiting = 'Waiting',
    Processing = 'Processing',
    Pause = 'Pause',
    Success = 'Success',
    Failed = 'Failed',
}
  
export interface SearchToken {
    ID: string
    ChainType: ChainType
    ChainID: string
    Contract: string
    TokenType: string
    TokenID: string
    Owner: string
    URI: string
    URIType: string
    ImageURL: string
    VideoURL: string
    Description: string
    Name: string
    VectorState: ConvertState
    VectorID: string
    Remark: string
    IPFSImageURL: string
    ImageSnapshotID: string
    SiblingTokens: SiblingToken[]
    SiblingsNum: number
    Distance: number
}

export interface SiblingToken {
    ID: string
    TokenID: string
    ImageURL: string
    IPFSImageURL: string
}

export interface Token {
    ID: string
    ChainType: ChainType
    ChainID: string
    Contract: string
    TokenType: TokenType
    TokenID: string
    Owner: string
    URI: string
    URIType: string
    ImageURL: string
    VideoURL: string
    Description: string
    Name: string
    VectorState: ConvertState
    VectorID: string
    Remark: string
    IPFSImageURL: string
    ImageSnapshotID: string
}

export interface GetTokenRequest extends BaseRequest {
    ID: string
}
  
export interface GetTokenResponse {
    Info: Token
}
  
export interface GetTokensRequest extends BaseRequest {
    StorageKey: string
    Page: number
}

export interface GetTokensResponse {
    Infos: SearchToken[]
    StorageKey: string
    Page: number
    TotalPages: number
    TotalTokens: number
    Limit: number
}