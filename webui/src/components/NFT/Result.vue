<template>
  <div class='nft-container'>
    <q-splitter v-model='splitterModel'>
      <template v-slot:before>
        <!-- left -->
        <div class='q-pa-md'>
          <div class='q-col-gutter-md row items-start'>
            <div class='col-12'>
              <q-img :src='currentImg'>
                <!-- <div class='absolute-bottom text-subtitle1 text-center'>
                  Distance: 0
                </div> -->
              </q-img>
            </div>
          </div>
          <div class='content-container column'>
            <h5>目标标题</h5>
            <div>
              <span class='label'>首发时间:</span>
              <span class='value'> XXXX-XX-XX</span>
            </div>
            <div>
              <span class='label'>公链信息:</span>
              <span class='value'> 由XX发布于XXXX公链</span>
            </div>
            <div>
              <span class='label'>网站信息:</span>
              <span class='value'> 由XX上传于XXXX网站</span>
            </div>
            <div>
              <span class='label'>作者:</span>
              <span class='value'> XXX</span>
            </div>
            <div>
              <span class='label'>描述:</span>
              <span class='value'> XXX</span>
            </div>
            <div>
              <span class='label'>版权声明:</span>
              <span class='value'> XXX</span>
            </div>
            <div>
              <span class='label'>稀缺度:</span>
              <span class='value'> 1%</span>
            </div>
          </div>
        </div>
      </template>
      <template v-slot:after>
        <!-- right -->
        <div class='q-pa-md'>
          <q-timeline color='secondary' v-for='nft in nfts' :key='nft.ID'>
            <q-timeline-entry subtitle='February 22, 1986'>
              <div class="row">
                <div class='col-md-2'>
                  <q-icon
                    v-if='nft?.ImageURL?.startsWith("img")'
                    size='300px' 
                    :name='nft.ImageURL' 
                  />
                  <q-img
                    v-else
                    :src="nft.ImageURL"
                    spinner-color="red"
                  />
                </div>
                
                <div class='column col-md-9' style="margin-left: 15px;">
                  <div>
                    <span class='label'>相似度:</span>
                    <span class='value'>&nbsp; {{ nft.Distance }}</span>
                  </div>
                  <div>
                    <span class='label'>发布时间:</span>
                    <span class='value'> XXX</span>
                  </div>
                  <div>
                    <span class='label'>公链信息:</span>
                    <span class='value'> {{ nft.ChainType }}</span>
                  </div>
                  <div>
                    <span class='label'>网站信息:</span>
                    <span class='value'> {{ nft.Contract }}</span>
                  </div>
                <div>
                    <span class='label'>作者:</span>
                    <span class='value'> XXX</span>
                  </div>
                </div>
              </div>
            </q-timeline-entry>
            
          </q-timeline>
      </div>
      </template>
    </q-splitter>
    <!-- <div class="transaction">
      <h5>目标交易历史</h5>
    </div> -->
  </div>
</template>
<script lang='ts' setup>
import { useNFTMetaStore } from 'src/localstore/nft';
import { NFTMeta } from 'src/localstore/nft/types';
import { useRetrieveStore } from 'src/teststore/retrieve';
import { computed, ref, watch } from 'vue';

const splitterModel = ref(40)

const nft = useNFTMetaStore()
const nfts = computed(() => {
  const rows = [] as Array<NFTMeta>
  nft.NTFMetas.NTFMetas.forEach((el) => {
    if(el.ImageURL?.startsWith('ipfs://')){
      el.ImageURL = el.ImageURL.replace('ipfs://', 'https://ipfs.io/ipfs/')
    }
    if (el.ImageURL?.startsWith('data:image')) {
      el.ImageURL = `img:${el.ImageURL}`
    }
    rows.push(el)
  })

  rows.sort((a, b) => a.Distance > b.Distance ? 1 : -1)
  return rows
})

const currentImg = computed(() => nft.NTFMetas.Current)

const retrieve = useRetrieveStore()
const retrieves = computed(() => retrieve.Retrieves.Retrieves)

enum ImageState {
  Normal = 'Normal',
  IPFS = 'IPFS',
  Retrieving = 'Retrieving',
  WaitRecover = 'WaitRecover'
}

const currentRetrieveNFTMeta = ref({} as NFTMeta)
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const checkImage = computed(() => (row: NFTMeta) => {
  return () => {
    const image = new Image()
    image.src = row.ImageURL
    if (image.width > 0 && image.height > 0) {
      return ImageState.Normal
    }

    image.src = row.IPFSImageURL
    if (image.width > 0 && image.height > 0) {
      return ImageState.IPFS
    }

    if (row.ImageSnapshotID?.length > 0) {
      currentRetrieveNFTMeta.value = row // trigger watch

      const _row = retrieves.value?.find((el) => el.ChainType === row.ChainType && el.ChainID === row.ChainID && el.Contract === row.Contract && el.TokenID === row.TokenID)
      if (!_row) {
        console.log('not found')
        return ImageState.WaitRecover
      }
      if (_row.RetrieveState?.length > 0) {
        return ImageState.Retrieving
      }
    }
    return ImageState.WaitRecover
  }
})


watch(currentRetrieveNFTMeta, () => {
  statRetrieve(currentRetrieveNFTMeta.value)
})

const statRetrieve = (row: NFTMeta) => {
  retrieve.statRetrieve({
    ChainType: row.ChainType,
    ChainID: row.ChainID,
    Contract: row.Contract,
    TokenID: row.TokenID,
    Message: {}
  }, () => {
    // TODO
  })
}

</script>
<style lang='sass' scoped>
.nft-container
  ::v-deep .q-splitter--vertical > .q-splitter__separator
    width: 0px
  ::v-deep .q-timeline
    margin-top: 0
  .nft-result
    margin-left: 15px
  
  .content-container h5
    margin: 20px 0 6px 0
  .content-container div
    height: 30px
    line-height: 30px           
</style>