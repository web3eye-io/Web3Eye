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

/** @default "Default" */
export enum CttypeSyncState {
  Default = 'Default',
  Start = 'Start',
  Pause = 'Pause',
  Finish = 'Finish',
  Failed = 'Failed',
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

export interface SynctaskConds {
  ID?: Web3EyeStringVal;
  ChainType?: Web3EyeStringVal;
  ChainID?: Web3EyeStringVal;
  Start?: Web3EyeUint64Val;
  End?: Web3EyeUint64Val;
  Current?: Web3EyeUint64Val;
  Topic?: Web3EyeStringVal;
  Description?: Web3EyeStringVal;
  SyncState?: Web3EyeStringVal;
  Remark?: Web3EyeStringVal;
  IDs?: Web3EyeStringSliceVal;
}

export interface SynctaskCountSyncTasksRequest {
  Conds?: SynctaskConds;
}

export interface SynctaskCountSyncTasksResponse {
  /** @format int64 */
  Info?: number;
}

export interface SynctaskCreateSyncTaskRequest {
  Info?: SynctaskSyncTaskReq;
}

export interface SynctaskCreateSyncTaskResponse {
  Info?: SynctaskSyncTask;
}

export interface SynctaskDeleteSyncTaskRequest {
  ID?: string;
}

export interface SynctaskDeleteSyncTaskResponse {
  Info?: SynctaskSyncTask;
}

export interface SynctaskExistSyncTaskCondsRequest {
  Conds?: SynctaskConds;
}

export interface SynctaskExistSyncTaskCondsResponse {
  Exist?: boolean;
}

export interface SynctaskExistSyncTaskRequest {
  ID?: string;
}

export interface SynctaskExistSyncTaskResponse {
  Exist?: boolean;
}

export interface SynctaskGetSyncTaskOnlyRequest {
  Conds?: SynctaskConds;
}

export interface SynctaskGetSyncTaskOnlyResponse {
  Info?: SynctaskSyncTask;
}

export interface SynctaskGetSyncTaskRequest {
  ID?: string;
}

export interface SynctaskGetSyncTaskResponse {
  Info?: SynctaskSyncTask;
}

export interface SynctaskGetSyncTasksRequest {
  Conds?: SynctaskConds;
  /** @format int32 */
  Offset?: number;
  /** @format int32 */
  Limit?: number;
}

export interface SynctaskGetSyncTasksResponse {
  Infos?: SynctaskSyncTask[];
  /** @format int64 */
  Total?: number;
}

export interface SynctaskSyncTask {
  ID?: string;
  ChainType?: ChainChainType;
  ChainID?: string;
  /** @format uint64 */
  Start?: string;
  /** @format uint64 */
  End?: string;
  /** @format uint64 */
  Current?: string;
  Topic?: string;
  Description?: string;
  SyncState?: CttypeSyncState;
  Remark?: string;
}

export interface SynctaskSyncTaskReq {
  ID?: string;
  ChainType?: ChainChainType;
  ChainID?: string;
  /** @format uint64 */
  Start?: string;
  /** @format uint64 */
  End?: string;
  /** @format uint64 */
  Current?: string;
  Topic?: string;
  Description?: string;
  SyncState?: CttypeSyncState;
  Remark?: string;
}

export interface SynctaskUpdateSyncTaskRequest {
  Info?: SynctaskSyncTaskReq;
}

export interface SynctaskUpdateSyncTaskResponse {
  Info?: SynctaskSyncTask;
}

export interface Web3EyeStringSliceVal {
  Op?: string;
  Value?: string[];
}

export interface Web3EyeStringVal {
  Op?: string;
  Value?: string;
}

export interface Web3EyeUint64Val {
  Op?: string;
  /** @format uint64 */
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
 * @title web3eye/entrance/v1/synctask/synctask.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  v1 = {
    /**
     * No description
     *
     * @tags Manager
     * @name ManagerCountSyncTasks
     * @request POST:/v1/count/synctasks
     */
    managerCountSyncTasks: (body: SynctaskCountSyncTasksRequest, params: RequestParams = {}) =>
      this.request<SynctaskCountSyncTasksResponse, RpcStatus>({
        path: `/v1/count/synctasks`,
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
     * @name ManagerCreateSyncTask
     * @request POST:/v1/create/synctask
     */
    managerCreateSyncTask: (body: SynctaskCreateSyncTaskRequest, params: RequestParams = {}) =>
      this.request<SynctaskCreateSyncTaskResponse, RpcStatus>({
        path: `/v1/create/synctask`,
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
     * @name ManagerDeleteSyncTask
     * @request POST:/v1/delete/synctask
     */
    managerDeleteSyncTask: (body: SynctaskDeleteSyncTaskRequest, params: RequestParams = {}) =>
      this.request<SynctaskDeleteSyncTaskResponse, RpcStatus>({
        path: `/v1/delete/synctask`,
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
     * @name ManagerExistSyncTask
     * @request POST:/v1/exist/synctask
     */
    managerExistSyncTask: (body: SynctaskExistSyncTaskRequest, params: RequestParams = {}) =>
      this.request<SynctaskExistSyncTaskResponse, RpcStatus>({
        path: `/v1/exist/synctask`,
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
     * @name ManagerExistSyncTaskConds
     * @request POST:/v1/exist/synctask/conds
     */
    managerExistSyncTaskConds: (body: SynctaskExistSyncTaskCondsRequest, params: RequestParams = {}) =>
      this.request<SynctaskExistSyncTaskCondsResponse, RpcStatus>({
        path: `/v1/exist/synctask/conds`,
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
     * @name ManagerGetSyncTask
     * @request POST:/v1/get/synctask
     */
    managerGetSyncTask: (body: SynctaskGetSyncTaskRequest, params: RequestParams = {}) =>
      this.request<SynctaskGetSyncTaskResponse, RpcStatus>({
        path: `/v1/get/synctask`,
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
     * @name ManagerGetSyncTaskOnly
     * @request POST:/v1/get/synctask/only
     */
    managerGetSyncTaskOnly: (body: SynctaskGetSyncTaskOnlyRequest, params: RequestParams = {}) =>
      this.request<SynctaskGetSyncTaskOnlyResponse, RpcStatus>({
        path: `/v1/get/synctask/only`,
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
     * @name ManagerGetSyncTasks
     * @request POST:/v1/get/synctasks
     */
    managerGetSyncTasks: (body: SynctaskGetSyncTasksRequest, params: RequestParams = {}) =>
      this.request<SynctaskGetSyncTasksResponse, RpcStatus>({
        path: `/v1/get/synctasks`,
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
     * @name ManagerUpdateSyncTask
     * @request POST:/v1/update/synctask
     */
    managerUpdateSyncTask: (body: SynctaskUpdateSyncTaskRequest, params: RequestParams = {}) =>
      this.request<SynctaskUpdateSyncTaskResponse, RpcStatus>({
        path: `/v1/update/synctask`,
        method: 'POST',
        body: body,
        type: ContentType.Json,
        format: 'json',
        ...params,
      }),
  };
}
