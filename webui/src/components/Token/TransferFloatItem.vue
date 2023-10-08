<template>
    <div class="inner row">
        <div class="inner-left column">
            <div class="title">Value</div>
            <template v-for='(item, index) in offers' :key='index'>
                <div class="token-number">#{{ item.TokenID }}</div>
                <q-separator />
                <div class="coin">+ {{ item.AmountStr }}</div>
            </template>
        </div>
        <div class="inner-center column justify-center">
            <div class="round">
                <q-icon name="img:icons/transfer.png" size="24px" style="margin-top: 11px" />
            </div>
        </div>
        <div class="inner-right column">
            <div class="title">Value</div>
            <template v-for='(item, index) in targets' :key='index'>
                <div class="token-number">#{{ item.TokenID }}</div>
                <q-separator />
                <div class="coin">+ {{ item.AmountStr }}</div>
            </template>
        </div>
    </div>
</template>

<script lang="ts" setup>
import {  computed, ref, toRef } from 'vue'
import { OrderItem } from 'src/teststore/transfer/types'

interface Props {
  offerItems: OrderItem[];
  targetItems: OrderItem[];
}

const props = defineProps<Props>()
const _offerItems = toRef(props, 'offerItems')
const _targetItems = toRef(props, 'targetItems')

const offers = computed(() => _offerItems.value)
const targets = computed(() => _targetItems.value)

</script>

<style lang="sass" scoped>
.inner
  width: 96%
  margin: 0 auto
  .token-number,.coin
    color: #1D1E20
    font-size: 16px
    font-weight: 700
    line-height: 20px
    padding: 5px 0
.inner-center
  .round
    margin: 0 20px
    height: 48px
    width: 48px
    border-radius: 50%
    background-color: #efefef
    text-align: center
</style>