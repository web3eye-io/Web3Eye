import { defineStore } from 'pinia'
import { NotificationState, Notification } from './types'

export const useNotificationStore = defineStore('local-notification-v3', {
  state: (): NotificationState => ({
    Notifications: []
  }),
  getters: {},
  actions: {
    pushNotification (notification: Notification) {
      this.Notifications.push(notification)
    },
    popNotification (): Notification | undefined {
      if (this.Notifications.length > 0) {
        const notification = this.Notifications[0]
        this.Notifications = this.Notifications.splice(0, 1)
        return notification
      }
      return undefined
    }
  }
})

export * from './helper'
export * from './types'
