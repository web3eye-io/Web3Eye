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
          >
            <q-tab name="Collections" label="Collections" />
            <q-tab name="Transfers" label="Activity" />
          </q-tabs>
        </div>
      </div>
      <div id="contract">
        <div class="inner grid-container" v-if="tab == 'Collections'">
          <div class="box" v-for="token in tokens" :key="token.ID">
            <TokenCard :token="token" />
          </div>
        </div>
        <div v-else>
          Coming Soon
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
import { useTransferStore } from 'src/teststore/transfer';
import { ChainType } from 'src/teststore/basetypes/const';
import { Transfer } from 'src/teststore/transfer/types';

const TokenCard = defineAsyncComponent(() => import('src/components/Token/TokenCard.vue'))

const tab = ref('Collections')
const contract = useContractStore()
const tokens = computed(() => contract.ShotTokens.ShotTokens)
const current = computed(() => contract.Contract)

interface Query {
  contract: string;
  chainID: string;
  chainType: ChainType;
}

const route = useRoute()
const query = computed(() => route.query as unknown as Query)
const _contract = computed(() => query.value.contract)
const _chainID = computed(() => query.value.chainID)
const _chainType = computed(() => query.value.chainType)

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

const transfer = useTransferStore()
const key = computed(() => transfer.setKey(_chainID.value, _contract.value))
const transfers = computed(() => transfer.getTransfersByKey(key.value))

const getTransfers = (offset: number, limit: number) => {
  transfer.getTransfers({
    ChainID: _chainID.value,
    ChainType: _chainType.value,
    Contract: _contract.value,
    Offset: offset, 
    Limit: limit,
    Message: {}
  },
  key.value, 
  (err: boolean, rows: Array<Transfer>) => {
    if (err || rows.length === 0) {
      return 
    }
    getTransfers(offset + limit, limit)
  })
}

onMounted(() => {
  if (transfers.value?.length === 0) {
    getTransfers(0, 100)
  }
})
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
  margin-bottom: 15px
  padding-left: 0px
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
  

</style>