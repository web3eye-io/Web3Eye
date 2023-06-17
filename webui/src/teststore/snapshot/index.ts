import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetSnapshotsRequest, GetSnapshotsResponse, Snapshot, CreateBackupRequest, CreateBackupResponse } from './types'

export const useSnapshotStore = defineStore('snapshot', {
  state: () => ({
    Snapshots: {
      Snapshots: [] as Array<Snapshot>,
      Total: 0,
    }
  }),
  getters: {},
  actions: {
    getSnapshots (req: GetSnapshotsRequest, done: (error: boolean, rows: Snapshot[]) => void) {
      doActionWithError<GetSnapshotsRequest, GetSnapshotsResponse>(
        API.GET_SNAPSHOTS,
        req,
        req.Message,
        (resp: GetSnapshotsResponse): void => {
          this.Snapshots.Snapshots = resp.Infos
          this.Snapshots.Total = resp.Total
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
    createBackup (req: CreateBackupRequest, done: (error: boolean, row: Snapshot) => void) {
      doActionWithError<CreateBackupRequest, CreateBackupResponse>(
        API.CREATE_BACKUP,
        req,
        req.Message,
        (resp: CreateBackupResponse): void => {
          this.Snapshots.Snapshots.push(resp.Info)
          this.Snapshots.Total += 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as Snapshot)
      })
    }
  }
})