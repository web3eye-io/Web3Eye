<template>
  <div id="contract">
    <div class="row justify-center q-pa-md">
      <div class="left col-3 self-end">
        <MyImage :url="current.ProfileURL" :height="'100%'" :width="'90%'" />
      </div>
      <div class="right flex column col">
        <div class="header">{{ current.Name }}</div>
        <div class="title">Milady 12</div>
        <div class="contract">
          <span>Contract</span>
          <span>{{current.Address}}</span>
        </div>
        <div class="chain">{{current.Name}} @ {{current.Creator}}</div>
      </div>
    </div>
    <div class="contracts q-pa-md">
      <h5>Tokens</h5>
      <div class="inner row">
        <div class="box column" v-for="token in tokens" :key="token.ID">
          <MyImage :url="current.ProfileURL" :height="'180px'" :width="'180px'" />
          <div class="content">
            <div class="line row justify-between">
              <span class="title">{{token.TokenID}}.</span>
              <span class="fee">{{token.ImageSnapshotID}}</span>
            </div>
            <div class="super row justify-between">
              <span>CloneX Super</span>
              <span>27</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useContractStore } from 'src/teststore/contract'
import { computed, defineAsyncComponent } from 'vue';

const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

const contract = useContractStore()
const tokens = computed(() => contract.ShotTokens.ShotTokens)
const current = computed(() => contract.Contract)
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

</style>