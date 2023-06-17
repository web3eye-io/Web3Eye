import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { StartRetrieveRequest, StartRetrieveResponse, Retrieve, StatRetrieveRequest, StatRetrieveResponse } from './types'

export const useRetrieveStore = defineStore('Retrieve', {
  state: () => ({
    Retrieves: {
      Retrieves: [] as Array<Retrieve>,
      Total: 0,
    }
  }),
  getters: {},
  actions: {
    startRetrieve (req: StartRetrieveRequest, done: (error: boolean, rows: Retrieve) => void) {
      doActionWithError<StartRetrieveRequest, StartRetrieveResponse>(
        API.START_RETRIEVE,
        req,
        req.Message,
        (resp: StartRetrieveResponse): void => {
          this.Retrieves.Retrieves.push(resp.Info)
          done(false, resp.Info)
        }, () => {
          done(true, {} as Retrieve)
      })
    },
    statRetrieve (req: StatRetrieveRequest, done: (error: boolean, row: Retrieve) => void) {
      doActionWithError<StatRetrieveRequest, StatRetrieveResponse>(
        API.STAT_RETRIEVE,
        req,
        req.Message,
        (resp: StatRetrieveResponse): void => {
          const index = this.Retrieves.Retrieves.findIndex((el) => el.TokenID === resp.Info.TokenID)
          this.Retrieves.Retrieves.splice(index < 0 ? 0 : index, index < 0 ? 0 : 1, resp.Info)
          this.Retrieves.Total += 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as Retrieve)
      })
    }
  }
})
export * from './const'
export * from './types'