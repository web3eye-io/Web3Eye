<template>
    <div class='outer-bg'>
    <div class='outer-container'>
  <q-page class='row items-center justify-evenly'>
    <div class='column'>
      <div class='row logo'>
        <q-img :src='logo' style='width: 640px;' />
      </div>
      <q-btn flat class='btn btn1' label='Download English Version' rounded disable />
      <q-btn flat class='btn btn2' label='下载中文版' rounded disable />
      <div class='row looking'>
        <q-space />
          <q-btn label='Request Other Language Copy' rounded flat>
            <q-tooltip>
              Coming soon
            </q-tooltip>
          </q-btn>
          <q-space />
        </div>
      <div class='occupier' />
    </div>
  </q-page>
  </div>
  </div>
</template>

<script setup lang='ts'>

import logo from '../assets/logo/logo.png'

</script>

<style scoped lang='sass'>
.logo
  margin-bottom: 48px

.btn
  margin-bottom: 16px
  height: 48px
  font-size: 20px

.occupier
  height: 200px

.btn1
  background-color: rgb(4, 143, 244)
  color: white

.btn2
  border: 1px solid rgb(231, 94, 1)
  color: $grey-8
  font-weight: bold

.looking
  margin: 10px 0 10px 0
  color: $grey-8
</style>
