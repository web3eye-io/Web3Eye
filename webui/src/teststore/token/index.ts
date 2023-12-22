import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { GetTokenRequest, GetTokenResponse, GetTokensRequest, GetTokensResponse, SearchToken, SearchTokenMessage, SearchTokensResponse, Token } from './types' 
import { Cookies } from 'quasar'

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
          this.SearchTokens.TotalPages = resp.TotalPages
          this.SearchTokens.TotalTokens = resp.TotalTokens
          this.SearchTokens.StorageKey = resp.StorageKey
          if (resp.StorageKey?.length > 0) {
            Cookies.set('Storage-Key', resp.StorageKey, { expires: '4h', secure: true, path: '/' })
          }
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
    getTokens (req: GetTokensRequest, done: (error: boolean, rows: SearchToken[]) => void) {
      const key = Cookies.get('Storage-Key')
      if (key && key?.length > 0) {
        req.StorageKey = key
      }
      doActionWithError<GetTokensRequest, GetTokensResponse>(
        API.SEARCH_PAGE,
        req,
        req.Message,
        (resp: GetTokensResponse): void => {
          this.addSearchTokens(resp.Infos)
          this.SearchTokens.StorageKey = resp.StorageKey
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
          this.Token.Token.set(resp.Info.ID, resp.Info)
          done(false, resp.Info)
        }, () => {
          done(true, {} as Token)
      })
    },
  }
})