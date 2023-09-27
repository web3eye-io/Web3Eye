<template>
    <Endpoint />
  </template>
  
  <script setup lang='ts'>
  import { defineAsyncComponent } from 'vue'
  
  const Endpoint = defineAsyncComponent(() => import('src/components/endpoint/Endpoint.vue'))
  
  </script>
  