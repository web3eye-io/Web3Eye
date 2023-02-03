<template>
  <div view='lHh Lpr lFf'>
    <div class='row logo'>
      <q-space />
      <q-img :src='logo' style='width: 400px' fit='contain' />
      <q-space />
    </div>
    <q-input
      v-model="search"
      rounded
      outlined
      style='width: 650px;'
      placeholder='Coming soon'
    >
      <template v-slot:append>
        <q-icon name='camera_enhance' @click='photographer = true' />
      </template>
    </q-input>
    <div class='row looking'>
      <q-space />
      <q-btn label='Just looking' rounded flat>
        <q-tooltip>
          Coming soon
        </q-tooltip>
      </q-btn>
      <q-space />

    </div>
    <div class='occupier' />
  </div>
  <q-dialog v-model='photographer'>
      <q-card style="width: 700px; max-width: 80vw; height: 300px; max-height: 300px;">
        <q-uploader
          label="Custom header"
          single
          color='white'
          text-color='black'
          class='upload-box'
          field-name='upload'
          :form-fields='[{name: "topN", value: "10"}]'
          auto-upload
          @failed='onFailed'
          @uploaded='onUploaded'
          :factory='factoryFn'
        >
      <template v-slot:header="scope">
        <div class="row no-wrap items-center q-pa-sm q-gutter-xs">
          <q-btn v-if="scope.queuedFiles.length > 0" icon="clear_all" @click="scope.removeQueuedFiles" round dense flat >
            <q-tooltip>Clear All</q-tooltip>
          </q-btn>
          <q-btn v-if="scope.uploadedFiles.length > 0" icon="done_all" @click="scope.removeUploadedFiles" round dense flat >
            <q-tooltip>Remove Uploaded Files</q-tooltip>
          </q-btn>
          <q-spinner v-if="scope.isUploading" class="q-uploader__spinner" />
          <!-- <div class="col">
            <div class="q-uploader__title">Upload your files</div>
            <div class="q-uploader__subtitle">{{ scope.uploadSizeLabel }} / {{ scope.uploadProgressLabel }}</div>
          </div> -->
          <div>
            <span>Drag an image here or</span>
            <q-btn flat color="primary" @click="scope.pickFiles">
              upload a file
              <q-uploader-add-trigger />
            </q-btn>
          </div>
          <!-- <q-btn v-if="scope.canUpload" icon="cloud_upload" @click="scope.upload" round dense flat >
            <q-tooltip>Upload Files</q-tooltip>
          </q-btn> -->

          <q-btn v-if="scope.isUploading" icon="clear" @click="scope.abort" round dense flat >
            <q-tooltip>Abort Upload</q-tooltip>
          </q-btn>
        </div>
      </template>
    </q-uploader>
      </q-card>
    </q-dialog>
</template>

<script setup lang='ts'>
import { read } from 'fs';
import { useNFTMetaStore } from 'src/localstore/nft';
import { UploadResponse } from 'src/localstore/nft/types';
import { ref } from 'vue'
import { useRouter } from 'vue-router';

import logo from '../../assets/logo/logo.png'

const search = ref('')
const photographer = ref(false)

const router = useRouter()

const nft = useNFTMetaStore()

const onUploaded = (info: {
    /**
     * Uploaded files
     */
    files: readonly any[];
    /**
     * XMLHttpRequest that has been used to upload this batch of files
     */
    xhr: XMLHttpRequest;
  }) => {
  const reader = new FileReader()
  reader.readAsDataURL(info.files[0] as Blob)
  reader.onload = function() { 
    // console.log('result: ', this.result) // binary
    nft.NTFMetas.Current = window.URL.createObjectURL(info.files[0] as Blob)
	}
  const response = JSON.parse(info.xhr.response as string) as UploadResponse
  nft.setNftMeta(response.data)
  void router.push({
    path: '/result'
  })
}

const factoryFn = () => {
  return new Promise((resolve) => {
    resolve({
      url: '/api/nft-meta/search/file',
      method: 'POST'
    })
  })
}

const onFailed = () => {
  console.log('onFailed...')
}
</script>

<style scoped lang='sass'>
.logo
  margin: 10px 0 20px 0

.looking
  margin: 10px 0 10px 0
  color: $grey-8

.occupier
  height: 240px

.upload-box
  width: 100%
  height: 300px
  max-height: 300px
</style>
