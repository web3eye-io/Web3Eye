import { defineStore } from 'pinia'
import { Cookies } from 'quasar'

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
    resetStorageKey () {
        return () => Cookies.remove('Storage-Key')
    },
  },
  actions: {}
})
