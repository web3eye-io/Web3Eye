<template>
  <div class="box" @click="onTokenClick">
    <div class="box-img">
        <MyImage :url="token.ImageURL" :height="'200px'" :width="'202px'" />
    </div>
    <div class="content">
        <div class="line row justify-between">
            <span class="title">#{{ token.TokenID }}</span>
            <div class="row fee">
                <span></span>
                <q-icon name="img:icons/ethereum-eth-logo.png" style="padding-top: 3px;" />
            </div>
        </div>
        <div class="super row justify-between">
            <div class="token-name">{{ token.Name }}</div>
            <div class="transfers">
                {{ token.TransfersNum }}
                <q-icon name="img:icons/transfers.png" />
            </div>
        </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ShotToken } from 'src/teststore/contract/types';
import { defineAsyncComponent, toRef } from 'vue';
const MyImage = defineAsyncComponent(() => import('src/components/Token/Image.vue'))

interface Props {
  token: ShotToken;
}

const props = defineProps<Props>()
const token = toRef(props, 'token')

const onTokenClick = () => {
  // TODO
}

</script>
<style lang="sass" scoped>
.box
  width: 220px
  max-width: 280px
  height: 290px
  border-radius: 10px
  border: 1px solid #efefef
  background-color: #FAFAFA
  &:hover
    cursor: pointer
  .box-img
    padding: 8px
  .content
    .line
      padding: 5px
    .super
      padding: 0 5px 2px 5px
      .token-name
        max-width: 160px
        overflow: hidden
        text-overflow: ellipsis
        height: 20px
    .transfers
      min-width: 20px
</style>