<template>
  <div class='q-pa-md q-gutter-sm'>
    <q-tree
      :nodes='contentTable'
      node-key='label'
      default-expand-all
    >
      <template v-slot:default-header='prop'>
        <div class='row items-center'>
          <div class='text-weight-bold text-primary'>{{ prop.node.label }}</div>
        </div>
      </template>
    </q-tree>
  </div>
</template>

<script setup lang='ts'>
import { QTreeNode } from 'quasar'
import { toRef } from 'vue'


interface Prop {
  contentTable: QTreeNode<unknown>[]
}

const prop = defineProps<Prop>()
const contentTable = toRef(prop, 'contentTable')

</script>
