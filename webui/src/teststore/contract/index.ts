import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { Contract, GetContractAndTokensRequest, GetContractAndTokensResponse, ShotToken } from './types'

export const useContractStore = defineStore('contract', {
  state: () => ({
    ShotTokens: {
      ShotTokens: new Map<string, Array<ShotToken>>(),
      Total: 0,
    },
    Contract: {} as Contract
  }),
  getters: {
    addShotTokens (): (contract: string, tokens: Array<ShotToken>) => void {
      return (contract: string, tokens: Array<ShotToken>) => {
        let shotTokens = this.ShotTokens.ShotTokens.get(contract) as Array<ShotToken>
        if (!shotTokens) {
          shotTokens = []
        }
        tokens.forEach((token) => {
          const index = shotTokens.findIndex((el) => el.ID === token.ID)
          shotTokens?.splice(index >= 0 ? index : 0, index >= 0 ? 1 : 0, token)
        })
        this.ShotTokens.ShotTokens.set(contract, shotTokens)
      }
    },
    shotTokens(): (contract: string) => Array<ShotToken> {
      return (contract:string) => {
        return this.ShotTokens.ShotTokens.get(contract) as Array<ShotToken>
      }
    }
  },
  actions: {
    getContractAndTokens (req: GetContractAndTokensRequest, done: (error: boolean, rows: Contract, tokens: ShotToken[]) => void) {
      doActionWithError<GetContractAndTokensRequest, GetContractAndTokensResponse>(
        API.GET_CONTRACT_AND_TOKENS,
        req,
        req.Message,
        (resp: GetContractAndTokensResponse): void => {
          this.addShotTokens(req.Contract, resp.Tokens)
          this.Contract = resp.Contract
          done(false, resp.Contract, resp.Tokens)
        }, () => {
          done(true, {} as Contract, [])
      })
    },
  }
})