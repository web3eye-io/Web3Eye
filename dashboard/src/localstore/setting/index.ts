import { defineStore } from 'pinia'
import { SettingState } from './types'

export const useLocalSettingStore = defineStore('local-setting', {
  state: (): SettingState => ({
    DisplayToolbarSearchBox: false
  }),
  getters: {},
  actions: {}
})
