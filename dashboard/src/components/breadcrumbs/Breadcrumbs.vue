<template>
  <div>
    <q-breadcrumbs
      separator='>'
    >
      <q-breadcrumbs-el
        v-for='router in breadcrumbs'
        :key='router.label'
        :icon='router.icon'
        :to='router.target'
        :label='router.label'
        disable
      />
    </q-breadcrumbs>
  </div>
</template>

<script setup lang='ts'>
import { MenuItem } from 'src/localstore'
import { toRef, defineProps } from 'vue'

interface Props {
  breadcrumbs: Array<MenuItem>
}

const props = defineProps<Props>()
const breadcrumbs = toRef(props, 'breadcrumbs')

</script>

<style lang='sass' scoped>
</style>
