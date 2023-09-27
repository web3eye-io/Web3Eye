<template>
  <q-btn
    flat
    dense
    :icon='icon'
    @click='onClick'
  />
</template>

<script setup lang='ts'>
import { defineProps, withDefaults, computed, defineEmits } from 'vue'

interface Props {
  icon: string
}

const props = withDefaults(defineProps<Props>(), {
  icon: 'icons/github.png'
})

const icon = computed(() => 'img:' + props.icon)

const emit = defineEmits<{(e: 'click'): void }>()
const onClick = () => {
  emit('click')
}

</script>
