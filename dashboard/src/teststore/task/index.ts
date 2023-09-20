import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { 
  GetSyncTasksRequest, 
  GetSyncTasksResponse, 
  SyncTask, 
  CreateSyncTaskRequest, 
  CreateSyncTaskResponse, 
  UpdateSyncTaskRequest, 
  UpdateSyncTaskResponse, 
  DeleteSyncTaskRequest, 
  DeleteSyncTaskResponse 
} from './types'

export const useSyncTaskStore = defineStore('synctask', {
  state: () => ({
    SyncTasks: {
      SyncTasks: [] as Array<SyncTask>,
      Total: 0,
    }
  }),
  getters: {},
  actions: {
    getSyncTasks (req: GetSyncTasksRequest, done: (error: boolean, rows: SyncTask[]) => void) {
      doActionWithError<GetSyncTasksRequest, GetSyncTasksResponse>(
        API.GET_SYNCTASKS,
        req,
        req.Message,
        (resp: GetSyncTasksResponse): void => {
          this.SyncTasks.SyncTasks = resp.Infos
          this.SyncTasks.Total = resp.Total
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
    createSyncTask (req: CreateSyncTaskRequest, done: (error: boolean, row?: SyncTask) => void) {
      doActionWithError<CreateSyncTaskRequest, CreateSyncTaskResponse>(
        API.CREATE_SYNCTASK,
        req,
        req.Message,
        (resp: CreateSyncTaskResponse): void => {
          this.SyncTasks.SyncTasks.push(resp.Info)
          this.SyncTasks.Total += 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as SyncTask)
      })
    },
    updateSyncTask (req: UpdateSyncTaskRequest, done: (error: boolean, row?: SyncTask) => void) {
      doActionWithError<UpdateSyncTaskRequest, UpdateSyncTaskResponse>(
        API.UPDATE_SYNCTASK,
        req,
        req.Message,
        (resp: UpdateSyncTaskResponse): void => {
          const index = this.SyncTasks.SyncTasks.findIndex((el) => el.ID === resp.Info.ID)
          this.SyncTasks.SyncTasks.splice(index > 0 ? index : 0, index > 0 ? 1 : 0, resp.Info)
          this.SyncTasks.Total += 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as SyncTask)
      })
    },
    deleteSyncTask (req: DeleteSyncTaskRequest, done: (error: boolean, row?: SyncTask) => void) {
      doActionWithError<DeleteSyncTaskRequest, DeleteSyncTaskResponse>(
        API.DELETE_SYNCTASK,
        req,
        req.Message,
        (resp: DeleteSyncTaskResponse): void => {
          const index = this.SyncTasks.SyncTasks.findIndex((el) => el.ID === resp.Info.ID)
          if(index < 0) return
          this.SyncTasks.SyncTasks.splice(index, 1)
          this.SyncTasks.Total -= 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as SyncTask)
      })
    },
  }
})