<template>
  <div view='lHh Lpr lFf search-container'>
    <div class='row logo'>
      <q-space />
      <q-img :src='logo' style='width: 400px' fit='contain' />
      <q-space />
    </div>
    <q-input
      v-if='isText'
      class='icontainer input-padding'
      rounded
      outlined
      v-model="search"
      @keyup.enter='handleEnter'
      placeholder="input text here"
    >
      <template v-slot:append>
        <InputOption v-model:option='curOption' :disable='uploading' />
      </template>
    </q-input>
    <q-uploader
      v-if='!isText'
      class='upload-box'
      url="/api/entrance/search/file"
      color='white'
      :square='false'
      field-name='upload'
      :form-fields='[{name: "topN", value: "10"}]'
      auto-upload
      flat
      :disable='uploading'
      @failed='onFailed'
      @uploaded='onUploaded'
      @added='onAdded'
      @uploading='uploading = true'
    >
      <template v-slot:header>
          <q-input
            class='icontainer'
            rounded
            outlined
            v-model="fileName"
            :loading="uploading"
            placeholder="drag a image here"
          >
          <q-uploader-add-trigger /><!-- trigger file picker -->
          <template v-slot:append>
            <InputOption v-model:option='curOption' :disable='uploading' />
          </template>
        </q-input>
      </template>
      <template v-slot:list>
      </template>
    </q-uploader>
  </div>
    <div class='occupier' />
</template>

<script setup lang='ts'>
import { useNFTMetaStore } from 'src/localstore/nft';
import { UploadResponse } from 'src/localstore/nft/types';
import { computed, ref } from 'vue'
import InputOption from 'src/components/Main/InputOption.vue'
import logo from '../../assets/logo/logo.png'
import { useRouter } from 'vue-router'

const curOption = ref('File')
const isText = computed(() => curOption.value === 'Text')

const search = ref('')
const handleEnter = () => {
  console.log('enter......')
}

const fileName = ref('')
const uploading = ref(false)

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const onAdded = (files: readonly any[]) => {
  const _file = files[0] as File
  fileName.value = _file.name
}


const router = useRouter()

const nft = useNFTMetaStore()
const onUploaded = (info: {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    files: readonly any[];
    xhr: XMLHttpRequest;
  }) => {
  const reader = new FileReader()
  reader.readAsDataURL(info.files[0] as Blob)
  reader.onload = function() {
    nft.NTFMetas.Current = window.URL.createObjectURL(info.files[0] as Blob)
	}
  const response = JSON.parse(info.xhr.response as string) as UploadResponse
  nft.setNftMeta(response.data)
  void router.push({
    path: '/result'
  })
  uploading.value = false
}

const onFailed = () => {
  uploading.value = false
  console.log('onFailed...')
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

.input-padding
  padding-bottom: 5px

.upload-box
  width: 650px
  flex-direction: row
.search-container
  ::v-deep > div.q-uploader
    width: auto
    max-height: 160px
</style>
