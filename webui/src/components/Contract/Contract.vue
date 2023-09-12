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
            <q-tab name="Transfers" label="Transfers" />
          </q-tabs>
        </div>
      </div>
      <div id="contract">
        <div class="inner row" v-if="tab == 'Collections'">
          <div class="box column" v-for="token in tokens" :key="token.ID">
            <MyImage :url="token.ImageURL" :height="'180px'" :width="'180px'" />
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
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

const tab = ref("Collections")
const contract = useContractStore()
const tokens = computed(() => contract.ShotTokens.ShotTokens)
const current = computed(() => contract.Contract)

interface Query {
  contract: string;
}

const route = useRoute()
const query = computed(() => route.query as unknown as Query)
const _contract = computed(() => query.value.contract)

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

const transfers = ref([] as Array<Transfer>) 
const columns = computed(() => [
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
  margin: 0 auto
  padding: 10px
  padding-left: 0
  .right
    padding: 0 15px 15px 15px
    .header,.title
      font-weight: bolder
      font-size: 16px
      padding: 5px 0
    .contract
      padding: 2px 0
      span
        padding-right: 20px
  .box
    width: 180px
    border-radius: 10px
    border: 1px solid #efefef
    margin: 0 15px 15px 0
    .content
      .line
        padding: 5px
      .super
        padding: 0 5px 2px 5px
      .transfers
        min-width: 20px

</style>