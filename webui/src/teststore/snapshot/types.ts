import { BaseRequest } from '../local';

export interface ContentItem {
  ID: string;
  URI: string;
  ChainType: string;
  ChainID: string;
  Contract: string;
  TokenID: string;
  FileName: string;
}

export enum BackupState {
  BackupStateNone  = 'BackupStateNone',
  BackupStateCreated  = 'BackupStateCreated',
  BackupStateProposed = 'BackupStateProposed',
  BackupStateAccepted = 'BackupStateAccepted',
  BackupStateSuccess  = 'BackupStateSuccess',
  BackupStateFail     = 'BackupStateFail',
}

export interface Snapshot {
  ID: string;
  Index: number;
  SnapshotCommP: string;
  SnapshotRoot: string;
  SnapshotURI: string;
  Items: ContentItem[];
  BackupState: BackupState;
  ProposalCID: string;
  DealID: number;
  // just for frontend
  Loading: boolean;
}

export interface GetSnapshotsRequest extends BaseRequest {
  Indexes?: number[];
}

export interface GetSnapshotsResponse {
  Infos: Snapshot[];
  Total: number;
}

export interface CreateBackupRequest extends BaseRequest {
  Index: number;
}

export interface CreateBackupResponse {
  Info: Snapshot;
}