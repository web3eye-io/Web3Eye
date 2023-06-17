<template>
  <div v-for='(values, index) in snapshotsMap' :key='index'>
    <q-table
      flat bordered
      :title='values?.[0]'
      :rows='values?.[1]'
      :columns='(columns as never)'
      row-key='name'
      binary-state-sort
    >
      <template v-slot:body='props'>
        <q-tr :props='props'>
          <q-td key='ID' :props='props'>
            {{ props.row.ID }}
          </q-td>
          <q-td key='Index' :props='props'>
            {{ props.row.Index }}
          </q-td>
          <q-td key='SnapshotCommP' :props='props'>
            {{ props.row.SnapshotCommP }}
          </q-td>
          <q-td key='SnapshotRoot' :props='props'>
            {{ props.row.SnapshotRoot }}
          </q-td>
          <q-td key='SnapshotURI' :props='props'>{{ props.row.SnapshotURI }}</q-td>
          <q-td key='BackupState' :props='props'>{{ props.row.BackupState }}</q-td>
          <!-- <q-td key='ProposalCID' :props='props'>{{ props.row.ProposalCID }}</q-td> -->
          <!-- <q-td key='DealID' :props='props'>{{ props.row.DealID }}</q-td> -->
          <!-- <q-td key='Items' :props='props'>{{ props.row.Items?.join(',') }}</q-td> -->
          <q-td key='Op' :props='props' v-if='props.row?.BackupState === BackupState.BackupStateCreated'>
            <q-btn outline rounded color="primary" label="Backup" @click='onBackupClick(props.row)' :loading='props.row.Loading' />
          </q-td>

        </q-tr>
    </template>
  </q-table>
  </div>
</template>

<script setup lang='ts'>
import { useSnapshotStore } from 'src/teststore/snapshot';
import { Snapshot, BackupState } from 'src/teststore/snapshot/types';
import { computed, onMounted, ref } from 'vue';

const snapshot = useSnapshotStore()
const snapshots = computed(() => snapshot.Snapshots.Snapshots)

const snapshotsMap = computed(() => {
  const rowMap = new Map<string, Snapshot[]>() 
  Object.values(BackupState).forEach((state) => {
    const stateStr = state.toString()
    const rows = [] as Array<Snapshot>
    snapshots.value.forEach((sl) => {
      sl.Loading = false
      if (sl.BackupState === state) {
        rows?.push(sl)
      }
    })
    rowMap.set(stateStr, rows)
  } )
  return rowMap
})

const onBackupClick = (row: Snapshot) => {
  row.Loading = true
  snapshot.createBackup({
    Index: row.Index,
    Message:{}
  }, () => {
    row.Loading = false
  })
}

onMounted(() => {
  if(snapshots.value?.length === 0) {
    getSnapshots()
  }
})

const getSnapshots = () => {
  snapshot.getSnapshots({
    Message: {}
  }, () => {
    // TODO
  })
}

const columns = computed(() => [
  {
    name: 'ID',
    label: 'ID',
    field: (row: Snapshot) => row.ID,
    align: 'left',
  },
  {
    name: 'Index',
    label: 'Index',
    field: (row: Snapshot) => row.Index,
    align: 'left',
  },
  {
    name: 'SnapshotCommP',
    label: 'SnapshotCommP',
    field: (row: Snapshot) => row.SnapshotCommP,
    align: 'left',
  },
  {
    name: 'SnapshotRoot',
    label: 'SnapshotRoot',
    field: (row: Snapshot) => row.SnapshotRoot,
    align: 'left',
  },
  {
    name: 'SnapshotURI',
    label: 'SnapshotURI',
    field: (row: Snapshot) => row.SnapshotURI,
    align: 'left',
  },
  {
    name: 'BackupState',
    label: 'BackupState',
    field: (row: Snapshot) => row.BackupState,
    align: 'left',
  },
  // {
  //   name: 'ProposalCID',
  //   label: 'ProposalCID',
  //   field: (row: Snapshot) => row.ProposalCID,
  //   align: 'left',
  // },
  // {
  //   name: 'DealID',
  //   label: 'DealID',
  //   field: (row: Snapshot) => row.DealID,
  //   align: 'left',
  // },
  // {
  //   name: 'Items',
  //   label: 'Items',
  //   field: (row: Snapshot) => row.Items.join(','),
  //   align: 'left',
  // },
  {
    name: 'Op',
    label: '',
    align: 'left',
  },
]) 
</script>
