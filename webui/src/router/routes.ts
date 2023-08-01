import { RouteRecordRaw } from 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    DisplayToolbarSearchBox: boolean
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue'), meta: { DisplayToolbarSearchBox: false } },
      { path: 'whitepaper', component: () => import('pages/Whitepaper.vue'), meta: { DisplayToolbarSearchBox: true } },
      { path: 'transaction', component: () => import('pages/Transaction.vue'), meta: { DisplayToolbarSearchBox: true } },
      { path: 'deck', component: () => import('pages/Deck.vue'), meta: { DisplayToolbarSearchBox: true } }
    ],
  },
  {
    path: '/token',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('components/Token/Token.vue'), 
        meta: {
          DisplayToolbarSearchBox: false 
        } 
      },
      // { path: 'whitepaper', component: () => import('pages/Whitepaper.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'transaction', component: () => import('pages/Transaction.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'deck', component: () => import('pages/Deck.vue'), meta: { DisplayToolbarSearchBox: true } }
    ],
  },
  {
    path: '/transfer',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('components/Transfer/Transfer.vue'), 
        meta: {
          DisplayToolbarSearchBox: false 
        } 
      },
      // { path: 'whitepaper', component: () => import('pages/Whitepaper.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'transaction', component: () => import('pages/Transaction.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'deck', component: () => import('pages/Deck.vue'), meta: { DisplayToolbarSearchBox: true } }
    ],
  },
  {
    path: '/contract',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('components/Contract/Contract.vue'), 
        meta: {
          DisplayToolbarSearchBox: false 
        } 
      },
      // { path: 'whitepaper', component: () => import('pages/Whitepaper.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'transaction', component: () => import('pages/Transaction.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'deck', component: () => import('pages/Deck.vue'), meta: { DisplayToolbarSearchBox: true } }
    ],
  },
  {
    path: '/backup',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('components/Backup/Backup.vue'), 
        meta: {
          DisplayToolbarSearchBox: false 
        } 
      },
      // { path: 'whitepaper', component: () => import('pages/Whitepaper.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'transaction', component: () => import('pages/Transaction.vue'), meta: { DisplayToolbarSearchBox: true } },
      // { path: 'deck', component: () => import('pages/Deck.vue'), meta: { DisplayToolbarSearchBox: true } }
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
