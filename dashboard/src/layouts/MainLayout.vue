<template>
  <q-layout view='hHh Lpr fFf'>
    <MainHeader />
    <MainDrawer />
    <q-page-container>
      <MainBreadcrumbs />
      <router-view />
    </q-page-container>
    <Footer />
  </q-layout>
</template>

<script setup lang='ts'>
import { defineAsyncComponent, } from 'vue'

const MainHeader = defineAsyncComponent(() => import('src/components/header/MainHeader.vue'))
const Footer = defineAsyncComponent(() => import('src/components/footer/Footer.vue'))
const MainBreadcrumbs = defineAsyncComponent(() => import('src/components/breadcrumbs/MainBreadcrumbs.vue'))
const MainDrawer = defineAsyncComponent(() => import('src/components/drawer/MainDrawer.vue'))
</script>

<style lang='sass' scoped>
.q-page-container
  ::v-deep .q-table__container
    padding-bottom: 60px
</style>
