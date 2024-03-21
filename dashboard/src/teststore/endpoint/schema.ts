/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

/** @default "ChainUnkonwn" */
export enum ChainChainType {
  ChainUnkonwn = 'ChainUnkonwn',
  Ethereum = 'Ethereum',
  Solana = 'Solana',
}

/** @default "EndpointDefault" */
export enum CttypeEndpointState {
  EndpointDefault = 'EndpointDefault',
  EndpointAvailable = 'EndpointAvailable',
  EndpointUnstable = 'EndpointUnstable',
  EndpointError = 'EndpointError',
}

export interface EndpointConds {
  ID?: Web3EyeStringVal;
  ChainType?: Web3EyeStringVal;
  ChainID?: Web3EyeStringVal;
  Address?: Web3EyeStringVal;
  State?: Web3EyeStringVal;
  Remark?: Web3EyeStringVal;
  IDs?: Web3EyeStringSliceVal;
}

export interface EndpointCountEndpointsRequest {
  Conds?: EndpointConds;
}

export interface EndpointCountEndpointsResponse {
  /** @format int64 */
  Info?: number;
}

export interface EndpointCreateEndpointRequest {
  Info?: EndpointEndpointReq;
}

export interface EndpointCreateEndpointResponse {
  Info?: EndpointEndpoint;
}

export interface EndpointCreateEndpointsRequest {
  Infos?: EndpointEndpointReq[];
}

export interface EndpointCreateEndpointsResponse {
  Infos?: EndpointEndpoint[];
}

export interface EndpointDeleteEndpointRequest {
  ID?: string;
}

export interface EndpointDeleteEndpointResponse {
  Info?: EndpointEndpoint;
}

export interface EndpointEndpoint {
  ID?: string;
  ChainType?: ChainChainType;
  ChainID?: string;
  Address?: string;
  State?: CttypeEndpointState;
  Remark?: string;
}

export interface EndpointEndpointReq {
  ID?: string;
  ChainType?: ChainChainType;
  ChainID?: string;
  Address?: string;
  State?: CttypeEndpointState;
  Remark?: string;
}

export interface EndpointExistEndpointCondsRequest {
  Conds?: EndpointConds;
}

export interface EndpointExistEndpointCondsResponse {
  Exist?: boolean;
}

export interface EndpointExistEndpointRequest {
  ID?: string;
}

export interface EndpointExistEndpointResponse {
  Exist?: boolean;
}

export interface EndpointFailedInfo {
  ID?: string;
  MSG?: string;
}

export interface EndpointGetEndpointOnlyRequest {
  Conds?: EndpointConds;
}

export interface EndpointGetEndpointOnlyResponse {
  Info?: EndpointEndpoint;
}

export interface EndpointGetEndpointRequest {
  ID?: string;
}

export interface EndpointGetEndpointResponse {
  Info?: EndpointEndpoint;
}

export interface EndpointGetEndpointsRequest {
  Conds?: EndpointConds;
  /** @format int32 */
  Offset?: number;
  /** @format int32 */
  Limit?: number;
}

export interface EndpointGetEndpointsResponse {
  Infos?: EndpointEndpoint[];
  /** @format int64 */
  Total?: number;
}

export interface EndpointUpdateEndpointRequest {
  Info?: EndpointEndpointReq;
}

export interface EndpointUpdateEndpointResponse {
  Info?: EndpointEndpoint;
}

export interface EndpointUpdateEndpointsRequest {
  Infos?: EndpointEndpointReq[];
}

export interface EndpointUpdateEndpointsResponse {
  Infos?: EndpointFailedInfo[];
}

