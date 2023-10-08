<template>
  <div id="transfer">
    <div class="row justify-center q-pa-md">
      <div class="left col-4">
        <MyImage :url="token.ImageURL" :height="'230px'" :width="'100%'" />
      </div>
      <div class="right flex column col">
        <div class="header">{{token.Name}}</div>
        <div class="title">{{token.Owner}}</div>
        <div class="description">
          {{token.Description}}
        </div>
        <div class="url" v-if="token?.ImageURL?.length > 0">
          <span>ImageURL</span>
          <span>{{token.ImageURL}}</span>
        </div>
        <div class="ipfs" v-if="token?.IPFSImageURL?.length > 0">
          <span>IPFS-URL</span>
          <span>{{token.IPFSImageURL}}</span>
        </div>
        <div class="contract">
          <span>Contract</span>
          <span>{{ token.Contract }}</span>
        </div>
        <div class="chain">{{token.ChainType}}</div>
      </div>
    </div>
    <div class="transfer-table">
      <div class="q-pa-md">
        <h5>Transfer</h5>
        <q-table
          row-key="Block" 
          flat bordered
          :columns="(columns as any)"
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
    label: 'Block',
    align: 'center',
    field: (row: Transfer) => row.BlockNumber
  },
  {
    name: 'Time',
    label: 'Time',
    align: 'center',
    field: (row: Transfer) => row.TxTime
  },
  {
    name: 'Amount',
    label: 'Value',
    align: 'center',
    field: (row: Transfer) => row.Amount
  },
  {
    name: 'From',
    label: 'From',
    align: 'center',
    field: (row: Transfer) => row.From
  },
  {
    name: 'To',
    label: 'To',
    align: 'center',
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
