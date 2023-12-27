<template>
  <div class="outer-bg">
    <div class="outer-container">
      <q-img :src="contractbg" :img-style='{borderRadius: "12px"}' />
      <div class="row items-center justify-center">
        <q-avatar v-if='current.ProfileURL?.length > 0'>
          <img :src="current.ProfileURL">
        </q-avatar>
      </div>
      <div class="collection column items-center justify-center">
        <div class="name">
          {{ current.Name }}
        </div>
        <!-- <div class="table row">
          <div class="listed">1.4 <span>%</span></div>
          <div class="owner">11 <span>K</span></div>
          <div class="price">0.03 <span>ETH</span></div>
          <div class="volume">0.03 <span>ETH</span></div>
        </div> -->
        <div class="description">
          {{ current.Description }}
        </div>
      </div>
      <div class="content">
        <div class="nav row">
          <q-tabs v-model="tab" dense class="text-grey" active-color="primary" indicator-color="primary" align="justify"
            narrow-indicator>
            <q-tab name="Collections" label="Collections" />
            <q-tab name="Transfers" label="Activity" />
          </q-tabs>
        </div>
      </div>
      <div id="contract">
        <div class="inner grid-container" v-if="tab == 'Collections'">
          <template class="box" v-for="token in tokens" :key="token.ID">
            <TokenCard :token="token" @click="onTokenClick(token)" />
          </template>
        </div>
        <div v-else>
          <q-table flat bordered :rows="transfers" :columns="(columns as any)" row-key="name" :rows-per-page-options='[20]'>
            <template v-slot:body="props">
              <q-tr :props="props">
            <q-td key="OfferItems" :props="props">
              <span v-if='props.row.OfferItems?.length === 0' />
              <div v-else class="row justify-start">
                <div class="left">
                  <MyImage
                    :url="(props.row.OfferItems?.[0].ImageURL as string)"
                    :height="'40px'"
                    :width="'40px'"
                  />
                </div>
                <div class="column items-start right">
                  <div class="token"># {{props.row.OfferItems?.[0]?.TokenID}}</div>
                  <div class="show-more">
                    Show More(Hover)
                    <q-tooltip
                      anchor="bottom right"
                      style="width: 400px"
                      self="center middle"
                      class="bg-white text-black shadow-2"
                      :offset="[60, 60]"
                    >
                      <TransferFloatItem :offer-items="props.row.OfferItems" :target-items="props.row.TargetItems" />
                    </q-tooltip>
                  </div>
                </div>
              </div>
            </q-td>
            <q-td key="Transfer" :props="props">
              <span v-if='props.row.OfferItems?.length === 0' />
              <div v-else class="row justify-start">
                <div class="column items-start right">
                  <div class="show-more">
                    <q-icon name="img:icons/transfer.png" size="20px"/>
                  </div>
                </div>
              </div>
            </q-td>
            <q-td key="Transfer1" :props="props">
            </q-td>
            <q-td key="TargetItems" :props="props">
              <span v-if='props.row.TargetItems?.length === 0' />
              <div v-else class="row justify-start">
                <div class="left">
                  <MyImage
                    :url="(props.row.TargetItems?.[0]?.ImageURL as string)"
                    :height="'40px'"
                    :width="'40px'"
                  />
                </div>
                <div class="column items-start right">
                  <div class="token"># {{props.row.TargetItems?.[0]?.TokenID}}</div>
                  <div class="show-more">
                    Show More(Hover)
                    <q-tooltip
                      anchor="bottom right"
                      style="width: 400px"
                      self="center middle"
                      class="bg-white text-black shadow-2"
                      :offset="[60, 60]"
                    >
                      <TransferFloatItem :offer-items="props.row.OfferItems" :target-items="props.row.TargetItems" />
                    </q-tooltip>
                  </div>
                </div>
              </div>
            </q-td>
            <q-td key="From" :props="props">
              <ToolTip :address="props.row.From" />
            </q-td>
            <q-td key="To" :props="props">
              <ToolTip :address="props.row.To" />
            </q-td>
            <q-td key="TxTime" :props="props">
              {{ formatTime(props.row.TxTime) }}
            </q-td>
          </q-tr>
            </template>
          </q-table>
        </div>
        <div style="height: 150px;"></div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useContractStore } from 'src/teststore/contract'
import { computed, defineAsyncComponent, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import contractbg from '../../assets/material/contract-bg.png'
import { useTransferStore } from 'src/teststore/transfer';
import { ChainType } from 'src/teststore/basetypes/const';
import { Transfer } from 'src/teststore/transfer/types';
import { formatTime } from 'src/teststore/util'
import { Contract, ShotToken } from 'src/teststore/contract/types';
const ToolTip = defineAsyncComponent(() => import('src/components/Token/ToolTip.vue'))
const TokenCard = defineAsyncComponent(() => import('src/components/Token/TokenCard.vue'))
const MyImage = defineAsyncComponent(
  () => import('src/components/Token/Image.vue')
)
const TransferFloatItem = defineAsyncComponent(
  () => import('src/components/Token/TransferFloatItem.vue')
)
const tab = ref('Collections')
const contract = useContractStore()
const tokens = computed(() => contract.shotTokens(_contract.value))
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
  if (url.startsWith('ipfs://')) {
    return url.replace('ipfs://', 'https://ipfs.io/ipfs/')
  }
  if (url.startsWith('data:image')) {
    return `img:${url}`
  }
  return url
})

const getContract = () => {
  contract.getContractAndTokens({
    Contract: _contract.value,
    Offset: 0,
    Limit: 100,
    Message: {}
  }, (error: boolean, row: Contract) => {
    if (!error) {
      getTransfers(0, 100, row.ChainID, row.ChainType)
    }
  })
}

const transfer = useTransferStore()
const key = computed(() => transfer.setKey(_chainID.value, _contract.value, undefined as unknown as string))
const transfers = computed(() => transfer.getTransfersByKey(key.value))

const getTransfers = (offset: number, limit: number, chainID: string, chainType:string) => {
  transfer.getTransfers({
    ChainID: chainID,
    ChainType: chainType as unknown as ChainType,
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
      getTransfers(offset + limit, limit, chainID, chainType)
    })
}

const columns = computed(() => [
  {
    name: 'OfferItems',
    label: 'Offer Items',
    align: 'left',
  },
  {
    name: 'Transfer',
    label: '',
    align: 'left',
  },
  {
    name: 'Transfer1',
    label: '',
    align: 'left',
  },
  {
    name: 'TargetItems',
    label: 'Target Items',
    align: 'left',
  },
  {
    name: 'From',
    label: 'From',
    align: 'left',
  },
  {
    name: 'To',
    label: 'To',
    align: 'left',
  },
  {
    name: 'TxTime',
    label: 'Time',
    align: 'left',
  },
])

const router = useRouter()
const onTokenClick = (token: ShotToken) => {
  void router.push({
    path: '/token/detail',
    query: {
      chainID: _chainID.value,
      chainType: token.ChainType,
      contract: _contract.value,
      tokenID: token.TokenID,
      id: token.ID,
    }
  })
}

onMounted(() => {
  if (transfers.value?.length === 0) {
    if (_chainID.value?.length === 0 || _chainType.value?.length === 0) return
    getTransfers(0, 100, _chainID.value, _chainType.value)
  }
  if (_contract?.value?.length > 0) {
    getContract()
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
  
.token
  font-size: 16px
.show-more
  color: #1772F8
  font-size: 12px
.token,.show-more
  padding-left: 5px
</style>