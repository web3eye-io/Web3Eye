<template>
    <div class="row box">
        <div class="left"><q-icon name="img:icons/search.png" size="20px" /></div>
        <div class="main">
            <input class="search-box" id="drop-area" placeholder="search contract address or drag an image here"
                v-model="contract" />
        </div>
        <input ref='loadFileButton' type='file' style='display: none;' @change='uploadFile'>
        <div class="right"><q-icon name="img:icons/finder.png" class="photography" size="18px"
                @click='loadFileButton?.click()' /></div>
    </div>
</template>
  
<script setup lang='ts'>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useContractStore } from 'src/teststore/contract'
import { useTokenStore } from 'src/teststore/token'
import { SearchTokenMessage } from 'src/teststore/token/types'
import { useStorageKeyStore } from 'src/localstore/storagekey'

const contract = ref('')
const _contract = useContractStore()

const getContractAndTokens = (offset: number, limit: number) => {
    _contract.getContractAndTokens({
        Contract: contract.value,
        Offset: offset,
        Limit: limit,
        Message: {}
    }, (error: boolean) => {
        if (error) {
            return
        }
        void router.push({path: '/contract', query: {contract: contract.value} })
    })
}

const loadFileButton = ref<HTMLInputElement>()

const uploadFile = (evt: Event) => {
    const target = evt.target as unknown as HTMLInputElement
    if (target.files) {
        const file = target.files[0]
        const reader = new FileReader()
        reader.onload = () => {
            handleUploadFile(file)
        }
        reader.readAsText(file)
    }
}

const router = useRouter()
const token = useTokenStore()

const localkey = useStorageKeyStore()

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const handleUploadFile = (file: any) => {
    let formData = new FormData()
    formData.append('UploadFile', file as Blob)
    formData.append('Limit', '8')
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment, @typescript-eslint/no-unsafe-member-access
    contract.value = file?.name 
    const reqMessage = {} as SearchTokenMessage
    token.$reset()
    localkey.reset()
    token.searchTokens(formData, reqMessage, (error: boolean) => {
        if (!error) {
            void router.push('/token')
        }
    })
}

onMounted(() => {
    const dropArea = document.getElementById('drop-area')
    dropArea?.addEventListener('drop', (e) => {
        e.stopPropagation()
        e.preventDefault()
        const file = e.dataTransfer?.files[0]
        handleUploadFile(file)
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
  