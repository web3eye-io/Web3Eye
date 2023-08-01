import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTokensRequest, GetTokensResponse, SearchToken } from './types' 

export const useTokenStore = defineStore('token', {
  state: () => ({
    SearchTokens: {
      SearchTokens: [] as Array<SearchToken>,
      Total: 0,
      Current: ''
    }
  }),
  getters: {
    setToken () {
      return (rows: Array<SearchToken>) => {
        this.SearchTokens.SearchTokens = rows
      }
    }
  },
  actions: {
    getTokens (req: GetTokensRequest, done: (error: boolean, rows: SearchToken[]) => void) {
      doActionWithError<GetTokensRequest, GetTokensResponse>(
        API.SEARCH_PAGE,
        req,
        req.Message,
        (resp: GetTokensResponse): void => {
          this.SearchTokens.SearchTokens = resp.Infos
          this.SearchTokens.Total = resp.TotalTokens
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
  }
})