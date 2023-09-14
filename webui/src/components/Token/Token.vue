<template>
  <div class="outer-bg">
    <div class="outer-container">
      <div class="token row nowrap">
        <div class="left">
          <q-list bordered class="rounded-borders">
            <q-expansion-item expand-separator default-opened label="Chains">
              <q-card>
                <q-card-section>
                  <q-option-group v-model="group" :options="options" color="blue" type="checkbox">
                    <template v-slot:label="row">
                      <div class="row justify-between">
                        <div>{{ row.label }}</div>
                        <q-badge color="blue" outline rounded text-color="black" :label="row.amount" />
                      </div>
                    </template>
                  </q-option-group>
                </q-card-section>
              </q-card>
            </q-expansion-item>
          </q-list>
        </div>
        <div class="right">
          <div class="title">Collections</div>
          <div class="row box" v-for="token in tokens" :key="token.ID">
            <div class="content-left" @click="onImageClick">
              <MyImage :url="token.ImageURL" :height="'230px'" :width="'230px'" />
            </div>
            <div class="content-right column">
              <div class="line-top row space-between items-center">
                <div class="distance">Distance: {{ token.Distance }}</div>
                <div class="block1">Block: {{ token.SiblingsNum }}</div>
                <q-space />
                <div>
                  <q-icon v-if="token.ChainType === ChainType.Ethereum" name="img:icons/ethereum-eth-logo.png" />
                  <q-icon v-if="token.ChainType === ChainType.Solana" name="img:icons/solana-sol-logo.png" />
                </div>
                <div class="chain-logo">{{ token.ChainType }}</div>
              </div>
              <div class="name">
                <span>{{ token.Name }}</span>
              </div>
              <div class="contract row">
                <a href="#" @click.prevent @click="onContractClick(token)">
                  <span>Contract: {{ token.Contract }}</span>
                </a>
                <div class="copy">
                  <q-img :src='copy' class='logo' width="14px" height="14px" @click="onCopyClick(token)" />
                </div>
              </div>
              <div class="total-transfers">
                  <span>Transfers: {{ token?.TransfersNum }}</span>
              </div>
              <div class="transfers row">
                <div v-for="item in token.SiblingTokens" :key="item.ID" @click="onImageClick" class="split-token">
                  <MyImage :url="item.ImageURL" :height="'70px'" :width="'70px'" :title="item.TokenID" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <q-dialog v-model="showing" id="transfer-card">
    <q-card style="width: 860px;">
      <TransferCard :transfers="targetTransfers" :token="target" />
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
import copy from '../../assets/material/copy.png'
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))
const TransferCard = defineAsyncComponent(() => import('src/components/Transfer/Transfer.vue'))
import { copyToClipboard } from 'quasar'

const group = ref(['op1'])
const options = ref(
  [
    {
      label: 'Ethereum',
      value: 'op1',
      amount: 100,
    },
    {
      label: 'Flow',
      value: 'op2',
      amount: 10,
    },
    {
      label: 'Tezos',
      value: 'op3',
      amount: 21,
    }
  ]
)


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
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const onTransferClick = (token: SearchToken) => {
  target.value = { ...token }
  showing.value = true
}

const onImageClick = () => {
  void router.push({
    path: '/token/detail',
  })
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
  }, (error: boolean, rows: Transfer[]) => {
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
      contract: token.Contract,
      tokenID: token.TokenID
    }
  })
}

const getTokens = (page: number) => {
  token.getTokens({
    StorageKey: token.SearchTokens.StorageKey,
    Page: page,
    Message: {}
  }, (error: boolean) => {
    if (error || page >= token.SearchTokens.TotalPages) return
    page += 1
    getTokens(page)
  })
}

const onCopyClick = (token: SearchToken) => {
  void copyToClipboard(token.Contract)
}

onMounted(() => {
  if (token.SearchTokens.SearchTokens.length < token.SearchTokens.Total) {
    getTokens(2)
  }
})
</script>
<style lang="sass" scoped>
.outer-container
  padding-top: 40px
.token
  background-color: $white
  .left
    width: 290px
    .rounded-borders
      border-radius: 10px
    ::v-deep .q-checkbox
      width: 100%
      .q-checkbox__label
        flex-grow: 1
  .right
    margin-left: 40px
    flex-grow: 1
    .title
      font-weight: 700
      font-size: 36px
      line-height: 33px
    .box
      height: 230px
      border: 1px solid #EFEFEF
      border-radius: 4px
      margin-top: 40px
      background-color: #fcfbfb
      border-radius: 16px
      .content-left
        position: relative
        cursor: pointer
        top: -20px
        left: 20px
      .content-right
        flex-grow: 1
        padding-left: 40px
        opacity: 0.9
        .line-top
          padding-top: 15px
          .block1
            padding-left: 30px
          .chain-logo
            padding-right: 20px
        .name
          color: #F5841F
          font-size: 20px
          min-height: 28px
        .contract,.total-transfers
          padding: 4px 0
          a
            color: inherit
            text-decoration: none
          .copy
            padding: 0 5px
        .transfers
          padding-top: 12px
          gap: 8px
          .split-token
            cursor: pointer
@media (min-width: 600px)
.q-dialog__inner--minimized > div
  max-width: 100%

</style>
