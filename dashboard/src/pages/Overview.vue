<template>
    <Overview />
  </template>
  
  <script setup lang='ts'>
  import { defineAsyncComponent } from 'vue'
  
  const Overview = defineAsyncComponent(() => import('src/components/overview/Overview.vue'))
  
  </script>
  