import {defineStore } from 'pinia'
import Web3 from 'web3'
import { Account } from './types'
import { Cookies } from 'quasar'

export const useWeb3jsStore = defineStore('local-web3js', {
  state: () => ({
    Account: {} as Account,
    Web3: {} as Web3
  }),
  getters: {
    getAccount () {
      return () =>  {
        if (this.Account.Address?.length > 0) {
          return this.Account
        }else {
          return this.getAccountFromCookie()
        }
      }
    },
    setAccount () {
      return (account: Account) => {
        Cookies.set('X-WEB3-ADDRESS', account.Address)
        Cookies.set('X-WEB3-BALANCE', account.Balance)
        Cookies.set('X-WEB3-CHAIN_ID', `${account.ChainID}`)
        Cookies.set('X-WEB3-NETWORK', account.Network)
        this.Account = account
      }
    },
    getAccountFromCookie () {
      return () => {
        const account = {} as Account
        account.Address = Cookies.get('X-WEB3-ADDRESS')
        account.Balance = Cookies.get('X-WEB3-BALANCE')
        account.ChainID = Cookies.get('X-WEB3-CHAIN_ID')
        account.Network = Cookies.get('X-WEB3-NETWORK')
        return account
      }
    },
    getWeb3 () {
      return () => this.Web3
    },
    setWeb3 () {
      return (web3: Web3) => this.Web3 = web3
    },
    logined () {
      // TODO
    },
    logout () {
      Cookies.remove('X-WEB3-ADDRESS')
      Cookies.remove('X-WEB3-BALANCE')
      Cookies.remove('X-WEB3-CHAIN_ID')
      Cookies.remove('X-WEB3-NETWORK')
      this.Web3 = {} as Web3
    }
  },
  actions: {
    //
  }
})
