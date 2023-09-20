<template>
  <q-header>
    <q-toolbar class='toolbar'>
      <!-- <Logo class='logo' /> -->
      <q-space />
      <!-- <HeaderTools /> -->
    </q-toolbar>
  </q-header>
</template>

<script setup lang='ts'>
// import { defineAsyncComponent } from 'vue'

// const Logo = defineAsyncComponent(() => import('src/components/logo/Logo.vue'))
// const HeaderTools = defineAsyncComponent(() => import('src/components/header/HeaderTools.vue'))

</script>

<style lang='sass' scoped>
.logo
  max-height: 100%
  width: 64px

.toolbar
  background-color: $blue-4
  height: 48px
</style>
