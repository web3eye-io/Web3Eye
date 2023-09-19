<template>
  <div class="outer-bg">
    <div class="outer-container">
      <div class="top row no-wrap">
        <div class="left">
          <MyImage 
            :url="(target?.ImageURL as string)" 
            :height="'460px'" 
            :width="'460px'"
          />
        </div>
        <div class="gap" />
        <div class="right column justify-between">
          <div class="name">
            {{ target?.Name }}
          </div>
          <div class="content">
            {{target?.Name}} #{{target?.TokenID}}
          </div>
          <div class="description">
            {{ target?.Description }}
          </div>
          <div class="author row justify-between">
            <div class="column">
              <div class="creator-title">Creator</div>
              <div class="row items-center">
                <q-avatar size="40px">
                  <!-- <img src="https://cdn.quasar.dev/img/avatar.png"> -->
                </q-avatar> 
                <div class="creator-name">
                  {{ target?.Owner }}
                </div>
              </div>
            </div>
            <div class="column">
              <div class="chain-title">Blockchain</div>
              <div class="row items-center justify-center">
                  <q-icon name="img:icons/ethereum-eth-logo.png" />
                  <div class="chain-name">{{ target?.ChainType }}</div>
              </div>
            </div>
            <div class="column">
              <div class="chain-title">TokenType</div>
              <div class="row items-center justify-center">
                  <div class="chain-name">{{ target?.TokenType }}</div>
              </div>
            </div>
            <div class="col-2" />
          </div>
          <div class="contract column">
            <div class="title">Contract</div>
            <div class="address">{{ target?.Contract }}</div>
          </div>
          <!-- <div class="price column">
            <div class="title">TokenType</div>
            <div class="amount">{{ target?.TokenType }}</div>
          </div> -->
          <q-btn class="buy" disable unelevated rounded color="primary" label="BUY NOW" />
        </div>
      </div>
      <div class="transfer">Transfer</div>
      <q-table
        flat 
        bordered
        :rows="transfers"
        :columns="(columns as any)"
        row-key="name"
        :rows-per-page-options='[20]'
      >
        <template v-slot:body="props">
          <q-tr :props="props">
            <q-td key="Block" :props="props">
              {{ props.row.BlockNumber }}
            </q-td>
            <q-td key="TxTime" :props="props">
                {{ formatTime(props.row.TxTime)}}
            </q-td>
            <q-td key="Value" :props="props">
                {{ props.row.Value }}
            </q-td>
            <q-td key="From" :props="props">
              <ToolTip :address="props.row.From" />
            </q-td>
            <q-td key="To" :props="props">
              <ToolTip :address="props.row.To" />
            </q-td>
          </q-tr>
        </template>
      </q-table>
      <div class="collections">More from this collection</div>
      <div class="inner grid-container">
          <template v-for="token in tokens" :key="token.ID">
            <TokenCard :token="token" @click="onShotTokenClick(token)" />
          </template>
      </div>
    </div>
  </div>
</template>
<script lang='ts' setup>
import { ChainType } from 'src/teststore/basetypes/const';
import { useContractStore } from 'src/teststore/contract';
import { useTokenStore } from 'src/teststore/token';
import { useTransferStore } from 'src/teststore/transfer';
import { formatTime } from 'src/teststore/util'
import { Transfer } from 'src/teststore/transfer/types';
import { computed, defineAsyncComponent, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { ShotToken } from 'src/teststore/contract/types';
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))
const TokenCard = defineAsyncComponent(() => import('src/components/Token/TokenCard.vue'))
const ToolTip = defineAsyncComponent(() => import('src/components/Token/ToolTip.vue'))

interface Query {
  chainID: string;
  chainType: ChainType;
  contract: string;
  tokenID: string;
  id: string
}

const route = useRoute()
const query = computed(() => route.query as unknown as Query)
const _contract = computed(() => query.value.contract)
const _chainID = computed(() => query.value.chainID)
const _chainType = computed(() => query.value.chainType)
const _tokenID = computed(() => query.value.tokenID)
const _id = computed(() => query.value.id)

const tokenID1 = ref(_tokenID.value)
const id1 = ref(_id.value)

const transfer = useTransferStore()
const transferKey = computed(() => transfer.setKey(_chainID.value, tokenID1.value))
const transfers = computed(() => transfer.Transfers.Transfers.get(transferKey.value))

const columns = computed(() => [
  {
    name: 'Block',
    label: 'BLOCK',
    align: 'center',
  },
  {
    name: 'TxTime',
    label: 'Time',
    align: 'center',
  },
  {
    name: 'Value',
    label: 'Value',
    align: 'center',
  },
  {
    name: 'From',
    label: 'From',
    align: 'center',
  },
  {
    name: 'To',
    label: 'To',
    align: 'center',
  },
])

const getTransfers = (offset: number, limit: number) => {
  transfer.getTransfers({
    ChainType: _chainType.value,
    ChainID: _chainID.value,
    Contract: _contract.value,
    TokenID: tokenID1.value,
    Offset: offset,
    Limit: limit,
    Message: {}
  },
  transferKey.value,
  (error:boolean, rows: Transfer[]) => {
    if (error || rows.length === 0) {
      return
    }
    getTransfers(offset+limit, limit)
  })
}

const token = useTokenStore()
const target = computed(() => token.getTokenByID(tokenID1.value))

const getToken = () => {
  token.getToken({
    ID: id1.value,
    Message: {}
  }, () => {
    // TODO
  })
}

const contract = useContractStore()
const tokens = computed(() => contract.ShotTokens.ShotTokens)
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

const onShotTokenClick = (token: ShotToken) => {
  tokenID1.value = token.TokenID
  id1.value = token.ID
  if (!target.value) {
    getToken()
  }
  if (!transfers.value || transfers.value?.length === 0) {
    getTransfers(0, 100)
  }
}

onMounted(() => {
  if (!target.value) {
    getToken()
  }
  if (!transfers.value || transfers.value?.length === 0) {
    getTransfers(0, 100)
  }
  if (_contract?.value?.length > 0) {
    getContract()
  }
})
</script>

<style lang="sass" scoped>
  .top
    .left
      width: 500px
      height: 500px
      border: 1px solid #f7f7f7
      border-radius: 16px
      padding-left: 20px
      padding-top: 20px
      box-shadow: 5px 5px 5px #f7f7f7
    .gap
      width: 30px
    .right
      flex-grow: 1
      .name
        font-size: 20px
        color: #1772f8
      .content
        font-size: 40px
        line-height: 48px
        font-weight: 800
        color: #F5841F
      .description
        margin-top: 16px
        font-size: 16px
        font-weight: 500
        line-height: 20px
        opacity: 0.8
        height: 26px
      .author
        padding-top: 20px
      .contract
        width: 500px
        margin-top: 25px
        border: 1px solid #efefef
        background-color: #f7f7f7
        border-radius: 16px
        .title,.address
          padding: 12px
        .address
          padding-top: 0
          color: #31373D
          font-size: 16px
          line-height: 20px
          font-weight: 700
      .author
        .creator-title,.chain-title
          padding: 5px 0
          opacity: 0.8
        .creator-name
          color: #31373D
          font-weight: 700
          padding-left: 12px
        .chain-name
          line-height: 40px
          margin-left: 5px
      .price
        padding-top: 25px
        .title
          opacity: 0.6
        .amount
          font-weight: 700
          font-size: 24px
      .buy
        margin: 10px 0
        width: 500px
        margin-bottom: 20px
.transfer,.collections
  margin-top: 40px
  font-size: 36px
  font-weight: 700
.transfer,.collections
  padding-bottom: 20px
</style>