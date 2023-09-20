<template>
  <q-drawer v-model='leftDrawerOpen' :mini='leftDrawerMini' show-if-above bordered>
    <div class='row'>
      <q-space />
      <q-btn
        flat dense round icon='swap_horiz'
        aria-label='Menu'
        class='drawer-toggle'
        @click='toggleLeftDrawer'
      />
    </div>
    <q-list separator>
      <DrawerMenu
        v-for='item in MainDrawerMenus'
        :key='item.label'
        :menu='item'
        @click='onItemClick(item)'
        :mini='leftDrawerMini'
      />
    </q-list>
  </q-drawer>
</template>
<script setup lang='ts'>
import { ref, defineAsyncComponent } from 'vue'
import { MenuItem, useMenuStore, HomePageBreadcrumbs } from 'src/localstore'
import { MainDrawerMenus } from 'src/menus/menus'
const DrawerMenu = defineAsyncComponent(() => import('src/components/drawer/DrawerMenu.vue'))

const leftDrawerOpen = ref(true)
const leftDrawerMini = ref(false)

const menus = useMenuStore()
const onItemClick = (menu: MenuItem) => {
  menus.MainBreadcrumbs = [HomePageBreadcrumbs, menu] as Array<MenuItem>
}
const toggleLeftDrawer = (): void => {
  leftDrawerMini.value = !leftDrawerMini.value
}
</script>
<style lang='sass' scoped>
.drawer-toggle
  margin-left: auto
  margin-right: auto
</style>
