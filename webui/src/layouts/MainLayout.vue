<template>
  <div class="outer">
    <q-layout view='lHh Lpr lFf'>
      <q-header>
        <q-toolbar>
        <div class='search row'>
          <q-img :src='logobottom' class='logo' fit="contain" />
          <q-input
            v-model='search'
            rounded
            outlined
            disable
            dense
            class='search-box'
            placeholder='Coming soon'
            v-if='false'
          >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </div>
        <q-space />
        <a href='#/whitepaper'>White Paper</a>
        <a  href='#/deck'>Deck</a>
        <a  href='#/blog'>Blog</a>
        <a  href='#/daily'>Daily</a>
        <a  href='#/schedule'>Schedule</a>
        <q-btn size="md" color="primary" outline rounded label="Connect Wallet" @click="onMetaMaskClick" />

        <!-- <q-btn avatar :icon='"img:" + metamask' flat dense round size='18px'>
          <q-menu auto-close>
            <q-list>
              <q-item clickable>
                <q-item-section  @click='onMetaMaskClick'>Profile</q-item-section>
              </q-item>
              <q-item clickable>
                <q-item-section @click='onTxClick'>Transfer</q-item-section>
              </q-item>
              <q-item clickable>
                <q-item-section @click='onLogout'>Logout</q-item-section>
              </q-item>
            </q-list>
          </q-menu>
        </q-btn> -->
      </q-toolbar>
    </q-header>
  
    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer>
      <q-toolbar class="justify-center">
        <div class='footer'>© 2022 - Cyber Tracer</div>
      </q-toolbar>
    </q-footer>
  </q-layout>
  </div>
</template>

<script setup lang='ts'>
import { ref, reactive } from 'vue'
import { useWeb3jsStore } from 'src/localstore'

import logobottom from '../assets/logo/logo-bottom.png'
// import metamask from '../assets/icon/metamask.webp'
import Web3 from 'web3'
import { Account } from 'src/localstore/web3js/types'
// import { useRouter } from 'vue-router'

// const setting = useLocalSettingStore()
// const displaySearchBox = computed(() => setting.DisplayToolbarSearchBox)

const search = ref('')

const web3js = useWeb3jsStore()
const account = reactive({} as Account)
let web3 = new Web3(window.ethereum)


const onMetaMaskClick = () => {
  web3.eth.requestAccounts((_, accounts) => {
    if (accounts?.length > 0) {
      account.Address = accounts[0]
    }
  })
  .then((result) => {
    console.log('result: ', result)
    web3js.setWeb3(web3)
    void getBalance()
  })
  .catch((error) => {
    if (!window.ethereum) {
      window.location.href = 'https://metamask.io/download/'
    }
    console.log('error: ', error)
  })
}

const getBalance = async() => {
  const balance = await web3.eth.getBalance(account.Address)
  account.Balance = web3.utils.fromWei(balance, 'ether')
  void getChainID()
}

const getChainID = async() => {
  const chainID = await web3.eth.getChainId()
  account.ChainID = chainID
  console.log('ChainID: ', chainID)
  web3js.setAccount(account)
  console.log('web3: ', web3js.getAccount())
}

// const onLogout = () => {
//   // TODO
// }

// const router = useRouter()
// const onTxClick = () => {
//   void router.push({path: '/transaction'})
// }

</script>

<style scoped lang='sass'>
.outer
  background:  $white
  background-image: url(../assets/material/background.png)
  background-repeat: repeat
  content: ""
  display: block
  position: absolute
  top: 0
  right: 0
  height: 100%
  width: 100%
.q-layout
  font-size: 14px
  font-weight: 500
  color: $light-black
  font-family: 'Manrope'
  width: 90%
  margin: 0 auto
  
  .q-header, .q-footer
    background-color: inherit
    color: $light-black
    height: 48px
    line-height: 48px
  .q-header
    position: inherit
    a,button
      margin: 0 18px 0 18px
    a
      text-decoration: none
      color: #31373D
      @media (max-width: $breakpoint-sm-max)
        display: none
    button
      text-transform: none
      ::v-deep .q-btn_context
        padding: 4px 0
  .q-toolbar
    padding: 0
  .q-footer
    background-color: $white
    opacity: 0.7

.logo
  width: 120px
  margin-right: 10px
  line-height: 72px
  @media (max-width: 660px)
    display: none

.q-page-container
  ::v-deep .justify-evenly
    justify-content: center
    min-height: 800px !important

.search-box
  width: 450px
  max-width: 100%

.search
  height: 56px
  padding: 8px
</style>
