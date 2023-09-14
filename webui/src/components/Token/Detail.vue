<template>
  <div class="outer-bg">
    <div class="outer-container">
      <div class="top row">
        <div class="left">
          <MyImage 
            :url="'https://ipfs.io/ipfs/QmR9sexEQLMxVNzjDpYphXKmi2cACfzdCM1afXh6e6cDL4/1109.png'" 
            :height="'460px'" 
            :width="'460px'"
        />
        </div>
        <div class="right col justify-between">
          <div class="name">
            Coolman's Universe
          </div>
          <div class="content">
            Coolman's Universe #8149
          </div>
          <div class="description">
            Spesh is looking for his best friend throughout Coolman's Universe. To travel through this universe, Spesh uses a surfboard and a magical compass, and find...
          </div>
          <div class="author row justify-between">
            <div class="column">
              <div class="creator-title">Creator</div>
              <div class="row items-center">
                <q-avatar size="40px">
                  <img src="https://cdn.quasar.dev/img/avatar.png">
                </q-avatar> 
                <div class="creator-name">
                  @CoolmansUniverseDeployer
                </div>
              </div>
            </div>
            <div class="column">
              <div class="chain-title">Blockchain</div>
              <div class="row items-center justify-center">
                  <q-icon name="img:icons/ethereum-eth-logo.png" />
                  <div class="chain-name">Ethereum</div>
              </div>
            </div>
            <div class="col-2"></div>
          </div>
          <div class="contract column">
              <div class="title">Contract</div>
              <div class="address">0xe525FAE3fC6fBB23Af05E54Ff413613A6573CFf2</div>
          </div>
          <div class="price column">
            <div class="title">Best Price</div>
            <div class="amount">0.03ETH</div>
          </div>
          <q-btn class="buy" disable unelevated rounded color="primary" label="BUY NOW" />
        </div>
      </div>
      <div class="transfer">Transfer</div>
      <q-table
        row-key="Block" 
        flat 
        bordered
        :columns="(columns as any)"
        :rows="transfers"
      />
      <div class="collections">More from this collection</div>
      <div class="inner grid-container">
          <div class="box" v-for="token in tokens" :key="token.ID">
            <TokenCard :token="token" />
          </div>
        </div>
    </div>
  </div>
</template>
<script lang='ts' setup>
import { useContractStore } from 'src/teststore/contract';
import { useTransferStore } from 'src/teststore/transfer';
import { Transfer } from 'src/teststore/transfer/types';
import { computed, defineAsyncComponent } from 'vue';
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))
const TokenCard = defineAsyncComponent(() => import('src/components/Token/TokenCard.vue'))

const transfer = useTransferStore()
const transfers = computed(() => transfer.Transfers.Transfers)

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

const contract = useContractStore()
const tokens = computed(() => contract.ShotTokens.ShotTokens)


const current = computed(() => contract.Contract)
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const getTransfers = (offset: number, limit: number) => {
  transfer.getTransfers({
    ChainType: current.value.ChainType,
    ChainID: current.value.ChainID,
    Contract: current.value.Address,
    TokenID: '',
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
  .top
    .left
      width: 500px
      height: 500px
      border: 1px solid #f7f7f7
      border-radius: 16px
      padding-left: 20px
      padding-top: 20px
      box-shadow: 5px 5px 5px #f7f7f7
    .right
      margin-left: 30px
      padding-top: 20px
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
        padding-top: 40px
      .contract
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
        gap: 50px
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
        width: 100%
.transfer,.collections
  margin-top: 40px
  font-size: 36px
  font-weight: 700
.grid-container
  margin-top: 20px
  display: grid
  grid-template-columns: repeat(auto-fill, minmax(auto, 220px))
  grid-gap: 12px  
  justify-content: space-between
</style>