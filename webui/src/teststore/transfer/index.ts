import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTransfersRequest, GetTransfersResponse, Transfer } from './types' 

export const useTransferStore = defineStore('Transfer', {
  state: () => ({
    Transfers: {
        Transfers: [] as Array<Transfer>,
        Total: 0,
    }
  }),
  getters: {},
  actions: {
    getTransfers (req: GetTransfersRequest, done: (error: boolean, rows: Transfer[]) => void) {
      doActionWithError<GetTransfersRequest, GetTransfersResponse>(
        API.GET_TRANSFERS,
        req,
        req.Message,
        (resp: GetTransfersResponse): void => {
          this.Transfers.Transfers = resp.Infos
          this.Transfers.Total = resp.Total
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
  }
})