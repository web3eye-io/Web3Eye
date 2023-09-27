<template>
    <SyncTask />
  </template>
  
  <script setup lang='ts'>
  import { defineAsyncComponent } from 'vue'
  
  const SyncTask = defineAsyncComponent(() => import('src/components/task/SyncTask.vue'))
  
  </script>
  