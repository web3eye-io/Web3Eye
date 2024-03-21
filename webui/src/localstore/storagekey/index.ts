import { defineStore } from 'pinia'
import { Cookies, LocalStorage } from 'quasar'


export const useStorageKeyStore = defineStore('local-storage-key', {
  state: () => ({}),
  getters: {
    setStorageKey () {
      return (key: string) => {
        if (key && key?.length > 0) {
            Cookies.set('Storage-Key', key,  { expires: '4h', secure: true, path: '/' })
        }
      }
    },
    getStorageKey () {
        return () => Cookies.get('Storage-Key')
    },
    setVector () {
      return (vectors: Array<number>) => {
        if (vectors?.length > 0) {
            LocalStorage.set('Vector', vectors)
        }
      }
    },
    getVector () {
        return () => {
          const vector = LocalStorage.getItem('Vector')
          if(vector) {
            return vector
          }
        }
    },
    reset () {
      return () => {
        Cookies.remove('Storage-Key')
        Cookies.remove('Vector')
      }
  },
  },
  actions: {}
})
