<template>
  <div id="token">
    <div class="row box" v-for="token in tokens" :key="token.ID">
      <div class="col-2 left">
        <MyImage :url="token.ImageURL" :height="'100%'" />
      </div>
      <div class="col flex column center">
        <div class="content col">
          <div class="line-top">
            <span class="distance">Distance: {{ token.Distance }}</span>
            <span class="block1">Block: {{ token.VectorID }}</span>
          </div>
          <div class="clone">
            <span>{{ token.Name }}</span>
          </div>
          <div class="total-transfers">{{token.SiblingTokens?.length}} transfers</div>
          <div class="contract">
            <span>Contract: {{ token.Contract }}</span>
          </div>
        </div>
        <div class="transfers col flex">
          <div class="col-9" v-for="item in token.SiblingTokens?.slice(0, 5)" :key="item.ID">
            <MyImage :url="item.ImageURL" :height="'100%'" :width="'120px'"/>
          </div>
          <div class="col-1 self-center">
            ... have {{token.SiblingTokens?.length}} items transfer
          </div>
        </div>
      </div>
      <div class="col-2">
        <div class="right column justify-between">
          <div class="right-top self-end">
            <span class="name">Ethereum</span>
            <span class="net">@mainnet</span>
          </div>
          <div class="right-bottom self-end">
            <span>ERC-721</span>
            <span></span>2021/09/6 23:56
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useTokenStore } from 'src/teststore/token';
import { computed, defineAsyncComponent } from 'vue';

const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

const token = useTokenStore()
const tokens = computed(() => {
  const rows = token.SearchTokens.SearchTokens
  rows.sort((a, b) => a.Distance > b.Distance ? 1 : -1)
  return rows
})

</script>
<style lang="sass" scoped>
#token
  width: 60%
  margin:  0 auto
  padding-top: 30px
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
        .clone
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
</style>
