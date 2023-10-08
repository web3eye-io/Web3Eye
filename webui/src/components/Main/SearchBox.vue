<template>
    <div>
        <input class="upload" id="drop-area" placeholder="search contract address or drag an image here"
        v-model="contract" />
    <q-icon name="img:icons/search.png" size="18px" class="search" />
    </div>
</template>
  
<script setup lang='ts'>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useContractStore } from 'src/teststore/contract'
import { useTokenStore } from 'src/teststore/token'
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
  .upload,.input-container
    margin: 0 auto
    width: 100%
  .upload
    display: block
    position: relative
    width: 100%
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
    top: -45px
  </style>
  