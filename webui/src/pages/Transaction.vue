<template>
  <Transaction />
</template>

<script setup lang='ts'>import { defineAsyncComponent } from 'vue';

const Transaction = defineAsyncComponent(() => import('src/components/Transaction/Transaction.vue'))

</script>