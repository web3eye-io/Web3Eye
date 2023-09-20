<template>
    <q-table 
        dense 
        flat 
        :title='"MSG_OVERVIEW"' 
        row-key='ID'
        :rows-per-page-options='[100]'
        :rows="[]"
        :columns='(columns as any)'
    />
</template>
<script setup lang='ts'>
import { computed } from 'vue';
import { useI18n } from 'vue-i18n'
const { t } = useI18n({ useScope: 'global' })

const columns = computed(() => [
    {
        name: 'ID',
        label: t('MSG_ID'),
        sortable: true,
        field: '',
    },
    {
        name: 'ChainID',
        label: t('MSG_CHAIN_ID'),
        sortable: true,
        field: '',
    },
    {
        name: 'ChainType',
        label: t('MSG_CHAIN_TYPE'),
        sortable: true,
        field: '',
    },
    {
        name: 'Block',
        label: t('MSG_BLOCK'),
        sortable: true,
        field: '',
    },
    {
        name: 'Transfers',
        label: t('MSG_TRANSFERS'),
        sortable: true,
        field: '',
    },
    {
        name: 'Tokens',
        label: t('MSG_TOKENS'),
        sortable: true,
        field: '',
    },
    {
        name: 'Collections',
        label: t('MSG_COLLECTIONS'),
        sortable: true,
        field: '',
    },
])
</script>