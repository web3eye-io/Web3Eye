<template>
  <div class="outer">
    <q-layout view='lHh Lpr lFf'>
      <q-header>
        <q-toolbar>
        <div class='search row'>
          <q-img :src='logobottom' class='logo' fit="contain" @click="onLogoClick" />
        </div>
        <div class="search-box column justify-center">
          <SearchBox  v-if="displaySearchBox" />
        </div>
        <q-space />
        <a href='#/whitepaper'>White Paper</a>
        <a  href='#/deck'>Deck</a>
        <!-- <a  href='#/blog'>Blog</a>
        <a  href='#/daily'>Daily</a>
        <a  href='#/schedule'>Schedule</a> -->
        <q-btn v-if="!logined" size="md" color="primary" outline rounded label="Connect Wallet" @click="onMetaMaskClick" />
        <q-avatar v-if="logined">
              <img src="https://cdn.quasar.dev/img/boy-avatar.png">
            </q-avatar>
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
        <div class='footer'>© 2022 - web3eye.io</div>
      </q-toolbar>
    </q-footer>
  </q-layout>
  </div>
</template>

<script setup lang='ts'>
import { ref, reactive, computed, defineAsyncComponent, onMounted } from 'vue'
import { useLocalSettingStore, useWeb3jsStore } from 'src/localstore'
import { event } from 'vue-gtag'
import logobottom from '../assets/logo/logo-bottom.png'
// import metamask from '../assets/icon/metamask.webp'
import Web3 from 'web3'
import { Account } from 'src/localstore/web3js/types'
import { Cookies } from 'quasar'
import { useRouter } from 'vue-router'
import { useRoute } from 'vue-router'
const SearchBox = defineAsyncComponent(() => import('src/components/Main/SearchBox.vue'))

const setting = useLocalSettingStore()
const displaySearchBox = computed(() => setting.DisplayToolbarSearchBox)

const web3js = useWeb3jsStore()
const account = reactive({} as Account)
let web3 = new Web3(window.ethereum)

const login = ref(false)
const logined = computed(() => {
  if(!login.value) {
    if (Cookies.get('Logined')) {
      return true
    }
  }
  return false
})
const onMetaMaskClick = () => {
  web3.eth.requestAccounts((_, accounts) => {
    if (accounts?.length > 0) {
      account.Address = accounts[0]
    }
  })
  .then((result) => {
    console.log('result: ', result)
    login.value = true
    Cookies.set('Logined', 'true')
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

const router = useRouter()
const onLogoClick = () => {
  void router.push({path: '/'})
}

interface FromChannel {
  channel: string
}

const route = useRoute()
const sendChannel = () => {
  const from = JSON.parse(JSON.stringify(route.query)) as FromChannel 
  if (from.channel?.length === 0) return
  event('channel', {'channel': from.channel})
}

onMounted(() => {
  sendChannel()
})
</script>

<style scoped lang='sass'>
.outer
  background-color:  $white
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
  .q-header, .q-footer
    background: linear-gradient(to right, transparent 0, #3187FF 0%, transparent 0%)
    color: $light-black
    height: 48px
    line-height: 48px
  .q-header
    width: 90%
    margin: 0 auto
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
  line-height: 56px
  cursor: pointer
  @media (max-width: 660px)
    display: none

.q-page-container
  padding-top: 10px !important
  ::v-deep .justify-evenly
    justify-content: center
    min-height: 800px !important

.search-box
  width: 400px
  height: 50px
  ::v-deep .upload
    margin-top: 0
    height: 40px
    margin-top: 45px
  ::v-deep .search
    top: -55px

.search
  height: 56px
  padding: 8px
</style>