export interface ProtobufAny {
  '@type'?: string;
  [key: string]: any;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface Web3EyeStringSliceVal {
  Op?: string;
  Value?: string[];
}

export interface Web3EyeStringVal {
  Op?: string;
  Value?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, 'body' | 'bodyUsed'>;

export interface FullRequestParams extends Omit<RequestInit, 'body'> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, 'body' | 'method' | 'query' | 'path'>;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, 'baseUrl' | 'cancelToken' | 'signal'>;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = 'application/json',
  FormData = 'multipart/form-data',
  UrlEncoded = 'application/x-www-form-urlencoded',
  Text = 'text/plain',
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = '';
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>['securityWorker'];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: 'same-origin',
    headers: {},
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === 'number' ? value : `${value}`)}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join('&');
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => 'undefined' !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join('&');
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : '';
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === 'object' || typeof input === 'string') ? JSON.stringify(input) : input,
    [ContentType.Text]: (input: any) => (input !== null && typeof input !== 'string' ? JSON.stringify(input) : input),
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === 'object' && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === 'boolean' ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ''}${path}${queryString ? `?${queryString}` : ''}`, {
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type && type !== ContentType.FormData ? { 'Content-Type': type } : {}),
      },
      signal: (cancelToken ? this.createAbortSignal(cancelToken) : requestParams.signal) || null,
      body: typeof body === 'undefined' || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title web3eye/entrance/v1/endpoint/endpoint.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  v1 = {
    /**
     * No description
     *
     * @tags Manager
     * @name ManagerCountEndpoints
     * @request POST:/v1/count/endpoints
     */
    managerCountEndpoints: (body: EndpointCountEndpointsRequest, params: RequestParams = {}) =>
      this.request<EndpointCountEndpointsResponse, RpcStatus>({
        path: `/v1/count/endpoints`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerCreateEndpoint
     * @request POST:/v1/create/endpoint
     */
    managerCreateEndpoint: (body: EndpointCreateEndpointRequest, params: RequestParams = {}) =>
      this.request<EndpointCreateEndpointResponse, RpcStatus>({
        path: `/v1/create/endpoint`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerCreateEndpoints
     * @request POST:/v1/create/endpoints
     */
    managerCreateEndpoints: (body: EndpointCreateEndpointsRequest, params: RequestParams = {}) =>
      this.request<EndpointCreateEndpointsResponse, RpcStatus>({
        path: `/v1/create/endpoints`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerDeleteEndpoint
     * @request POST:/v1/delete/endpoint
     */
    managerDeleteEndpoint: (body: EndpointDeleteEndpointRequest, params: RequestParams = {}) =>
      this.request<EndpointDeleteEndpointResponse, RpcStatus>({
        path: `/v1/delete/endpoint`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerExistEndpoint
     * @request POST:/v1/exist/endpoint
     */
    managerExistEndpoint: (body: EndpointExistEndpointRequest, params: RequestParams = {}) =>
      this.request<EndpointExistEndpointResponse, RpcStatus>({
        path: `/v1/exist/endpoint`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerExistEndpointConds
     * @request POST:/v1/exist/endpoint/conds
     */
    managerExistEndpointConds: (body: EndpointExistEndpointCondsRequest, params: RequestParams = {}) =>
      this.request<EndpointExistEndpointCondsResponse, RpcStatus>({
        path: `/v1/exist/endpoint/conds`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerGetEndpoint
     * @request POST:/v1/get/endpoint
     */
    managerGetEndpoint: (body: EndpointGetEndpointRequest, params: RequestParams = {}) =>
      this.request<EndpointGetEndpointResponse, RpcStatus>({
        path: `/v1/get/endpoint`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerGetEndpointOnly
     * @request POST:/v1/get/endpoint/only
     */
    managerGetEndpointOnly: (body: EndpointGetEndpointOnlyRequest, params: RequestParams = {}) =>
      this.request<EndpointGetEndpointOnlyResponse, RpcStatus>({
        path: `/v1/get/endpoint/only`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerGetEndpoints
     * @request POST:/v1/get/endpoints
     */
    managerGetEndpoints: (body: EndpointGetEndpointsRequest, params: RequestParams = {}) =>
      this.request<EndpointGetEndpointsResponse, RpcStatus>({
        path: `/v1/get/endpoints`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerUpdateEndpoint
     * @request POST:/v1/update/endpoint
     */
    managerUpdateEndpoint: (body: EndpointUpdateEndpointRequest, params: RequestParams = {}) =>
      this.request<EndpointUpdateEndpointResponse, RpcStatus>({
        path: `/v1/update/endpoint`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),

    /**
     * No description
     *
     * @tags Manager
     * @name ManagerUpdateEndpoints
     * @request POST:/v1/update/endpoints
     */
    managerUpdateEndpoints: (body: EndpointUpdateEndpointsRequest, params: RequestParams = {}) =>
      this.request<EndpointUpdateEndpointsResponse, RpcStatus>({
        path: `/v1/update/endpoints`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),
  };
}
