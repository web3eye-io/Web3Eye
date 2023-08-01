<template>
  <div id="transfer">
    <div class="row justify-center q-pa-md">
      <div class="left col-3 self-end">
        <q-img
          src="https://picsum.photos/500/300"
          height="100%"
          width="100%"
          class="transfer-img"
          fit="fill"
        />
      </div>
      <div class="right flex column col">
        <div class="header">#1562</div>
        <div class="title">Milady 12</div>
        <div class="description">
          Milady Maker is a collection of 10,000 generative pfpNFT's in a
          neochibi aesthetic inspired by street style tribes.
        </div>
        <div class="url">
          <span>ImageURL</span>
          <span>https://www.miladymaker.net/milady/12.png</span>
        </div>
        <div class="ipfs">
          <span>IPFS-URL</span>
          <span>IPFS://qmdjk8kfwacmt4fpxedtqax9bvsyxdy5xnr5rzwq46f3ip/teddies_hidden.jpg</span>
        </div>
        <div class="contract">
          <span>Contract</span>
          <span>0x8eF4459129029Db6a90a021762247dDe75899074</span>
        </div>
        <div class="chain">Ethereum @ Goerli</div>
      </div>
    </div>
    <div class="transfer-table">
      <div class="q-pa-md">
        <h5>Transfer</h5>
        <q-table :rows="rows" row-key="Block" flat bordered />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
const rows = [
  {
    Block: '91504562',
    Time: '2019/07/10 14:56',
    Value: '0.01 ETH',
    From: '0x8eF4459129029Db6a90a021762247dDe75899074',
    To: '0x8eF4459129029Db6a90a021762247dDe75899074',
  },
  {
    Block: '91856432',
    Time: '2019/07/10 14:56',
    Value: '0.01 ETH',
    From: '0x8eF4459129029Db6a90a021762247dDe75899074',
    To: '0x8eF4459129029Db6a90a021762247dDe75899074',
  },
  {
    Block: '918751514',
    Time: '2019/07/10 14:56',
    Value: '0.01 ETH',
    From: '0x8eF4459129029Db6a90a021762247dDe75899074',
    To: '0x8eF4459129029Db6a90a021762247dDe75899074',
  },
];
</script>
<style lang="sass" scoped>
#transfer
  width: 80%
  margin: 0 auto
  padding: 10px
  padding-left: 0
  .right
    padding: 0 15px 15px 15px
    .header,.title
      font-weight: bolder
      font-size: 16px
      padding: 5px 0
    .description
        padding: 20px 0
    .url,.ipfs,.contract
      padding: 2px 0
      span
        padding-right: 20px
    .chain
      padding-top: 20px
            
</style>
