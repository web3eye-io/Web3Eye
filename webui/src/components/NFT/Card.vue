<template>
  <div class="row container">
    <div class="col-2 left">
      <q-img src="https://picsum.photos/500/300" height="100%" fit="fill" class="rounded-borders" />
    </div>
    <div class="col flex column center">
      <div class="content col">
        <div class="line-top">
          <span class="distance">Distance: 0.12</span>
          <span class="block1">Block: 19180569</span>
        </div>
        <div class="clone">
          <span>CloneX #4643</span>
        </div>
        <div class="total-transfers">27 transfers</div>
        <div class="contract">
          <span>Contract: 0xe525FAE3fC6fBB23Af05E54Ff413613A6573CFf2</span>
        </div>
      </div>
      <div class="transfers col flex">
        <div class="col-9 flex">
          <q-img 
            v-for="index in [1,2,3,4]" 
            :key="index" 
            src="https://picsum.photos/500/300" 
            height="100%" 
            width="120px"
            class="transfer-img rounded-borders"
            fit="fill" 
          />
        </div>
        <div class="col-1 self-center">
          ... have 21 items transfer
        </div>
        <!-- <CardSmall /> -->
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
</template>
<script lang="ts" setup>
</script>
<style lang="sass" scoped>
.container
  width: 60%
  margin:  0 auto
  height: 230px
  border: 1px solid #DDDDDD
  border-radius: 4px
  .center
    padding: 10px 10px
  .transfer-img
    margin-right: 10px
  .content
    .line-top
      .distance,.block1
        font-weight: bolder
        .block1
          padding-left: 15px
      .clone
        padding: 10px 0
.right
    height: 100%
    .right-top, .right-bottom
      padding: 10px 10px 10px 0
    .right-top
      .name
        font-weight: 700px
        color: #7D7D7D
</style>
