export interface NFTMeta {
  ChainType: string;
  Contract: string;
  Description: string;
  Distance: number;
  ID: string;
  ImageURL: string;
  Name: string;
  TokenID: string;
  URI: string;
  URIType: string;
  VectorID:string;
  VectorState: number;
}

export interface UploadResponse {
  data: Array<NFTMeta>
  info: string
}