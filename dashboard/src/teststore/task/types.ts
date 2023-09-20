import { ChainType } from '../basetypes/const';
import { SyncState } from '../basetypes/synctask/const';
import { BaseRequest } from '../local';

export interface SyncTask {
    ID: string;
    ChainType: ChainType;
    ChainID: string;
    /** @format uint64 */
    Start: string;
    /** @format uint64 */
    End: string;
    /** @format uint64 */
    Current: string;
    Topic: string;
    Description: string;
    SyncState: SyncState;
    Remark: string;
}

export interface CreateSyncTaskRequest extends BaseRequest{
    ID: string;
    ChainType: ChainType;
    ChainID: string;
    /** @format uint64 */
    Start: string;
    /** @format uint64 */
    End: string;
    /** @format uint64 */
    Current: string;
    Topic: string;
    Description: string;
    SyncState: SyncState;
    Remark: string;
  }
  
  export interface CreateSyncTaskResponse{
    Info: SyncTask;
  }
  
  export interface DeleteSyncTaskRequest extends BaseRequest{
    ID: string;
  }
  
  export interface DeleteSyncTaskResponse {
    Info: SyncTask;
  }
  
  export interface UpdateSyncTaskRequest extends BaseRequest{
    ID: string;
    ChainType: ChainType;
    ChainID: string;
    /** @format uint64 */
    Start: string;
    /** @format uint64 */
    End: string;
    /** @format uint64 */
    Current: string;
    Topic: string;
    Description: string;
    SyncState: SyncState;
    Remark: string;
  }
  
  export interface UpdateSyncTaskResponse {
    Info: SyncTask;
  }

  export interface GetSyncTasksRequest extends BaseRequest{
    /** @format int32 */
    Offset: number;
    /** @format int32 */
    Limit: number;
  }
  
  export interface GetSyncTasksResponse {
    Infos: SyncTask[];
    /** @format int64 */
    Total: number;
  }
  