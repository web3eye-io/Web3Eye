import { defineStore } from 'pinia'
import { doActionWithError } from '../action'
import { API } from './const'
import { 
  GetEndpointsRequest, 
  GetEndpointsResponse, 
  Endpoint, 
  CreateEndpointRequest, 
  CreateEndpointResponse, 
  UpdateEndpointRequest, 
  UpdateEndpointResponse, 
  DeleteEndpointRequest, 
  DeleteEndpointResponse 
} from './types'

export const useEndpointStore = defineStore('endpoint', {
  state: () => ({
    Endpoints: {
      Endpoints: [] as Array<Endpoint>,
      Total: 0,
    }
  }),
  getters: {},
  actions: {
    getEndpoints (req: GetEndpointsRequest, done: (error: boolean, rows: Endpoint[]) => void) {
      doActionWithError<GetEndpointsRequest, GetEndpointsResponse>(
        API.GET_ENDPOINTS,
        req,
        req.Message,
        (resp: GetEndpointsResponse): void => {
          this.Endpoints.Endpoints = resp.Infos
          this.Endpoints.Total = resp.Total
          done(false, resp.Infos)
        }, () => {
          done(true, [])
      })
    },
    createEndpoint (req: CreateEndpointRequest, done: (error: boolean, row?: Endpoint) => void) {
      doActionWithError<CreateEndpointRequest, CreateEndpointResponse>(
        API.CREATE_ENDPOINT,
        req,
        req.Message,
        (resp: CreateEndpointResponse): void => {
          this.Endpoints.Endpoints.push(resp.Info)
          this.Endpoints.Total += 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as Endpoint)
      })
    },
    updateEndpoint (req: UpdateEndpointRequest, done: (error: boolean, row?: Endpoint) => void) {
      doActionWithError<UpdateEndpointRequest, UpdateEndpointResponse>(
        API.UPDATE_ENDPOINT,
        req,
        req.Message,
        (resp: UpdateEndpointResponse): void => {
          const index = this.Endpoints.Endpoints.findIndex((el) => el.ID === resp.Info.ID && el.EntID === resp.Info.EntID)
          this.Endpoints.Endpoints.splice(index > 0 ? index : 0, index > 0 ? 1 : 0, resp.Info)
          this.Endpoints.Total += 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as Endpoint)
      })
    },
    deleteEndpoint (req: DeleteEndpointRequest, done: (error: boolean, row?: Endpoint) => void) {
      doActionWithError<DeleteEndpointRequest, DeleteEndpointResponse>(
        API.DELETE_ENDPOINT,
        req,
        req.Message,
        (resp: DeleteEndpointResponse): void => {
          const index = this.Endpoints.Endpoints.findIndex((el) => el.ID === resp.Info.ID)
          if(index < 0) return
          this.Endpoints.Endpoints.splice(index, 1)
          this.Endpoints.Total -= 1
          done(false, resp.Info)
        }, () => {
          done(true, {} as Endpoint)
      })
    },
  }
})