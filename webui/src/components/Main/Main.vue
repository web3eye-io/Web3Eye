<template>
  <div view='lHh Lpr lF'>
    <div class="summary">
      <div class="easier row items-center">
        <q-space />
        <div class="left" />
        <div class="center">Here we make web3.0 easier</div>
        <div class="right" />
        <q-space />
      </div>
      <div class="name">
        <q-img :src='largelogo' class='logo' fit="contain" />
      </div>
      <div class="mission column items-center">
        <div>Our mission is to organize the web3.0 information of different blockchain /</div>  
        <div>ecosystem and make the web3.0 easier to the whole world.</div>
      </div>
    </div>
    <!-- <div class='row logo'>
      <q-space />
      <q-img :src='logo' style='width: 400px' fit='contain' />
      <q-space />
    </div> -->
    <q-input
      v-if='isText'
      class='icontainer input-padding'
      rounded
      outlined
      v-model="contract"
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
      field-name='UploadFile'
      :form-fields='[{name: "Limit", value: "20"}]'
      auto-upload
      flat
      accept=".jpg, image/*"
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
import { computed, ref } from 'vue'
import InputOption from 'src/components/Main/InputOption.vue'
import largelogo from '../../assets/logo/large-logo.png'

import { useRouter } from 'vue-router'
import { useTokenStore } from 'src/teststore/token';
import { GetTokensResponse } from 'src/teststore/token/types';
import { useContractStore } from 'src/teststore/contract';

const curOption = ref('File')
const isText = computed(() => curOption.value === 'Text')

const contract = ref('')
const _contract = useContractStore()

const handleEnter = () => {
  console.log('enter......', contract)
  getContractAndTokens(0, 100)
}
const getContractAndTokens = (offset: number, limit: number) => {
  _contract.getContractAndTokens({
    Contract: contract.value,
    Offset: offset,
    Limit: limit,
    Message: {}
  }, (error: boolean) => {
    if (error) return
    void router.push('/contract')
  })
}
const fileName = ref('')
const uploading = ref(false)

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const onAdded = (files: readonly any[]) => {
  const _file = files[0] as File
  fileName.value = _file.name
}


const router = useRouter()

const token = useTokenStore()
const onUploaded = (info: {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    files: readonly any[];
    xhr: XMLHttpRequest;
  }) => {
  const reader = new FileReader()
  reader.readAsDataURL(info.files[0] as Blob)
  reader.onload = function() {
    token.SearchTokens.Current = window.URL.createObjectURL(info.files[0] as Blob)
	}
  const response = JSON.parse(info.xhr.response as string) as GetTokensResponse
  token.setToken(response.Infos)
  token.SearchTokens.Total = response.TotalTokens
  token.SearchTokens.StorageKey = response.StorageKey
  token.SearchTokens.TotalPages = response.TotalPages
  void router.push({
    path: '/token'
  })
  uploading.value = false
}

const onFailed = () => {
  uploading.value = false
  console.log('onFailed...')
}

</script>

<style lang='sass' scoped>
.summary
  width: 100%
  margin: 0 auto
  .easier
    font-size: 16px
    color: #1772F8

    .left,.right
      width: 100px
      height: 4px
      border-radius: 2px
    .left
      background: linear-gradient(to left, transparent 0, #3187FF 0%, transparent 100%)
    .right
      background: linear-gradient(to right, transparent 0, #3187FF 0%, transparent 100%)
    .center
      padding: 0 15px
  .name 
    font-size: 8rem
  .mission
    font-size: 20px
    opacity: 0.8
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
