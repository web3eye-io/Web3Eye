import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTransferRequest, GetTransferResponse, GetTransfersRequest, GetTransfersResponse, Transfer } from './types' 

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
    getTransfer (req: GetTransferRequest, done: (error: boolean, row: Transfer) => void) {
      doActionWithError<GetTransferRequest, GetTransferResponse>(
        API.GET_TRANSFER,
        req,
        req.Message,
        (resp: GetTransferResponse): void => {
          const index = this.Transfers.Transfers.findIndex((al) => al.ID === resp.Info?.ID)
          this.Transfers.Transfers.splice(index > -1 ? index : 0, index > -1 ? 1 : 0, resp.Info)
          done(false, resp.Info)
        }, () => {
          done(true, {} as Transfer)
      })
    },
  }
})