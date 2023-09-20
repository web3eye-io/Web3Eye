<template>
  <div view='lHh Lpr lF' class="main-container">
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
        <div>Our mission is to organize the web3.0 information of different blockchain / ecosystem </div>
        <div>and make the web3.0 easier to the whole world.</div>
      </div>
    </div>
    <div class="column input-container">
      <input class="upload" id="drop-area" v-model="contract" />
      <q-icon name="img:icons/search.png" size="18px" class="search" />
    </div>
  </div>
</template>

<script setup lang='ts'>
import { onMounted, ref } from 'vue'
import largelogo from '../../assets/logo/large-logo.png'
import { useRouter } from 'vue-router'
import { useContractStore } from 'src/teststore/contract';
import { useTokenStore } from 'src/teststore/token';
import { Cookies } from "quasar"
const contract = ref('')
const _contract = useContractStore()

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

const router = useRouter()

const token = useTokenStore()

onMounted(() => {
  const dropArea = document.getElementById('drop-area')
  dropArea?.addEventListener('drop', (e) => {
    e.stopPropagation()
    e.preventDefault()
    let formData = new FormData()
    const file = e.dataTransfer?.files[0]
    formData.append('UploadFile', file as Blob)
    formData.append('Limit', '20')
    contract.value = file?.name as string
    token.searchTokens(formData, (error: boolean) => {
      if (!error) {
        void router.push('/token')
      }
    })
  })
  dropArea?.addEventListener('dragenter', (e) => {
    e.stopPropagation()
    e.preventDefault()
    console.log('enter')
  })
  dropArea?.addEventListener('dragleave', (e) => {
    e.stopPropagation()
    e.preventDefault()
    console.log('leave')
  })
  dropArea?.addEventListener('keypress', (e) => {
    if (e.key != 'Enter') {
      return
    }
    e.stopPropagation()
    e.preventDefault()
    if (contract.value?.length === 0) return
    getContractAndTokens(0, 100)
  })
})
</script>

<style lang='sass' scoped>
.main-container
  flex-grow: 1
.summary
  width: 840px
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

.icontainer,.upload-box
  width: 850px
.icontainer
  margin-top: 50px
.looking
  margin: 10px 0 10px 0
  color: $grey-8

.occupier
  height: 240px

.input-padding
  padding-bottom: 5px

.upload-box
  background: none
  flex-direction: row
.search-container
  ::v-deep > div.q-uploader
    width: auto
    max-height: 160px
.q-uploader 
  ::v-deep .bg-white
    background: none !important
.upload,.input-container
  margin: 0 auto
  width: 940px
.upload
  display: block
  position: relative
  width: 940px
  margin: 0 auto
  margin-top: 40px
  padding-left: 40px
  height: 48px
  border-radius: 24px
  border: 1px solid #3187FF
  &:focus
    outline: 1px solid #3187FF
.search
  display: inline-block
  position: relative
  padding-left: 20px
  line-height: 45px
  top: -43px
</style>
