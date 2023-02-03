<template>
  <q-page class='row items-center justify-evenly'>
    <search-main />
  </q-page>
</template>

<script setup lang='ts'>
import { defineAsyncComponent } from 'vue'

const SearchMain = defineAsyncComponent(() => import('../components/Main/Main.vue'))

</script>
