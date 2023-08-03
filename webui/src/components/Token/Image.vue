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
  >
    <div class="absolute-bottom-left text-subtitle2">
      #{{title}}
    </div>
  </q-img>
</template>

<script lang="ts" setup>
import { computed, toRef } from 'vue';

interface Props {
  url: string;
  height?: string
  width?: string
  title?: string
}

const props = defineProps<Props>()
const url = toRef(props, 'url')
const height = toRef(props, 'height')
const width = toRef(props, 'width')
const title = toRef(props, 'title')

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
<style lang="sass" scoped>
.q-img__content > div
  padding: 2px
  background: none
</style>