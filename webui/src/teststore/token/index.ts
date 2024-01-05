import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTokenRequest, GetTokenResponse, GetTokensRequest, GetTokensResponse, SearchToken, SearchTokenMessage, SearchTokensResponse, Token } from './types' 
import { useStorageKeyStore } from 'src/localstore/storagekey'

export const useTokenStore = defineStore('token', {
  state: () => ({
    SearchTokens: {
      SearchTokens: [] as Array<SearchToken>,
      Total: 0,
      Current: '',
      StorageKey: '',
      Pages: 0
    },
    Token: {
      Token: new Map<number, Token>(),
    }
  }),

  getters: {
    setSearchToken () {
      return (rows: Array<SearchToken>) => {
        this.SearchTokens.SearchTokens = rows
      }
    },
    getTokenByID () {
      return (id: number) => {
        return this.Token.Token.get(id)
      }
    },
    addSearchTokens (): (tokens: Array<SearchToken>) => void {
      return (tokens: Array<SearchToken>) => {
        tokens.forEach((token) => {
          const index = this.SearchTokens.SearchTokens.findIndex((el) => el.ID === token.ID)
          this.SearchTokens.SearchTokens.splice(index >= 0 ? index : 0, index >= 0 ? 1 : 0, token)
        })
      }
    },
  },
  actions: {
    searchTokens (req: FormData, reqMessage: SearchTokenMessage, done: (error: boolean, rows?: SearchToken[]) => void) {
      doActionWithError<object, SearchTokensResponse>(
        API.SEARCH_FILE,
        req,
        reqMessage.Message,
        (resp: SearchTokensResponse): void => {
          this.addSearchTokens(resp.Infos)
          this.SearchTokens.Pages = resp.Pages
          this.SearchTokens.Total = resp.Total
          this.SearchTokens.StorageKey = resp.StorageKey
          const localkey = useStorageKeyStore()
          if (resp.StorageKey?.length > 0) {
            localkey.setStorageKey(resp.StorageKey)
          }
          if (resp.Vector?.length > 0) {
            localkey.setVector(resp.Vector)
          }
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
    getTokens (req: GetTokensRequest, done: (error: boolean, rows: SearchToken[], totalPages?: number) => void) {
      const localkey = useStorageKeyStore()
      const key = localkey.getStorageKey()
      if (key === null) {
        return
      }
      if (key && key?.length > 0) {
        req.StorageKey = key
      }
      const vector = localkey.getVector()
      if (vector == null || vector === undefined) {
        return
      }
      req.Vector = vector as unknown as Array<number>
      doActionWithError<GetTokensRequest, GetTokensResponse>(
        API.SEARCH_PAGE,
        req,
        req.Message,
        (resp: GetTokensResponse): void => {
          this.addSearchTokens(resp.Infos)
          this.SearchTokens.StorageKey = resp.StorageKey
          this.SearchTokens.Pages = resp.Pages
          done(false, resp.Infos, resp.Pages)
        }, () => {
          done(true, [], 0)
      })
    },
    getToken (req: GetTokenRequest, done: (error: boolean, row: Token) => void) {
      doActionWithError<GetTokenRequest, GetTokenResponse>(
        API.GET_TOKEN,
        req,
        req.Message,
        (resp: GetTokenResponse): void => {
          this.Token.Token.set(resp.Info.ID, resp.Info)
          done(false, resp.Info)
        }, () => {
          done(true, {} as Token)
      })
    },
  }
})