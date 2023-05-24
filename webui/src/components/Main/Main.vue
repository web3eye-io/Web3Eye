<template>
  <div view='lHh Lpr lFf'>
    <div class='row logo'>
      <q-space />
      <q-img :src='logo' style='width: 400px' fit='contain' />
      <q-space />
    </div>
    <q-input
      v-if='isText'
      class='icontainer'
      rounded
      outlined
      v-model="search"
      @keyup.enter='handleEnter'
      placeholder="input text here"
    >
      <template v-slot:append>
        <InputOption v-model:option='curOption' />
      </template>
    </q-input>
    <q-file
      v-if='!isText'
      :clearable='uploading'
      class='icontainer'
      v-model="file"
      rounded
      outlined
      :loading='uploading'
      name='upload'
      @update:model-value='onUpdate'
      placeholder="drag a image here"
      @clear='handleClear'
      @rejected='handleReject'
    >
      <template v-slot:append>
        <InputOption v-model:option='curOption' />
      </template>
    </q-file>
    <div class='occupier' />
  </div>
</template>

<script setup lang='ts'>
import { useNFTMetaStore } from 'src/localstore/nft';
import { NFTMeta } from 'src/localstore/nft/types';
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router';
import InputOption from 'src/components/Main/InputOption.vue'
import logo from '../../assets/logo/logo.png'
import { api } from 'src/boot/axios';

const curOption = ref('File')
const isText = computed(() => curOption.value === 'Text')

const search = ref('')
const file = ref({} as File)

const router = useRouter()

const nft = useNFTMetaStore()

const uploading = ref(false)
// Contract search
const handleEnter = () => {
  console.log('enter......')
}

// image search
const onUpdate  = () => {
  uploading.value = true
  const reader = new FileReader()
  reader.readAsBinaryString(file.value)
  reader.onload = function () {
    api.post('/api/nft-meta/search/file', {
      'topN': 10,
      'file': reader.result
    }, {
      headers: {'Content-Type': 'multipart/form-data'}}
    )
    .then((response) => {
      console.log('response: ', response.data)
      nft.setNftMeta(response.data as Array<NFTMeta>)
      void router.push({
        path: '/result'
      })
    })
    .catch((error)=> {
      console.log('error: ', error)
    })
    .finally(() => {
      uploading.value = false
    })
  }
}

const handleReject = () => {
  console.log('handleReject')
  uploading.value = false
}

const handleClear = () => {
  console.log('clear')
  file.value = {} as File
}

</script>

<style scoped lang='sass'>
.logo
  margin: 10px 0 20px 0

.icontainer
  width: 650px
.looking
  margin: 10px 0 10px 0
  color: $grey-8

.occupier
  height: 240px

.upload-box
  width: 100%
  height: 300px
  max-height: 300px
</style>
