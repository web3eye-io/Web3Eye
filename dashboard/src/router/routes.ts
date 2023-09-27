import { RouteRecordRaw } from 'vue-router'

declare module 'vue-router' {
  // eslint-disable-next-line @typescript-eslint/no-empty-interface
  interface RouteMeta {}
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('pages/Index.vue'), 
        meta: { 
        } 
      },
      { 
        path: 'manager/task', 
        component: () => import('pages/SyncTask.vue'), 
        meta: { 
        } 
      },
      { 
        path: 'manager/endpoint', 
        component: () => import('pages/Endpoint.vue'), 
        meta: { 
        } 
      },
      { 
        path: 'manager/overview', 
        component: () => import('pages/Overview.vue'), 
        meta: { 
        } 
      }
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue'),
  },
];

export default routes
