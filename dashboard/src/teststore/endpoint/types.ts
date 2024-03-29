import { ChainType } from '../basetypes/const';
import { EndpointState } from '../basetypes/endpoint/const';
import { BaseRequest } from '../local';

export interface Endpoint {
    ID: number;
    EntID: string;
    ChainType: ChainType;
    ChainID: string;
    Address: string;
    State: EndpointState;
    Remark: string;
    RPS: number
}

export interface CreateEndpointRequest extends BaseRequest{
    ChainType: ChainType;
    ChainID: string;
    Address: string;
    State: EndpointState;
    Remark: string;
    RPS: number
}
  
export interface CreateEndpointResponse {
    Info: Endpoint;
}
  

export interface DeleteEndpointRequest extends BaseRequest{
    ID: number;
    EntID: string;
}
  
export interface DeleteEndpointResponse {
    Info: Endpoint;
}
  
export interface GetEndpointsRequest extends BaseRequest{
    Offset: number;
    Limit: number;
}
export interface GetEndpointsResponse {
    Infos: Array<Endpoint>
    Total: number
}

export interface UpdateEndpointRequest extends BaseRequest{
    ID: number;
    EntID: string;
    ChainType: ChainType;
    ChainID: string;
    Address: string;
    State: EndpointState;
    Remark: string;
    RPS?: number
}
  
export interface UpdateEndpointResponse {
    Info: Endpoint;
}