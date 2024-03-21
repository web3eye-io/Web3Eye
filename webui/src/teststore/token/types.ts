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
export type SearchTokenMessage = BaseRequest

export interface SearchTokensResponse {
    Infos: Array<SearchToken>
    Vector: Array<number>
    StorageKey: string
    Page: number
    Pages: number
    Total: number
    Limit: number
}

export interface SearchToken {
    ID: number
    EntID: string
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
    TransfersNum: number
}

export interface SiblingToken {
    ID: number
    EntID: string
    TokenID: string
    ImageURL: string
    IPFSImageURL: string
}

export interface Token {
    ID: number
    EntID: string
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
    ID: number
}
  
export interface GetTokenResponse {
    Info: Token
}
  
export interface GetTokensRequest extends BaseRequest {
    StorageKey?: string
    Vector?: Array<number>
    Page: number
    Limit: number
}

export interface GetTokensResponse {
    Infos: SearchToken[]
    StorageKey: string
    Vector: Array<number>
    Page: number
    Pages: number
    Total: number
    Limit: number
}