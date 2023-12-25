<template>
    <div class="row box">
        <div class="left"><q-icon name="img:icons/search.png" size="20px" /></div>
        <div class="main">
            <input
                class="search-box"
                id="drop-area"
                placeholder="search contract address or drag an image here"
                v-model="contract" 
            />
        </div>
        <div class="right"><q-icon name="img:icons/camera.png" class="photography" size="20px" /></div>
    </div>
</template>
  
<script setup lang='ts'>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useContractStore } from 'src/teststore/contract'
import { useTokenStore } from 'src/teststore/token'
import { SearchTokenMessage } from 'src/teststore/token/types'

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
        formData.append('Limit', '8')
        contract.value = file?.name as string
        const reqMessage = {} as SearchTokenMessage
        token.$reset()
        token.searchTokens(formData, reqMessage, (error: boolean) => {
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
  </style>
  