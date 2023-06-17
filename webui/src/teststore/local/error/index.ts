import { defineStore } from 'pinia'
import { SwitchTarget } from './const'
import { ErrorSwitcherState, ErrorTarget } from './types'

export const useErrorStore = defineStore('local-errorswitcher-v3', {
  state: (): ErrorSwitcherState => ({
    ErrorTargets: [
      {
        ErrorCode: 403,
        Target: SwitchTarget.LOGIN
      }
    ],
    ErrorTrigger: undefined as unknown as ErrorTarget
  }),
  getters: {},
  actions: {}
})

export * from './const'
export * from './types'
