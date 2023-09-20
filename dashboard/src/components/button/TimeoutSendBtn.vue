<template>
  <button class='send-code alt' @click='onClick' :disabled='disabled'>
    {{ disabled ? timeout.toString() + 'S' : $t('MSG_SEND_CODE') }}
  </button>
</template>

<script setup lang='ts'>
import { defineEmits, ref, onMounted, onUnmounted, defineProps, toRef, watch } from 'vue'

interface Props {
  initialClicked: boolean;
  targetError: boolean;
}

const props = defineProps<Props>()
const initialClicked = toRef(props, 'initialClicked')
const targetError = toRef(props, 'targetError')

const disabled = ref(initialClicked.value)
watch(targetError, () => {
  if (targetError.value) {
    disabled.value = false
  }
})

const ticker = ref(-1)
const timeout = ref(60)

const startTimer = () => {
  timeout.value = 60
  disabled.value = true
  ticker.value = window.setInterval(() => {
    timeout.value -= 1
    if (timeout.value <= 0) {
      window.clearInterval(ticker.value)
      ticker.value = -1
      disabled.value = false
    }
  }, 1000)
}

const emit = defineEmits<{(e: 'click'): void}>()
const onClick = () => {
  if (targetError.value) {
    return
  }
  emit('click')
  startTimer()
}

onMounted(() => {
  if (initialClicked.value) {
    startTimer()
  }
})

onUnmounted(() => {
  if (ticker.value >= 0) {
    window.clearInterval(ticker.value)
    ticker.value = -1
  }
})

</script>
