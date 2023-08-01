<template>
  <div id="transfer">
    <div class="row justify-center q-pa-md">
      <div class="left col-4">
        <!-- <q-img
          src="https://picsum.photos/500/300"
          height="100%"
          width="100%"
          class="transfer-img rounded-borders"
          fit="fill"
        /> -->
        <MyImage :url="token.ImageURL" :height="'100%'" :width="'100%'" />
      </div>
      <div class="right flex column col">
        <div class="header">{{token.Name}}</div>
        <div class="title">{{token.Owner}}</div>
        <div class="description">
          {{token.Description}}
        </div>
        <div class="url">
          <span>ImageURL</span>
          <span>{{token.ImageURL}}</span>
        </div>
        <div class="ipfs">
          <span>IPFS-URL</span>
          <span>{{token.IPFSImageURL}}</span>
        </div>
        <div class="contract">
          <span>Contract</span>
          <span>{{ token.Contract }}</span>
        </div>
        <div class="chain">{{token.ChainType}} @ Goerli</div>
      </div>
    </div>
    <div class="transfer-table">
      <div class="q-pa-md">
        <h5>Transfer</h5>
        <q-table
          row-key="Block" 
          flat bordered
          :columns="columns"
          :rows="transfers"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { SearchToken } from 'src/teststore/token/types'
import { Transfer } from 'src/teststore/transfer/types'
import { computed, defineAsyncComponent, ref } from 'vue'
import { toRef } from 'vue'

const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

interface Props {
  transfers: Transfer[];
  token: SearchToken
}

const props = defineProps<Props>()
const _transfers = toRef(props, 'transfers')
const _token = toRef(props, 'token')

const transfers = ref(_transfers) 
const token = ref(_token) 

const columns = computed(() => [
  {
    name: 'Block',
    label: 'BLOCK',
    field: (row: Transfer) => row.BlockNumber
  },
  {
    name: 'Time',
    label: 'Time',
    field: (row: Transfer) => row.TxTime
  },
  {
    name: 'Value',
    label: 'Value',
    field: (row: Transfer) => row.Amount
  },
  {
    name: 'From',
    label: 'From',
    field: (row: Transfer) => row.From
  },
  {
    name: 'To',
    label: 'To',
    field: (row: Transfer) => row.To
  },
])
</script>
<style lang="sass" scoped>
#transfer
  width: 100%
  margin: 0 auto
  padding: 10px
  padding-left: 0
  .right
    padding: 0 15px 15px 15px
    .header,.title
      font-weight: bolder
      font-size: 16px
      padding: 5px 0
    .description
        padding: 20px 0
    .url,.ipfs,.contract
      padding: 2px 0
      span
        padding-right: 20px
    .chain
      padding-top: 20px
            
</style>
