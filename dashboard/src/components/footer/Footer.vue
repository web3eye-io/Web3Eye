<template>
  <q-footer>
    <div class='row announcement main'>
      <q-icon class='icon start' name='copyright' />
      {{ $t('MSG_DASHBOARD_COPYRIGHT', { THIS_YEAR: new Date().getFullYear() }) }}
    </div>
  </q-footer>
</template>

<style lang='sass' scoped>
.announcement
  color: $blue-6
</style>
