<template>
  <div id="contract">
    <div class="row justify-center">
      <div class="left self-end">
        <q-carousel
          animated
          v-model="slide"
          infinite
          autoplay
          arrows
          transition-prev="slide-right"
          transition-next="slide-left"
          style="height: 220px;width: 180px;"
        >
          <q-carousel-slide  
            v-for="(token,index) in tokens" 
            :key="token.ID"
            :img-src="getImageUrl(token.ImageURL)"
            :name="index"  
          />
        </q-carousel>
        <!-- <MyImage :url="current.ProfileURL" :height="'100%'" :width="'90%'" /> -->
      </div>
      <div class="right flex column col">
        <div class="header">{{ current.Name }}</div>
        <!-- <div class="title">Milady 12</div> -->
        <div class="contract">
          <span>Contract</span>
          <span>{{current.Address}}</span>
        </div>
        <div class="chain">
          <span>Name</span>
          {{current.Name}}
        </div>
      </div>
    </div>
    <div class="contracts q-pa-md">
      <h5>Tokens</h5>
      <div class="inner row">
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
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useContractStore } from 'src/teststore/contract'
import { computed, defineAsyncComponent, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

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
const slide = ref(1)
</script>
<style lang="sass" scoped>
#contract
  width: 60%
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