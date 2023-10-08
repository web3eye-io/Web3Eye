import { ChainType, OrderItemType, TokenType } from '../basetypes/const'
import { BaseRequest } from '../local'

export interface OrderItem {
  Contract: string
  TokenType: TokenType
  TokenID: string
  Amount: string
  AmountStr: string
  Remark: string
  Name: string
  Symbol: string
  Decimals: number
  OrderItemType: OrderItemType
  ImageURL: string
}

export interface Transfer {
  ID: string
  ChainType: ChainType
  ChainID: string
  Contract: string
  TokenType: string
  TokenID: string
  From: string
  To: string
  Amount: string
  BlockNumber: string
  TxHash: string
  BlockHash: string
  TxTime: number
  Remark: string
  OfferItems: OrderItem[]
  TargetItems: OrderItem[]
}

export interface GetTransferRequest extends BaseRequest {
  ID: string
}

export interface GetTransferResponse {
  Info: Transfer
}

export interface GetTransfersRequest extends BaseRequest {
  ChainType: ChainType
  ChainID: string
  Contract: string
  TokenID?: string
  Offset: number
  Limit: number
}

export interface GetTransfersResponse {
  Infos: Transfer[]
  Total: number
}
