<template>
  <div class="q-gutter-md">
    <div>Account: {{ account.Address }}</div> 
    <div>Balance: {{ account.Balance }}</div> 
    <div>Transfer To: </div>
    <div class="q-gutter-md" style="width: 35%;">
      <q-input
        v-model='targetAddress'
        rounded
        outlined
        dense
        placeholder='address'
      />
      <q-input
        v-model.number='targetValue'
        rounded
        outlined
        dense
        type='number'
        placeholder='value'
      />
    </div>
    <q-btn color="white" text-color="black" label="转账" @click='onClick' :disable='transferring' />
    <!-- <q-btn color="white" text-color="black" label="GASPrice" @click='onGetGasPrice' /> -->
  </div>

</template>
<script setup lang='ts'>
import { useWeb3jsStore } from 'src/localstore';
import { computed, ref } from 'vue';

const web3js = useWeb3jsStore()
const account = computed(() => web3js.getAccount())
const web3 = computed(() => web3js.getWeb3())

const targetAddress = ref('')
const targetValue = ref(1)

const transferring = ref(false)

const onClick =  () => {
  console.log('Value: ', targetValue.value)
  transferring.value = true
  web3.value.eth.sendTransaction({
    from: account.value.Address,
    to: targetAddress.value?.trim(),
    value: `${targetValue.value}`
  })
  .then((result) => {
    transferring.value = false
    console.log('result: ', result)
  })
  .catch((error) => {
    transferring.value = false
    console.log('error: ', error)
  })
}

const onGetGasPrice = async() => {
  const price = await web3.value.eth.getGasPrice()
  console.log('gasPrice: ', price)
}
</script>