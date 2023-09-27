<template>
    <q-btn :label="displayLabel" flat @click="onUpdate">
        <q-tooltip anchor="top middle" self="center middle" class="bg-white text-black shadow-2">
          {{ displayMessage }}
        </q-tooltip>
    </q-btn>
</template>
<script lang="ts" setup>
import { computed, ref, toRef } from 'vue';
import { copyToClipboard } from 'quasar'
interface Props {
  address: string;
}

const props = defineProps<Props>()
const address = toRef(props, 'address')

const displayMessage = ref(address.value)
const displayLabel = computed(() => {
    const start = address.value?.substring(0, 6) 
    const end = address.value?.substring(address.value?.length - 4)
    return `${start}...${end}`
})

const onUpdate = () => {
    void copyToClipboard(address.value)
    displayMessage.value = 'Copied!'
    setInterval(() => {
        displayMessage.value = address.value
    }, 3000)
}
</script>
<style lang="sass" scoped>
button
  ::v-deep .q-hoverable
    &:hover   
      background-color: none
.tooltip-background
    background: $white
</style>
