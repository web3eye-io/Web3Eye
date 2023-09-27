import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTokenRequest, GetTokenResponse, GetTokensRequest, GetTokensResponse, SearchToken, SearchTokensRequest, SearchTokensResponse, Token } from './types' 

export const useTokenStore = defineStore('token', {
  state: () => ({
    SearchTokens: {
      SearchTokens: [] as Array<SearchToken>,
      TotalTokens: 0,
      Current: '',
      StorageKey: '',
      TotalPages: 0
    },
    Token: {
      Token: new Map<string, Token>(),
    }
  }),

  getters: {
    setSearchToken () {
      return (rows: Array<SearchToken>) => {
        this.SearchTokens.SearchTokens = rows
      }
    },
    getTokenByID () {
      return (tokenID: string) => {
        return this.Token.Token.get(tokenID)
      }
    }
  },
  actions: {
    searchTokens (req: FormData, done: (error: boolean, rows?: SearchToken[]) => void) {
        const xhttp = new XMLHttpRequest();
        xhttp.open('POST', '/api/entrance/search/file');
        xhttp.send(req)
        xhttp.onload = () => {
          const response = JSON.parse(xhttp.responseText) as SearchTokensResponse
          this.SearchTokens.SearchTokens = response.Infos
          this.SearchTokens.TotalPages = response.TotalPages
          this.SearchTokens.TotalTokens = response.TotalTokens
          this.SearchTokens.StorageKey = response.StorageKey
        }
        if (xhttp.status === 0) {
          done(false)
          return
        }
        done(true)
    },
    getTokens (req: GetTokensRequest, done: (error: boolean, rows: SearchToken[]) => void) {
      doActionWithError<GetTokensRequest, GetTokensResponse>(
        API.SEARCH_PAGE,
        req,
        req.Message,
        (resp: GetTokensResponse): void => {
          this.SearchTokens.SearchTokens.push(...resp.Infos)
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
    getToken (req: GetTokenRequest, done: (error: boolean, row: Token) => void) {
      doActionWithError<GetTokenRequest, GetTokenResponse>(
        API.GET_TOKEN,
        req,
        req.Message,
        (resp: GetTokenResponse): void => {
          const tokenID = resp.Info.TokenID
          this.Token.Token.set(tokenID, resp.Info)
          done(false, resp.Info)
        }, () => {
          done(true, {} as Token)
      })
    },
  }
})