<template>
    <div class="loading">
        <q-inner-loading :showing="loadingStatus" :style='"color: " + customColor' />
    </div>
</template>
<script setup lang='ts'>
import {ref, defineProps, toRef } from 'vue'

interface Props {
  loading: boolean;
  color?: string
}

const props = defineProps<Props>()
const _loading = toRef(props, 'loading')
const _color = toRef(props, 'color')

const loadingStatus = ref(_loading.value)
const customColor = ref(_color.value)
</script>

<style lang="sass" scoped>
.loading
  ::v-deep .absolute-full
    position: relative
    top: auto
</style>