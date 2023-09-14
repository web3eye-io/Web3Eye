<template>
  <div class="outer-bg">
    <div class="outer-container">
      <div class="top row">
        <div class="left">

        </div>
        <div class="right col">
          <div class="name">
            Coolman's Universe
          </div>
          <div class="content">
            Coolman's Universe #8149
          </div>
          <div class="description">
            Spesh is looking for his best friend throughout Coolman's Universe. To travel through this universe, Spesh uses a surfboard and a magical compass, and find...
          </div>
          <div class="author row justify-between">
            <div class="column">
              <div class="creator-title">Creator</div>
              <div class="row items-center">
                <q-avatar size="40px">
                  <img src="https://cdn.quasar.dev/img/avatar.png">
                </q-avatar> 
                <div class="creator-name">
                  @CoolmansUniverseDeployer
                </div>
              </div>
            </div>
            <div class="column">
              <div class="chain-title">Blockchain</div>
              <div class="row items-center justify-center">
                  <q-icon name="img:icons/ethereum-eth-logo.png" />
                  <div class="chain-name">Ethereum</div>
              </div>
            </div>
            <div class="col-2"></div>
          </div>
          <div class="contract column">
              <div class="title">Contract</div>
              <div class="address">0xe525FAE3fC6fBB23Af05E54Ff413613A6573CFf2</div>
          </div>
          <div class="price column">
            <div class="title">Best Price</div>
            <div class="amount">0.03ETH</div>
          </div>
          <q-btn class="buy" disable unelevated rounded color="primary" label="BUY NOW" />
        </div>
      </div>

    </div>
  </div>
</template>
<script lang='ts' setup>
</script>

<style lang="sass" scoped>
  .top
    height: 620px
    .left
      width: 620px
    .right
      .name
        font-size: 20px
        color: #1772f8
      .content
        font-size: 40px
        line-height: 48px
        font-weight: 800
        color: #F5841F
      .description
        margin-top: 16px
        font-size: 16px
        font-weight: 500
        line-height: 20px
        opacity: 0.8
      .author
        padding-top: 40px
      .contract
        margin-top: 25px
        border: 1px solid #efefef
        background-color: #f7f7f7
        border-radius: 16px
        .title,.address
          padding: 12px
        .address
          padding-top: 0
          color: #31373D
          font-size: 16px
          line-height: 20px
          font-weight: 700
      .author
        gap: 50px
        .creator-title,.chain-title
          padding: 5px 0
          opacity: 0.8
        .creator-name
          color: #31373D
          font-weight: 700
          padding-left: 12px
        .chain-name
          line-height: 40px
          margin-left: 5px
      .price
        padding-top: 25px
        .title
          opacity: 0.6
        .amount
          font-weight: 700
          font-size: 24px
      .buy
        margin: 10px 0
        width: 100%
</style>