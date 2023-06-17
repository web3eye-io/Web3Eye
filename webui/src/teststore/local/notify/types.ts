export enum NotifyType {
  Error = 'error',
  Info = 'info',
  Warning = 'warning',
  Waiting = 'waiting',
  Success = 'success'
}

export interface Notification {
  Title?: string
  Message?: string
  Description?: string
  Popup?: boolean
  Type?: NotifyType
}

export interface ReqMessage {
  Info?: Notification
  Error?: Notification
}


export interface BaseRequest {
  Message: ReqMessage
}

export interface MyRequest {
  NotifyMessage: ReqMessage
}

export interface NotificationState {
  Notifications: Array<Notification>
}