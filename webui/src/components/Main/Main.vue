<template>
  <div view='lHh Lpr lF' class="main-container" id="index">
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
      <div class="row box" id="normal-box" :class="[opening ? 'hidden' : '']">
        <div class="left"><q-icon name="img:icons/search.png" size="18px" /></div>
        <div class="main">
          <input class="search-box" id="search-box" placeholder="search contract address or drag an image here"
            v-model="contract" />
        </div>
        <!-- camera start -->
        <input ref='loadFileButton' type='file' style='display: none;' @change='uploadFile'>
        <div class="right">
          <q-icon name="img:icons/finder.png" class="photography" size="18px" @click='loadFileButton?.click()' />
        </div>
        <!-- camera end -->
      </div>
      <!-- drop zone start -->
      <div class="row big-box" id="drop-target" :class="[opening ? '' : 'hidden']">
        <q-icon name="img:icons/picture.png" size="42px" />
        <div class="drag-image-here">Drag an image here</div>
      </div>
      <div class="row big-box" :class="[state === State.Drop ? '' : 'hidden']">
        <Loading v-model:loading="loading" color="#1772F8" />
      </div>
      <!-- drop zone end -->
    </div>
  </div>
</template>

<script setup lang='ts'>
import { defineAsyncComponent, onMounted, ref } from 'vue'
import largelogo from '../../assets/logo/large-logo.png'
import { useRouter } from 'vue-router'
import { useContractStore } from 'src/teststore/contract'
import { useTokenStore } from 'src/teststore/token'
import { SearchTokenMessage } from 'src/teststore/token/types'
const Loading = defineAsyncComponent(() => import('src/components/Loading/Loading.vue'))

const loadFileButton = ref<HTMLInputElement>()

const uploadFile = (evt: Event) => {
  const target = evt.target as unknown as HTMLInputElement
  if (target.files) {
    const file = target.files[0]
    const reader = new FileReader()
    reader.onload = () => {
      handleUploadFile(file, false)
    }
    reader.readAsText(file)
  }
}

const router = useRouter()
const token = useTokenStore()

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const handleUploadFile = (file: any, fromDropArea: boolean) => {
  let formData = new FormData()
  formData.append('UploadFile', file as Blob)
  formData.append('Limit', '8')
  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment, @typescript-eslint/no-unsafe-member-access
  contract.value = file?.name
  const reqMessage = {} as SearchTokenMessage
  token.$reset()
  token.searchTokens(formData, reqMessage, (error: boolean) => {
    if (!error) {
      const normalBox = document.getElementById('normal-box')
      normalBox?.classList?.add('hidden')
      loading.value = false
      void router.push('/token')
      return
    }
    opening.value = false
    if (error) {
      if (fromDropArea) {
        state.value = State.Normal
        const dropArea = document.getElementById('drop-target')
        dropArea?.classList.remove('hidden')
      }
    }

  })
}

const loading = ref(true)
enum State {
  Normal,
  Dragging,
  Drop,
}
const state = ref(State.Normal)

onMounted(() => {
  const dropArea = document.getElementById('drop-target')
  dropArea?.addEventListener('drop', (e) => {
    e.stopPropagation()
    e.preventDefault()
    state.value = State.Drop
    dropArea?.classList.add('hidden')
    const file = e.dataTransfer?.files[0]
    handleUploadFile(file, true)
  })
  dropArea?.addEventListener('dragenter', (e) => {
    e.stopPropagation()
    e.preventDefault()
    state.value = State.Dragging
    dropArea.classList.add('highlight')
  })

  const searchBox = document.getElementById('search-box')
  searchBox?.addEventListener('keypress', (e) => {
    if (e.key != 'Enter') {
      return
    }
    e.stopPropagation()
    e.preventDefault()
    if (contract.value?.length === 0) return
    getContractAndTokens(0, 100)
  })
})

const opening = ref(false)

onMounted(() => {
  const dropZone = document.getElementById('index')
  dropZone?.addEventListener('dragover', function (e) {
    e.preventDefault()
    e.stopPropagation()
    opening.value = true
  })
  dropZone?.addEventListener('dragleave', (e) => {
    e.stopPropagation()
    e.preventDefault()
    let relatedTarget = e.relatedTarget
    if (!relatedTarget) { // leave window
      opening.value = false
    }
  })
})

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


.looking
  margin: 10px 0 10px 0
  color: $grey-8

.occupier
  height: 240px

.input-container
  margin: 0 auto
  width: 940px
  position: absolute
  left: 0 // important
  right: 0 // important
.box
  margin-top: 40px
  border: 1px solid #3187FF
  border-radius: 24px
  background: $white
  justify-content: center
.left
  width: 40px
  align-self: center
  padding-left: 20px
.main
  flex-grow: 1
  .search-box
    padding: 4px
    width: 100%
    height: 48px
    border-radius: 24px
    background: $white
    border: none
    &:focus
       outline: none
.right
  width: 40px
  align-self: center
  cursor: pointer
.big-box
    margin-top: 40px
    height: 130px
    border: 1px dashed #a5a5a6
    background: rgb(248,249,250)
    border-radius: 19px
    justify-content: center
    align-items: center
    .drag-image-here
        padding-left: 8px
        color: rgb(95,99,104)
.hidden
    display: none
#drop-target.highlight
    background: #f0f6ff
    border: 1px dashed #1772F8
</style>
