<template>
    <q-table 
        dense 
        flat 
        :title='$t("MSG_SYNCTASKS")' 
        :rows='tasks' 
        row-key='ID' 
        :rows-per-page-options='[100]'
        @row-click='(evt, row, index) => onRowClick(row as SyncTask)' 
        selection='single'
        v-model:selected='selectedSyncTasks' 
        :columns='columns'
    >
        <template #top-right>
            <div class='row indent flat'>
                <q-btn dense flat class='btn flat' :label='$t("MSG_CREATE")' @click='onCreate' />
                <q-btn dense flat class='btn flat' :label='$t("MSG_DELETE")' :disable='selectedSyncTasks?.length === 0'
                    @click='onDelete(selectedSyncTasks?.[0])' />
            </div>
        </template>
    </q-table>
    <q-dialog v-model='showing' @hide='onMenuHide' position='right'>
        <q-card class='popup-menu' style="width: 600px">
            <q-card-section>
                <span>{{ $t('MSG_SYNCTASK') }}</span>
            </q-card-section>
            <q-card-section>
                <q-select :disable="updating"  :options='Object.keys(ChainType)' v-model='target.ChainType' :label='$t("MSG_CHAIN_TYPE")' />
                <q-input :disable="updating" v-model='target.ChainID' :label='$t("MSG_CHAIN_ID")' />
                <q-input :disable="updating" v-model='target.Start' :label='$t("MSG_START")' />
                <q-input :disable="updating" v-model='target.End' :label='$t("MSG_END")' />
                <q-select :options='Object.keys(SyncState)' v-model='target.SyncState' :label='$t("MSG_SYNC_STATE")' />
                <q-input type="textarea" v-model='target.Description' :label='$t("MSG_DESCRIPTION")' />
            </q-card-section>
            <q-item class='row'>
                <LoadingButton loading :label='$t("MSG_SUBMIT")' @click='onSubmit' />
                <q-btn class='btn round' :label='$t("MSG_CANCEL")' @click='onCancel' />
            </q-item>
        </q-card>
    </q-dialog>
</template>
  
<script setup lang='ts'>
import { ChainType } from 'src/teststore/basetypes/const'
import { SyncState } from 'src/teststore/basetypes/synctask/const'
import { NotifyType } from 'src/teststore/local/notify'
import { useSyncTaskStore } from 'src/teststore/task'
import { SyncTask } from 'src/teststore/task/types'
import { computed, defineAsyncComponent, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n({ useScope: 'global' })

const LoadingButton = defineAsyncComponent(() => import('src/components/button/LoadingButton.vue'))

const task = useSyncTaskStore()
const tasks = computed(() => task.SyncTasks.SyncTasks)

const target = ref({} as SyncTask)

const showing = ref(false)
const updating = ref(false)

const onCreate = () => {
    showing.value = true
    updating.value = false
}

const onMenuHide = () => {
    target.value = {} as SyncTask
    showing.value = false
}

const onCancel = () => {
    onMenuHide()
}

onMounted(() => {
    if (tasks.value?.length === 0) {
        getSyncTasks(0, 100)
    }
})

const getSyncTasks = (offset: number, limit: number) => {
    task.getSyncTasks({
        Offset: offset,
        Limit: limit,
        Message: {
            Error: {
                Title: t('MSG_GET_SYNCTASKS'),
                Message: t('MSG_GET_SYNCTASKS_FAIL'),
                Popup: true,
                Type: NotifyType.Error
            }
        }
    }, (error: boolean, rows: Array<SyncTask>) => {
        if (error || rows.length < limit) {
            return
        }
        getSyncTasks(offset + limit, limit)
    })
}

const onSubmit = (done: () => void) => {
    updating.value ? updateSyncTask(done) : createSyncTask(done)
}

const onRowClick = (row: SyncTask) => {
    target.value = { ...row }
    updating.value = true
    showing.value = true
}

const updateSyncTask = (done: () => void) => {
    task.updateSyncTask({
        ...target.value,
        Message: {
            Error: {
                Title: 'MSG_UPDATE_SYNCTASK',
                Message: 'MSG_UPDATE_SYNCTASK_FAIL',
                Popup: true,
                Type: NotifyType.Error
            },
            Info: {
                Title: 'MSG_UPDATE_SYNCTASK',
                Message: 'MSG_UPDATE_SYNCTASK_SUCCESS',
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

const createSyncTask = (done: () => void) => {
    task.createSyncTask({
        ...target.value,
        Message: {
            Error: {
                Title: 'MSG_CREATE_SYNCTASK',
                Message: 'MSG_CREATE_SYNCTASK_FAIL',
                Popup: true,
                Type: NotifyType.Error
            },
            Info: {
                Title: 'MSG_CREATE_SYNCTASK',
                Message: 'MSG_CREATE_SYNCTASK_SUCCESS',
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

const selectedSyncTasks = ref([] as Array<SyncTask>)
const onDelete = (row: SyncTask) => {
    task.deleteSyncTask({
        ID: row.ID,
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
        field: (row: SyncTask) => row.ID
    },
    {
        name: 'ChainID',
        label: t('MSG_CHAIN_ID'),
        sortable: true,
        field: (row: SyncTask) => row.ChainID
    },
    {
        name: 'ChainType',
        label: t('MSG_CHAIN_TYPE'),
        sortable: true,
        field: (row: SyncTask) => row.ChainType
    },
    {
        name: 'Start',
        label: t('MSG_START'),
        sortable: true,
        field: (row: SyncTask) => row.Start
    },
    {
        name: 'End',
        label: t('MSG_END'),
        sortable: true,
        field: (row: SyncTask) => row.End
    },
])
</script>
  