import { defineStore } from 'pinia'
import Web3 from 'web3'
import { Account } from './types'

export const useWeb3jsStore = defineStore('local-web3js', {
  state: () => ({
    Account: {} as Account,
    Web3: {} as Web3
  }),
  getters: {
    getAccount () {
      return () =>  this.Account
    },
    setAccount () {
      return (account: Account) => this.Account = account
    },
    getWeb3 () {
      return () => this.Web3
    },
    setWeb3 () {
      return (web3: Web3) => this.Web3 = web3
    }
  },
  actions: {
    // TODO
  }
})
