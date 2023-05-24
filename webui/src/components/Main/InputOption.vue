<template>
    <q-btn-toggle
      v-model="_option"
      class="my-custom-toggle"
      no-caps
      rounded
      unelevated
      toggle-color="primary"
      color="white"
      size='md'
      text-color="primary"
      :options="options"
      @update:model-value='onUpdate'
      :disable='disable'
    />
</template>
<script setup lang='ts'>
import { ref, toRef } from 'vue';

interface Props {
  option: string;
  disable: boolean;
}

const props = defineProps<Props>()
const option = toRef(props, 'option')
const disable = toRef(props, 'disable')

const _option = ref(option.value)
const options = ref([
    {label: 'F', value: 'File'},
    {label: 'T', value: 'Text'}  
])

const emit = defineEmits<{(e: 'update:option', option: string): void}>()
const onUpdate = () => {
  emit('update:option', _option.value)
}

</script>
<style lang="sass" scoped>
.my-custom-toggle
  border: 1px solid #027be3
</style>