<template>
    <q-table 
        dense 
        flat 
        :title='$t("MSG_ENDPOINTS")' 
        :rows='endpoints' 
        row-key='ID' 
        :rows-per-page-options='[100]'
        @row-click='(evt, row, index) => onRowClick(row as Endpoint)' 
        selection='single'
        v-model:selected='selectedEndpoints' 
        :columns='columns'
    >
        <template #top-right>
            <div class='row indent flat'>
                <q-btn dense flat class='btn flat' :label='$t("MSG_CREATE")' @click='onCreate' />
                <q-btn dense flat class='btn flat' :label='$t("MSG_DELETE")' :disable='selectedEndpoints?.length === 0'
                    @click='onDelete(selectedEndpoints?.[0])' />
            </div>
        </template>
    </q-table>
    <q-dialog v-model='showing' @hide='onMenuHide' position='right'>
        <q-card class='popup-menu' style="width: 600px">
            <q-card-section>
                <span>{{ $t('MSG_ENDPOINT') }}</span>
            </q-card-section>
            <q-card-section>
                <q-select v-if="updating" :disable="updating"  :options='Object.keys(ChainType)' v-model='target.ChainType' :label='$t("MSG_CHAIN_TYPE")' />
                <q-input v-if="updating" :disable="updating" v-model='target.ChainID' :label='$t("MSG_CHAIN_ID")' />
                <q-input  v-model='target.Address' :label='$t("MSG_ADDRESS")' />
                <q-input  v-model='target.RPS' :label='$t("MSG_RPS")' />
                <q-select :disable="updating"  :options='Object.keys(ChainType)' v-model='target.ChainType' :label='$t("MSG_CHAIN_TYPE")' />
                <q-select v-if="updating" :options='Object.keys(EndpointState)' v-model='target.State' :label='$t("MSG_SYNC_STATE")' />
            </q-card-section>
            <q-item class='row'>
                <LoadingButton loading :label='$t("MSG_SUBMIT")' @click='onSubmit' />
                <q-btn class='btn round' :label='$t("MSG_CANCEL")' @click='onCancel' />
            </q-item>
        </q-card>
    </q-dialog>
</template>
  
<script setup lang='ts'>
import { useEndpointStore } from 'src/teststore/endpoint'
import { Endpoint } from 'src/teststore/endpoint/types'
import { NotifyType } from 'src/teststore/local/notify'
import { computed, defineAsyncComponent, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ChainType } from 'src/teststore/basetypes/const'
import { EndpointState } from 'src/teststore/basetypes/endpoint/const'
const { t } = useI18n({ useScope: 'global' })

const LoadingButton = defineAsyncComponent(() => import('src/components/button/LoadingButton.vue'))

const endpoint = useEndpointStore()
const endpoints = computed(() => endpoint.Endpoints.Endpoints)

const target = ref({} as Endpoint)

const showing = ref(false)
const updating = ref(false)

const onCreate = () => {
    showing.value = true
    updating.value = false
}

const onMenuHide = () => {
    target.value = {} as Endpoint
    showing.value = false
}

const onCancel = () => {
    onMenuHide()
}

onMounted(() => {
    if (endpoints.value?.length === 0) {
        getEndpoints(0, 100)
    }
})

const getEndpoints = (offset: number, limit: number) => {
    endpoint.getEndpoints({
        Offset: offset,
        Limit: limit,
        Message: {
            Error: {
                Title: t('MSG_GET_ENDPOINTS'),
                Message: t('MSG_GET_ENDPOINTS_FAIL'),
                Popup: true,
                Type: NotifyType.Error
            }
        }
    }, (error: boolean, rows: Array<Endpoint>) => {
        if (error || rows.length < limit) {
            return
        }
        getEndpoints(offset + limit, limit)
    })
}

const onSubmit = (done: () => void) => {
    updating.value ? updateEndpoint(done) : createEndpoint(done)
}

const onRowClick = (row: Endpoint) => {
    target.value = { ...row }
    updating.value = true
    showing.value = true
}

const updateEndpoint = (done: () => void) => {
    endpoint.updateEndpoint({
        ...target.value,
        Message: {
            Error: {
                Title: 'MSG_UPDATE_ENDPOINT',
                Message: 'MSG_UPDATE_ENDPOINT_FAIL',
                Popup: true,
                Type: NotifyType.Error
            },
            Info: {
                Title: 'MSG_UPDATE_ENDPOINT',
                Message: 'MSG_UPDATE_ENDPOINT_SUCCESS',
                Popup: true,
                Type: NotifyType.Success
            }
        }
    }, (error: boolean) => {
        done()
        if (error) {
            return
        }
        onMenuHide()
    })
}

const createEndpoint = (done: () => void) => {
    endpoint.createEndpoint({
        ...target.value,
        Message: {
            Error: {
                Title: 'MSG_CREATE_ENDPOINT',
                Message: 'MSG_CREATE_ENDPOINT_FAIL',
                Popup: true,
                Type: NotifyType.Error
            },
            Info: {
                Title: 'MSG_CREATE_ENDPOINT',
                Message: 'MSG_CREATE_ENDPOINT_SUCCESS',
                Popup: true,
                Type: NotifyType.Success
            }
        }
    }, (error: boolean) => {
        done()
        if (error) {
            return
        }
        onMenuHide()
    })
}

const selectedEndpoints = ref([] as Array<Endpoint>)
const onDelete = (row: Endpoint) => {
    endpoint.deleteEndpoint({
        ID: row.ID,
        EntID: row.EntID,
        Message: {}
    }, () => {
        // TODO
    })
}

const columns = computed(() => [
    {
        name: 'ID',
        label: t('MSG_ID'),
        sortable: true,
        field: (row: Endpoint) => row.ID
    },
    {
        name: 'EntID',
        label: t('MSG_ENT_ID'),
        sortable: true,
        field: (row: Endpoint) => row.EntID
    },
    {
        name: 'ChainID',
        label: t('MSG_CHAIN_ID'),
        sortable: true,
        field: (row: Endpoint) => row.ChainID
    },
    {
        name: 'ChainType',
        label: t('MSG_CHAIN_TYPE'),
        sortable: true,
        field: (row: Endpoint) => row.ChainType
    },
    {
        name: 'Address',
        label: t('MSG_ADDRESS'),
        sortable: true,
        field: (row: Endpoint) => row.Address
    },
    {
        name: 'State',
        label: t('MSG_STATE'),
        sortable: true,
        field: (row: Endpoint) => row.State
    },
])
</script>
  