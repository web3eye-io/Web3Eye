<template>
  <div id="contract">
    <div class="row justify-center q-pa-md">
      <div class="left col-3 self-end">
        <q-img
          src="https://picsum.photos/500/300"
          height="100%"
          width="90%"
          class="transfer-img"
          fit="fill"
        />
      </div>
      <div class="right flex column col">
        <div class="header">#1562</div>
        <div class="title">Milady 12</div>
        <div class="contract">
          <span>Contract</span>
          <span>0x8eF4459129029Db6a90a021762247dDe75899074</span>
        </div>
        <div class="chain">Ethereum @ Goerli</div>
      </div>
    </div>
    <div class="contracts q-pa-md">
      <h5>Tokens</h5>
      <div class="inner row">
        <div class="box column" v-for="index in [1,23,4,5,6,7,8,9,10]" :key="index">
          <q-img
            src="https://picsum.photos/500/300"
            height="180px"
            width="180px"
            fit="fill"
          />
          <div class="content">
            <div class="line row justify-between">
              <span class="title">#756241...</span>
              <span class="fee">4.75</span>
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