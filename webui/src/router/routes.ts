import { RouteRecordRaw } from 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    DisplayToolbarSearchBox: boolean
    NeedLogined?: boolean
  }
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue'), meta: { DisplayToolbarSearchBox: false } },
      { path: 'whitepaper', component: () => import('pages/Whitepaper.vue'), meta: { DisplayToolbarSearchBox: true } },
      { path: 'transaction', component: () => import('pages/Transaction.vue'), meta: { DisplayToolbarSearchBox: true, NeedLogined: true } },
      { path: 'deck', component: () => import('pages/Deck.vue'), meta: { DisplayToolbarSearchBox: true } }
    ],
  },
  {
    path: '/result',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('components/NFT/Result.vue'), 
        meta: {
          DisplayToolbarSearchBox: false,
          NeedLogined: false,
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
