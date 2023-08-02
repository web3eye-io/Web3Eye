import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { Contract, GetContractAndTokensRequest, GetContractAndTokensResponse, ShotToken } from './types'

export const useContractStore = defineStore('contract', {
  state: () => ({
    ShotTokens: {
      ShotTokens: [] as Array<ShotToken>,
      Total: 0,
    },
    Contract: {} as Contract
  }),
  getters: {},
  actions: {
    getContractAndTokens (req: GetContractAndTokensRequest, done: (error: boolean, rows: Contract, tokens: ShotToken[]) => void) {
      doActionWithError<GetContractAndTokensRequest, GetContractAndTokensResponse>(
        API.GET_CONTRACT_AND_TOKENS,
        req,
        req.Message,
        (resp: GetContractAndTokensResponse): void => {
          this.ShotTokens.ShotTokens = resp.Tokens
          this.Contract = resp.Contract
          done(false, resp.Contract, resp.Tokens)
        }, () => {
          done(true, {} as Contract, [])
      })
    },
  }
})