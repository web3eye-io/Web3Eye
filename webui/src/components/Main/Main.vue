<template>
  <div view='lHh Lpr lF' class="main-container">
    <div class="summary">
      <div class="easier row items-center">
        <q-space />
        <div class="left" />
        <div class="center">Here we make web3.0 easier</div>
        <div class="right" />
        <q-space />
      </div>
      <div class="name">
        <q-img :src='largelogo' class='logo' fit="contain" />
      </div>
      <div class="mission column items-center">
        <div>Our mission is to organize the web3.0 information of different blockchain / ecosystem </div>
        <div>and make the web3.0 easier to the whole world.</div>
      </div>
    </div>
    <div class="column input-container">
      <SearchBox />
    </div>
  </div>
</template>

<script setup lang='ts'>
import { defineAsyncComponent } from 'vue';
import largelogo from '../../assets/logo/large-logo.png'
const SearchBox = defineAsyncComponent(() => import('src/components/Main/SearchBox.vue'))

</script>

<style lang='sass' scoped>
.main-container
  flex-grow: 1
.summary
  width: 840px
  margin: 0 auto
  .easier
    font-size: 16px
    color: #1772F8

    .left,.right
      width: 100px
      height: 4px
      border-radius: 2px
    .left
      background: linear-gradient(to left, transparent 0, #3187FF 0%, transparent 100%)
    .right
      background: linear-gradient(to right, transparent 0, #3187FF 0%, transparent 100%)
    .center
      padding: 0 15px
  .name 
    font-size: 8rem
  .mission
    font-size: 20px
    opacity: 0.8
.logo
  margin: 10px 0 20px 0

.icontainer,.upload-box
  width: 850px
.icontainer
  margin-top: 50px
.looking
  margin: 10px 0 10px 0
  color: $grey-8

.occupier
  height: 240px

.input-padding
  padding-bottom: 5px

.upload-box
  background: none
  flex-direction: row
.search-container
  ::v-deep > div.q-uploader
    width: auto
    max-height: 160px
.q-uploader 
  ::v-deep .bg-white
    background: none !important
.upload,.input-container
  margin: 0 auto
  width: 940px
</style>
