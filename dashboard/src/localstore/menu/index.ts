import { defineStore } from 'pinia'
import { HomePageBreadcrumbs } from './const'
import { MenuState } from './types'

export const useMenuStore = defineStore('menu', {
  state: (): MenuState => ({
    MainBreadcrumbs: [HomePageBreadcrumbs],
    ActiveMainBreadcrumb: HomePageBreadcrumbs
  }),
  getters: {},
  actions: {}
})

export * from './types'
export * from './const'
