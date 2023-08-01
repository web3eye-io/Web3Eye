<template>
  <q-icon
    v-if='imageType === ImageType.ICON'
    size='300px' 
    :name='imageUrl'
  />
  <q-img
    v-else
    :src="imageUrl"
    spinner-color="red"
    :height="height" 
    :width="width" 
    fit="fill" 
    class="rounded-borders"
  />
</template>

<script lang="ts" setup>
import { computed, toRef } from 'vue';

interface Props {
  url: string;
  height?: string
  width?: string
}

const props = defineProps<Props>()
const url = toRef(props, 'url')
const height = toRef(props, 'height')
const width = toRef(props, 'width')

enum ImageType {
  ICON = 'ICON',
  IMG = 'IMG'
}

const imageType = computed(() => {
  if (url.value?.startsWith('data:image')) {
    return ImageType.ICON
  }
  return ImageType.IMG
})

const imageUrl = computed(() => {
  if(url.value?.startsWith('ipfs://')) {
      return url.value?.replace('ipfs://', 'https://ipfs.io/ipfs/')
  }
  if (url.value?.startsWith('data:image')) {
    return `img:${url.value}`
  }
  return url.value
})

</script>