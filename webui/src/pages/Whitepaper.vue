<template>
  <blog-page title='White Paper' author='KK' download-url='' :letters=1000 v-model:lang='lang' :time-need=3.3>
    <template #CN>
      <WhitepaperCN v-if='lang === "CN"' />
    </template>
    <template #EN>
      <WhitepaperEN v-if='lang === "EN"' />
    </template>
  </blog-page>
</template>

<script setup lang='ts'>
import { defineAsyncComponent, ref } from 'vue'

const BlogPage = defineAsyncComponent(() => import('../components/BlogPage/Page.vue'))
const WhitepaperEN = defineAsyncComponent(() => import('../components/Whitepaper/EN/Whitepaper.vue'))
const WhitepaperCN = defineAsyncComponent(() => import('../components/Whitepaper/CN/Whitepaper.vue'))

const lang = ref('CN')

</script>
