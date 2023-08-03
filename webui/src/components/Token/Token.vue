<template>
  <div id="token">
    <div class="top row">
      <div class="col-2">
        <MyImage :url="token?.SearchTokens?.Current" :height="'230px'" />
      </div>
    </div>
    <h5>Result</h5>
    <div class="row box" v-for="token in tokens" :key="token.ID">
      <div class="col-2 left">
        <MyImage :url="token.ImageURL" :height="'230px'" :title="token.TokenID" />
      </div>
      <div class="col flex column center">
        <div class="content col">
          <div class="line-top">
            <span class="distance">Distance: {{ token.Distance }}</span>
            <!-- <span class="block1">Block: {{ token.SiblingsNum }}</span> -->
          </div>
          <div class="name">
            <span>{{ token.Name }}</span>
          </div>
          <div class="total-transfers">
            <a href="#" @click.prevent @click="onTransferClick(token)" v-if="token?.SiblingTokens?.length > 0">{{token?.TransfersNum}} transfers</a>
          </div>
          <div class="contract">
            <a href="#" @click.prevent @click="onContractClick(token)">
              <span>Contract: {{ token.Contract }}</span>
            </a>
          </div>
        </div>
        <div class="transfers col flex">
          <div class="col-9" v-for="item in token.SiblingTokens?.slice(0, 5)" :key="item.ID">
            <MyImage :url="item.ImageURL" :height="'110px'" :width="'120px'" :title="item.TokenID" />
          </div>
          <div class="col-1 self-center" v-if="token?.SiblingTokens?.length > 5">
            ...
          </div>
        </div>
      </div>
      <div class="col-2">
        <div class="right column justify-between">
          <div class="right-top self-end">
            <span class="name">
              <q-icon v-if="token.ChainType === ChainType.Ethereum" name="img:icons/ethereum-eth-logo.png" style="padding-bottom: 3px;" />
              <q-icon v-if="token.ChainType === ChainType.Solana" name="img:icons/solana-sol-logo.png" style="padding-bottom: 2px;" />
              {{ token.ChainType }}
            </span>
            <span class="net">@mainnet</span>
          </div>
          <div class="right-bottom self-end">
            <span>{{token.TokenType}}</span>
            <span>  ChainID-{{token.ChainID}}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
  <q-dialog v-model="showing" id="transfer-card">
    <q-card style="width: 860px;">
      <TransferCard :transfers="targetTransfers" :token="target"  />
    </q-card>
  </q-dialog>
</template>
<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { useTokenStore } from 'src/teststore/token';
import { SearchToken } from 'src/teststore/token/types';
import { useTransferStore } from 'src/teststore/transfer';
import { Transfer } from 'src/teststore/transfer/types';
import { computed, defineAsyncComponent, onMounted, ref, watch } from 'vue';
import { ChainType } from 'src/teststore/basetypes/const';

const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))
const TransferCard = defineAsyncComponent(() => import('src/components/Transfer/Transfer.vue'))

const token = useTokenStore()
const tokens = computed(() => {
  const rows = token.SearchTokens.SearchTokens
  rows.sort((a, b) => a.Distance > b.Distance ? 1 : -1)
  return rows
})

const target = ref({} as SearchToken)

const transfer = useTransferStore()
const targetTransfers = ref([] as Array<Transfer>)

const showing = ref(false)
const onTransferClick = (token: SearchToken) => {
  target.value = {...token}
  showing.value = true
}

watch(() => target.value?.ID, () => {
  if (!target.value) return
  if (!transfer.getTransferByToken(target.value?.ChainID, target?.value?.ChainType, target.value?.Contract, target.value?.TokenID)) {
    getTransfers(0, 100)
  }
})

const getTransfers = (offset: number, limit: number) => {
  transfer.getTransfers({
    ChainType: target.value.ChainType,
    ChainID: target.value.ChainID,
    Contract: target.value.Contract,
    TokenID: target.value.TokenID,
    Offset: offset,
    Limit: limit,
    Message: {}
  }, (error:boolean, rows: Transfer[]) => {
    if (error || rows.length < limit) {
      targetTransfers.value = rows
      return
    }
    getTransfers(offset, offset + limit)
  })
}

const router = useRouter()

const onContractClick = (token: SearchToken) => {
  void router.push({
    path: '/contract',
    query: {
      contract: token.Contract
    }
  })
}
const getTokens = (page: number) => {
  token.getTokens({
    StorageKey: token.SearchTokens.StorageKey,
    Page: page,
    Message: {}
  }, (error:boolean) => {
    if (error || page >= token.SearchTokens.TotalPages) return
    page += 1
    getTokens(page)
  })
}

onMounted(() => {
  if (token.SearchTokens.SearchTokens.length < token.SearchTokens.Total) {
    getTokens(2)
  }
})
</script>
<style lang="sass" scoped>
#token
  width: 60%
  margin:  0 auto
  padding-top: 30px
  .top
    padding-bottom: 20px
  .box
    height: 230px
    border: 1px solid #f4eeee
    border-radius: 4px
    margin-bottom: 30px
    .center
      padding: 10px 10px
    .content
      .line-top
        .distance,.block1
          font-weight: bolder
        .block1
          padding-left: 15px
        .name
          padding: 10px 0
    .transfers div
      margin-right: 5px
    .right
        height: 100%
        .right-top, .right-bottom
          padding: 10px 10px 10px 0
        .right-top
          .name
            font-weight: 700px
            color: #7D7D7D

@media (min-width: 600px)
.q-dialog__inner--minimized > div
  max-width: 100%

</style>
