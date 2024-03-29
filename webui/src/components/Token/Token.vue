<template>
  <div class="token-container">
    <div class="token row wrap">
      <div class="left">
        <q-list bordered class="rounded-borders">
          <q-expansion-item expand-separator default-opened label="Chains">
            <q-card>
              <q-card-section>
                <q-option-group v-model="groups" :options="options" color="blue" type="checkbox">
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
        <div id="tokens">
          <div class="row boxes" v-for="token in displayTokens" :key="token.ID">
            <div class="content-left" @click="onTokenClick(token)">
              <MyImage :url="token.ImageURL" :height="'230px'" :width="'230px'" />
            </div>
            <div class="content-right column">
              <div class="line-top row space-between items-center">
                <div class="distance">Distance: {{ token.Distance }}</div>
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
                  <q-img :src='copy' class='contract-copy' width="14px" height="14px" @click="onCopyClick(token)" />
                </div>
              </div>
              <div class="total-transfers">
                <span>Transfers: {{ token?.TransfersNum }}</span>
              </div>
              <div class="transfers row">
                <div v-for="item in token.SiblingTokens" :key="item.EntID" @click="onShotTokenClick(token, item)"
                  class="split-token">
                  <MyImage :url="item.ImageURL" :height="'70px'" :width="'70px'" :title="item.TokenID" />
                </div>
              </div>
            </div>
          </div>
          <div class="loading">
            <q-inner-loading :showing="loading" style="color:#b8b1b1" />
          </div>
          <div v-if="haveMore" class="no-more row">no more content</div>
          <div id="bottom" style="padding-bottom: 50px;"></div>
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
import { SearchToken, SiblingToken } from 'src/teststore/token/types'
import { Transfer } from 'src/teststore/transfer/types'
import { computed, defineAsyncComponent, onMounted, ref, watch } from 'vue'
import { ChainType } from 'src/teststore/basetypes/const'
import copy from '../../assets/material/copy.png'
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))
const TransferCard = defineAsyncComponent(() => import('src/components/Transfer/Transfer.vue'))
import { copyToClipboard } from 'quasar'
import { Notify } from 'quasar'
import { useStorageKeyStore } from 'src/localstore/storagekey';

const token = useTokenStore()
const tokens = computed(() => {
  const rows = token.SearchTokens.SearchTokens
  rows.forEach((el) => {
    if (el.ChainType === ChainType.Ethereum) {
      ethereums.value += 1
    }
    if (el.ChainType === ChainType.Solana) {
      solanas.value += 1
    }
  })
  rows.sort((a, b) => a.Distance > b.Distance ? 1 : -1)
  return rows
})

const ethereums = ref(0)
const solanas = ref(0)
const displayTokens = computed(() => tokens.value.filter((el) => {
  if (groups.value?.length === 0) {
    return true
  }
  if (groups.value?.length > 0) {
    if (el.ChainType.includes(groups.value[0])) return true
  }
  if (groups.value?.length > 1) {
    if (el.ChainType.includes(groups.value[1])) return true
  }
  return false
}))

const groups = ref([])
const options = computed(() => {
  const rows = token.SearchTokens.SearchTokens
  let ethereums = 0
  let solanas = 0
  rows.forEach((el) => {
    if (el.ChainType === ChainType.Ethereum) {
      ethereums += 1
    }
    if (el.ChainType === ChainType.Solana) {
      solanas += 1
    }
  })
  return [
    {
      label: 'Ethereum',
      value: 'Ethereum',
      amount: ethereums,
    },
    {
      label: 'Solana',
      value: 'Solana',
      amount: solanas,
    }
  ]
})

const target = ref({} as SearchToken)
const targetTransfers = ref([] as Array<Transfer>)

const showing = ref(false)

const onTokenClick = (token: SearchToken) => {
  void router.push({
    path: '/token/detail',
    query: {
      chainID: token.ChainID,
      chainType: token.ChainType,
      contract: token.Contract,
      tokenID: token.TokenID,
      id: token.ID,
    }
  })
}

const onShotTokenClick = (token: SearchToken, shotToken: SiblingToken) => {
  void router.push({
    path: '/token/detail',
    query: {
      chainID: token.ChainID,
      chainType: token.ChainType,
      contract: token.Contract,
      tokenID: shotToken.TokenID,
      id: shotToken.ID,
    }
  })
}

const router = useRouter()

const onContractClick = (token: SearchToken) => {
  void router.push({
    path: '/contract',
    query: {
      contract: token.Contract,
      chainID: token.ChainID,
      chainType: token.ChainType
    }
  })
}

const onCopyClick = (token: SearchToken) => {
  void copyToClipboard(token.Contract)
  Notify.create({
    position: 'bottom-right',
    message: 'Contract Copied',
    color: 'green',
    timeout: 2000
  })
}

const haveMore = ref(false)
const currentPage = ref(1)
const loading = ref(false)

const localkey = useStorageKeyStore()

watch(() => [tokens.value], () => {
  if (tokens.value?.length === 0) {
    haveMore.value = false
    isLoading.value = false
    loading.value = false
    currentPage.value = 1
  }
})

const loadMore = () => {
  if (localkey.getStorageKey() === null || localkey.getStorageKey() === '') {
    haveMore.value = false
    isLoading.value = false
    loading.value = true
    currentPage.value = 1
    return
  }
  if (currentPage.value > token.SearchTokens.Pages && token.SearchTokens.Pages !== 0) {
    haveMore.value = true
    return
  }
  loading.value = true
  if (currentPage.value === 1 && token.SearchTokens.SearchTokens?.length !== 0) {
    currentPage.value += 1
  }
  token.getTokens({
    Limit: 8,
    Page: currentPage.value,
    Message: {}
  }, (error: boolean) => {
    loading.value = false
    isLoading.value = false
    if (!error) {
      currentPage.value += 1
    }
  })
}

const isLoading = ref(false)
const handleObserve = (entries: IntersectionObserverEntry[]) => {
  entries.forEach((entry) => {
    if (isLoading.value) return
    if (entry.isIntersecting) {
      isLoading.value = true
      loadMore()
    }
  })
}

onMounted(() => {
  const observer = new IntersectionObserver(handleObserve)
  const target = document.getElementById('bottom')
  observer.observe(target as Element)
})

</script>
<style lang="sass" scoped>
.token-container
  background: $white
  padding-top: 40px
.token
  background-color: $white
  height: 100vh
  .left
    width: 290px
    margin-left: 90px
    .rounded-borders
      border-radius: 10px
    ::v-deep .q-checkbox
      width: 100%
      .q-checkbox__label
        flex-grow: 1
  .right
    margin-left: 40px
    margin-right: 90px
    flex-grow: 1
    .title
      font-weight: 700
      font-size: 36px
      line-height: 33px
    .boxes
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
            cursor: pointer
        .transfers
          padding-top: 12px
          gap: 8px
          .split-token
            cursor: pointer
    .loading
      ::v-deep .absolute-full
        position: relative
        top: auto
    .no-more
      padding-top: 5px
      display: block
      text-align: center
      color: #a3a3a3
@media (min-width: 600px)
.q-dialog__inner--minimized > div
  max-width: 100%

</style>
