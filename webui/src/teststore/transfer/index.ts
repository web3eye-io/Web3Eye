import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTransfersRequest, GetTransfersResponse, Transfer } from './types' 

export const useTransferStore = defineStore('Transfer', {
  state: () => ({
    Transfers: {
      Transfers: new Map<string, Array<Transfer>>(),
      Total: 0,
    }
  }),
  getters: {
    setKey() {
      return (chainID: string, contract: string, tokenID: string) => {
        // when get contract transfers, TokenID is Contract
        return `${chainID}-${contract}-${tokenID}`
      }
    },
    getTransfersByKey() {
      return (key: string) => {
        const transfers = this.Transfers.Transfers.get(key)
        return !transfers? [] : transfers
      }
    },
  },
  actions: {
    getTransfers (req: GetTransfersRequest, key: string, done: (error: boolean, rows: Transfer[]) => void) {
      doActionWithError<GetTransfersRequest, GetTransfersResponse>(
        API.GET_TRANSFERS,
        req,
        req.Message,
        (resp: GetTransfersResponse): void => {
          resp.Infos.forEach((el) => {
            let transfers = this.Transfers.Transfers.get(key)
            if (!transfers) {
                transfers = [] as Array<Transfer>
            }
            const index = transfers.findIndex((al) => al.ID === el.ID)
            transfers.splice(index > -1 ? index : 0, index > -1 ? 1 : 0, el)
            this.Transfers.Transfers.set(key, transfers)
          })
          this.Transfers.Total = resp.Total
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    }
  }
})