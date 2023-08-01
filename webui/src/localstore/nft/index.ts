import { defineStore } from 'pinia'
import { NFTMeta } from './types'

export const useNFTMetaStore = defineStore('local-token', {
  state: () => ({
    NTFMetas: {
      NTFMetas: [] as Array<NFTMeta>,
      Current: ''
    }
  }),
  getters: {
    setNftMeta () {
      return (rows: Array<NFTMeta>) => {
        this.NTFMetas.NTFMetas = rows
      }
    }
  },
  actions: {}
})
