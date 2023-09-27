<template>
  <q-btn
    :loading='loadingStatus'
    @click='simulateProgress'
    class='btn round'
    :disable='_disabled'
    :label='label'
  />
</template>

<script setup lang='ts'>
import { defineEmits, ref, defineProps, toRef, computed } from 'vue'

interface Props {
  loading: boolean;
  label: string;
  disabled?: boolean
}

const props = defineProps<Props>()
const loading = toRef(props, 'loading')
const label = toRef(props, 'label')
const disabled = toRef(props, 'disabled')
const _disabled = computed(() => !!disabled.value)

const loadingStatus = ref(false)

const emit = defineEmits<{(e: 'click', done: () => void): void}>()

const simulateProgress = () => {
  if (loading.value) {
    loadingStatus.value = true
  }
  emit('click', () => {
    loadingStatus.value = false
  })
}
</script>
