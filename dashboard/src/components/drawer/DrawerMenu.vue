<template>
  <q-expansion-item
    v-if='children.length > 0 && !mini'
    :key='menu.label'
    :content-inset-level='insetLevel'
    class='item'
    :clickable='logined'
    @click='onItemClick'
  >
    <template #header>
      <div class='row container'>
        <q-icon size='2em' class='icon' v-if='menu.icon' :name='menu.icon' />
        <q-item-section>
          <q-item-label>{{ menu.label }}</q-item-label>
          <q-item-label v-if='menu.caption !== ""' caption>
            {{ menu.caption }}
          </q-item-label>
        </q-item-section>
        <q-space />
      </div>
    </template>
    <DrawerMenu
      v-for='item in children'
      :key='item.label'
      :menu='item'
      :mini='mini'
    />
  </q-expansion-item>
  <q-item
    v-else
    :clickable='logined'
    :target='menu.target'
    class='item'
    active-class='active'
    :active='active'
    @click='onItemClick'
  >
    <div class='row'>
      <q-icon size='2em' class='icon' v-if='menu.icon' :name='menu.icon' />
      <q-item-section>
        <q-item-label>{{ menu.label }}</q-item-label>
        <q-item-label v-if='menu.caption !== ""' caption>
          {{ menu.caption }}
        </q-item-label>
      </q-item-section>
    </div>
  </q-item>
</template>

<script setup lang='ts'>
import { MenuItem, useMenuStore } from 'src/localstore'
import { defineProps, toRef, computed } from 'vue'
import { useRouter } from 'vue-router'

interface Props {
  menu: MenuItem
  mini: boolean
}
const props = defineProps<Props>()
const menu = toRef(props, 'menu')
const mini = toRef(props, 'mini')

const insetLevel = computed(() => {
  return menu.value.level === undefined ? 0 : menu.value.level * 0.1 + 0.3
})
const children = computed(() => {
  return menu.value.children === undefined ? [] : menu.value.children
})

const router = useRouter()
const menus = useMenuStore()
const active = computed(() => menus.ActiveMainBreadcrumb?.menuId === menu.value.menuId)

const logined = computed(() => true)

const onItemClick = () => {
  menus.ActiveMainBreadcrumb = menu.value
  menus.MainBreadcrumbs = menus.MainBreadcrumbs.filter((_, index) => index <= menu.value.level)

  menus.MainBreadcrumbs.push(menu.value)
  if (children.value.length === 0) {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    void router.push({ path: menu.value.target, query: menu.value.query })
  }
}

</script>

<style lang='sass' scoped>
.item
  line-height: 56px
  width: 100%

.container
  width: 100%

.ctive
  background-color: $grey-3
  font-weight: bold

.icon
  margin-right: 10px
</style>
