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
                  <q-img
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
import { computed, ref } from 'vue';

const splitterModel = ref(40)

const nft = useNFTMetaStore()
const nfts = computed(() => nft.NTFMetas.NTFMetas)
const currentImg = computed(() => nft.NTFMetas.Current)
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