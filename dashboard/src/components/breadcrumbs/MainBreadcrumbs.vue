<template>
  <Breadcrumbs
    class='main'
    :breadcrumbs='breadcrumbs'
  />
</template>

<script setup lang='ts'>
import { useMenuStore } from 'src/localstore'
import { defineAsyncComponent, computed } from 'vue'

const Breadcrumbs = defineAsyncComponent(() => import('src/components/breadcrumbs/Breadcrumbs.vue'))

const menu = useMenuStore()
const breadcrumbs = computed(() => menu.MainBreadcrumbs)

</script>

<style lang='sass' scoped>
.main
  border-bottom: 1px solid $grey-4
  font-size: 0.875em
  line-height: 1em
  padding: 3px 0
</style>
