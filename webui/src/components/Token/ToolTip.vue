<template>
    <q-btn :label="displayLabel" flat @click="onUpdate">
        <q-tooltip>
          Some text as content of Tooltip
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
    displayMessage.value = 'copied!'
    setInterval(() => {
        displayMessage.value = address.value
    }, 2000)
}
</script>
