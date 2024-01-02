export enum ImageState {
  Normal = 'Normal',
  IPFS = 'IPFS',
  Retrieving = 'Retrieving',
  WaitRecover = 'WaitRecover'
}

export interface NFTMeta {
  ChainType: string;
  ChainID: string;
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
  IPFSImageURL: string;
  ImageSnapshotID: string;
  VideoURL: string;
  // just for frontend
  ImageState: ImageState;
  Loading: boolean;
  LoadError: boolean
}

export interface UploadResponse {
  Infos: Array<NFTMeta>
  Msg: string
  Page: number
  StorageKey : string
  Pages: number
  Total: number
  Limit: number
}