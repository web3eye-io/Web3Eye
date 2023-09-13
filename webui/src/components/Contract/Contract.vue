<template>
  <div class="outer-bg">
    <div class="outer-container">
      <q-img :src="contractbg" />
      <div class="row items-center justify-center">
        <q-avatar>
          <img src="https://cdn.quasar.dev/img/avatar.png">
        </q-avatar>
      </div>
      <div class="collection column items-center justify-center">
        <div class="name">
          {{ current.Name }}
        </div>
        <div class="table row">
          <div class="listed">1.4 <span>%</span></div>
          <div class="owner">11 <span>K</span></div>
          <div class="price">0.03 <span>ETH</span></div>
          <div class="volume">0.03 <span>ETH</span></div>
        </div>
        <div class="description">
          {{ current.Description }}
          Spesh is looking for his best friend throughout Coolman's Universe. To travel through this universe, Spesh uses a surfboard and a magical compass.
        </div>
      </div>
      <div class="content">
        <div class="nav row">
          <q-tabs
            v-model="tab"
            dense
            class="text-grey"
            active-color="primary"
            indicator-color="primary"
            align="justify"
            narrow-indicator
            @update:model-value="onChangeTab"
          >
            <q-tab name="Collections" label="Collections" />
            <q-tab name="Transfers" label="Transfers" />
          </q-tabs>
        </div>
      </div>
      <div id="contract">
        <div class="inner grid-container" v-if="tab == 'Collections'">
          <div class="box" v-for="token in tokens" :key="token.ID">
            <div class="box-img">
              <MyImage :url="token.ImageURL" :height="'220px'" />
            </div>
            <div class="content">
              <div class="line row justify-between">
                <span class="title">#{{token.TokenID}}</span>
                <div class="row fee">
                  <span>4.75</span>
                  <q-icon name="img:icons/ethereum-eth-logo.png" style="padding-top: 3px;" />
                </div>
              </div>
              <div class="super row justify-between">
                <div>{{token.Name}}</div>
                <div class="transfers">
                  {{token.TransfersNum}}
                  <q-icon name="img:icons/transfers.png" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else>
          <q-table
          row-key="Block" 
          flat bordered
          :columns="(columns as any)"
          :rows="transfers"
        />
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useContractStore } from 'src/teststore/contract'
import { computed, defineAsyncComponent, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import contractbg from '../../assets/material/contract-bg.png'
import { Transfer } from 'src/teststore/transfer/types';
import { useTransferStore } from 'src/teststore/transfer';
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

const tab = ref('Collections')
const contract = useContractStore()
const tokens = computed(() => contract.ShotTokens.ShotTokens)
const current = computed(() => contract.Contract)

interface Query {
  contract: string;
  tokenID: string;
}

const route = useRoute()
const query = computed(() => route.query as unknown as Query)
const _contract = computed(() => query.value.contract)
const tokenID = computed(() => query.value.tokenID)

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const getImageUrl = computed(() => (url: string) => {
  if(url.startsWith('ipfs://')) {
      return url.replace('ipfs://', 'https://ipfs.io/ipfs/')
  }
  if (url.startsWith('data:image')) {
    return `img:${url}`
  }
  return url
})

onMounted(() => {
  if (_contract?.value?.length > 0) {
    getContract()
  }
})

const getContract = () => {
  contract.getContractAndTokens({
    Contract: _contract.value,
    Offset: 0, 
    Limit: 100,
    Message: {}
  }, () => {
    // TODO
  })
}

const columns = computed(() => [
  {
    name: 'Item',
    label: 'Item',
    align: 'center',
    field: () => ''
  },
  {
    name: 'Block',
    label: 'BLOCK',
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
    name: 'Value',
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

const transfer = useTransferStore()
const transfers = computed(() => transfer.Transfers.Transfers)

const onChangeTab = () => {
  if (transfers.value?.length == 0) {
    getTransfers(0, 500)
  }
}

const getTransfers = (offset: number, limit: number) => {
  transfer.getTransfers({
    ChainType: current.value.ChainType,
    ChainID: current.value.ChainID,
    Contract: current.value.Address,
    TokenID: tokenID.value,
    Offset: offset,
    Limit: limit,
    Message: {}
  }, (error:boolean, rows: Transfer[]) => {
    if (error || rows.length < limit) {
      return
    }
    getTransfers(offset, offset + limit)
  })
}

</script>
<style lang="sass" scoped>
.q-avatar
  top: -40px
  height: 1.5em
  width: 1.5em
  max-height: 33px
.q-avatar__content
  flex-grow: 1
.name
  color: #F5841F
  font-size: 40px
.description
  width: 480px
  margin: 0 auto
  padding-top: 10px
  opacity: 0.8
  text-align: center
.q-tab
  text-transform: none
  margin-top: 20px
::v-deep .q-tab__label
  font-size: 24px
.collection
  .table
    margin-top: 15px
    border-radius: 10px
    font-size: 24px
    background-color: #FAFAFA
    border: 1px solid #efeded
    div
      padding: 20px
      line-height: 32px
      border-right: 1px solid #efeded
      &:last-child
        border-right: none
#contract
  padding: 10px
  padding-right: 0
  padding-left: 0
  
  .grid-container
    display: grid
    grid-template-columns: repeat(auto-fill, minmax(auto, 220px))
    grid-gap: 12px  
    justify-content: space-between
    .box
      width: 220px
      max-width: 280px
      height: 308px
      border-radius: 10px
      border: 1px solid #efefef
      background-color: #FAFAFA
      .box-img
        padding: 8px
      .content
        .line
          padding: 5px
        .super
          padding: 0 5px 2px 5px
        .transfers
          min-width: 20px
</style>