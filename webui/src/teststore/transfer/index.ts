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
    getKey() {
      return (chainID: string, chainType: string) => {
        return `${chainID}-${chainType}`
      }
    },
    getTransfers() {
      return (chainID: string, chainType: string) => {
        const key = this.getKey(chainID, chainType)
        const transfers = this.Transfers.Transfers.get(key)
        return !transfers? [] : transfers
      }
    },
  },
  actions: {
    getTransfers (req: GetTransfersRequest, done: (error: boolean, rows: Transfer[]) => void) {
      doActionWithError<GetTransfersRequest, GetTransfersResponse>(
        API.GET_TRANSFERS,
        req,
        req.Message,
        (resp: GetTransfersResponse): void => {
          resp.Infos.forEach((el) => {
            // key: ChainID-ChainType
            let transfers = this.Transfers.Transfers.get(`${el.ChainID}-${el.ChainType}`)
            if (!transfers) {
                transfers = [] as Array<Transfer>
            }
            transfers.push(el)
            this.Transfers.Transfers.set(`${el.ChainID}-${el.ChainType}`, transfers)
          })
          this.Transfers.Total = resp.Total
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    }
  }
})