<template>
  <div>
    {{ account }}
  </div>
  <q-btn color="white" text-color="black" label="转账" @click='onClick' />
  <q-btn color="white" text-color="black" label="GASPrice" @click='onGetGasPrice' />
</template>
<script setup lang='ts'>
import { useWeb3jsStore } from 'src/localstore';
import { computed } from 'vue';

const web3js = useWeb3jsStore()
const account = computed(() => web3js.getAccount())
const web3 = computed(() => web3js.getWeb3())

const onClick =  () => {
  web3.value.eth.sendTransaction({
    from: account.value.Address,
    to: '0x656D969e412C47E7638c8a8843F86E3CaFE9a503',
    value: '3000000000000000000'
  })
  .then((result) => {
    console.log('result: ', result)
  })
  .catch((error) => {
    console.log('error: ', error)
  })
}

const onGetGasPrice = async() => {
  const price = await web3.value.eth.getGasPrice()
  console.log('gasPrice: ', price)
}
</script>