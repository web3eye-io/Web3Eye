<template>
  <div class='nft-container'>
    <q-splitter v-model='splitterModel'>
      <template v-slot:before>
        <!-- left -->
        <div class='q-pa-md'>
          <div class='q-col-gutter-md row items-start'>
            <div class='col-12'>
              <q-img :src='currentImg'>
                <!-- <div class='absolute-bottom text-subtitle1 text-center'>
                  Distance: 0
                </div> -->
              </q-img>
            </div>
          </div>
          <div class='content-container column'>
            <!-- <h5>目标标题</h5> -->
            <!-- <div>
              <span class='label'>首发时间:</span>
              <span class='value'> XXXX-XX-XX</span>
            </div>
            <div>
              <span class='label'>公链信息:</span>
              <span class='value'> 由XX发布于XXXX公链</span>
            </div>
            <div>
              <span class='label'>网站信息:</span>
              <span class='value'> 由XX上传于XXXX网站</span>
            </div>
            <div>
              <span class='label'>作者:</span>
              <span class='value'> XXX</span>
            </div>
            <div>
              <span class='label'>描述:</span>
              <span class='value'> XXX</span>
            </div>
            <div>
              <span class='label'>版权声明:</span>
              <span class='value'> XXX</span>
            </div>
            <div>
              <span class='label'>稀缺度:</span>
              <span class='value'> 1%</span>
            </div> -->
          </div>
        </div>
      </template>
      <template v-slot:after>
        <!-- right -->
        <div class='q-pa-md'>
          <q-timeline color='secondary' v-for='nft in nfts' :key='nft.ID'>
            <q-timeline-entry subtitle='February 22, 1986'>
              <div class="row">
                <div class='col-md-2'>
                  <div v-if='getImageState(nft) === ImageState.Normal'>
                    <!-- for svg display -->
                    <q-icon
                      v-if='nft?.ImageURL?.startsWith("img")'
                      size='300px' 
                      :name='nft.ImageURL'
                    />
                    <q-img
                      v-else
                      :src="nft.ImageURL"
                      spinner-color="red"
                      @error='() => onLoadImageError(nft)'
                    />
                  </div>
                  <div v-if='getImageState(nft) === ImageState.IPFS'>
                    IPFS
                  </div>
                  <div v-if='getImageState(nft) === ImageState.Retrieving'>
                    Retrieving
                  </div>
                  <div v-if='getImageState(nft) === ImageState.WaitRecover'>
                    WaitRecover
                    <q-btn outline rounded color="primary" label="Recover" @click='startRetrieve(nft)' :loading='nft.Loading' />
                  </div>
                </div>
                <div class='column col-md-9' style="margin-left: 15px;">
                  <div>
                    <span class='label'>Similarity:</span>
                    <span class='value'>&nbsp; {{ nft.Distance }}</span>
                  </div>
                  
                  <div>
                    <span class='label'>ReleaseTime:</span>
                    <span class='value'></span>
                  </div>
                  <div>
                    <span class='label'>ChainID:</span>
                    <span class='value'>&nbsp; {{ nft.ChainID }}</span>
                  </div>
                  <div>
                    <span class='label'>ChainType:</span>
                    <span class='value'> {{ nft.ChainType }}</span>
                  </div>
                  <div>
                    <span class='label'>Contract:</span>
                    <span class='value'> {{ nft.Contract }}</span>
                  </div>
                  <div>
                    <span class='label'>TokenID:</span>
                    <span class='value'> {{ nft.TokenID }}</span>
                  </div>
                <div>
                    <span class='label'>Link:</span>
                    <span class='value'> {{nft.ImageURL}}</span>
                  </div>
                </div>
              </div>
            </q-timeline-entry>
            
          </q-timeline>
      </div>
      </template>
    </q-splitter>
    <!-- <div class="transaction">
      <h5>目标交易历史</h5>
    </div> -->
  </div>
</template>
<script lang='ts' setup>
import { useNFTMetaStore } from 'src/localstore/nft';
import { NFTMeta, ImageState } from 'src/localstore/nft/types';
import { useRetrieveStore } from 'src/teststore/retrieve';
import { computed, onMounted, ref, watch } from 'vue';

const splitterModel = ref(40)

const nft = useNFTMetaStore()
const nfts = computed(() => {
  const rows = [] as Array<NFTMeta>
  nft.NTFMetas.NTFMetas?.forEach((el) => {
    if(el.ImageURL?.startsWith('ipfs://')){
      el.ImageURL = el.ImageURL.replace('ipfs://', 'https://ipfs.io/ipfs/')
    }
    if (el.ImageURL?.startsWith('data:image')) {
      el.ImageURL = `img:${el.ImageURL}`
    }
    rows.push(el)
  })

  rows.sort((a, b) => a.Distance > b.Distance ? 1 : -1)
  return rows
})

const currentImg = computed(() => nft.NTFMetas.Current)

const retrieve = useRetrieveStore()
const retrieves = computed(() => retrieve.Retrieves.Retrieves)

const getImageState = computed(() => (row: NFTMeta) => {
  if (!row.LoadError) {
    return ImageState.Normal
  }
  if (row.ImageURL?.startsWith('img')) { // svg 
    return ImageState.Normal
  }

  if(checkImageExist.value(row.ImageURL)) {
    return ImageState.Normal
  }
  if (checkImageExist.value(row.IPFSImageURL)) {
    return ImageState.IPFS
  }

  if (row.ImageSnapshotID?.length > 0) {
    const _row = retrieves.value?.find((el) => el?.ChainType === row.ChainType && 
      el?.ChainID === row.ChainID && 
      el?.Contract === row.Contract && 
      el?.TokenID === row.TokenID
    )
    console.log('retrieves: ', retrieves.value)
    if (!_row) {
      console.log('not found')
      return ImageState.WaitRecover
    }
    if (_row.RetrieveState?.length > 0) {
      return ImageState.Retrieving
    }
  }
  return ImageState.WaitRecover
})

const onLoadImageError = (nft: NFTMeta) => {
  nft.LoadError = true
}

const checkImageExist = computed(() => (url: string) => {
  if (url === undefined) return false
  const image = new Image()
  image.src = url
  console.log('url: ', url)
  if (image.height > 0 && image.width > 0) {
    return true
  }
  return false
} )

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const checkImgExist = (imgUrl: string) => {
  return new Promise(function (resolve, reject) {
    const image = new Image();
    image.src = imgUrl;
    image.onload = (res) => {
      resolve(res);
    };
    image.onerror = (err) => {
      reject(err);
    };
  });
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const validateImage = (url: string) => {    
  var xmlHttp = {} as XMLHttpRequest;
  if (window.XMLHttpRequest){
    xmlHttp = new XMLHttpRequest();
  } 
  xmlHttp.open('Get', url, false);
  xmlHttp.send();
  if(xmlHttp.status === 404){
    return false;
  }
  return true;
}

const statRetrieve = (rows: NFTMeta[]) => {
  rows.forEach((row) => {
    retrieve.statRetrieve({
    ChainType: row.ChainType,
    ChainID: row.ChainID,
    Contract: row.Contract,
    TokenID: row.TokenID,
    Message: {}
  }, () => {
    // TODO
  })
})
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const startRetrieve = (row: NFTMeta) => {
  row.Loading = true
  retrieve.startRetrieve({
    ChainType: row.ChainType,
    ChainID: row.ChainID,
    Contract: row.Contract,
    TokenID: row.TokenID,
    Message: {}
  }, () => {
    // TODO
    row.Loading = false
  })
}

// const raws = ref([{
//     ID:'fc735773-fe95-4130-828e-88a0c6c08739',
//     ChainType:'Ethereum',
//     ChainID:'5',
//     Contract:'0x41cc069871054C1EfB4Aa40aF12f673eA2b6a1fC',
//     TokenID:'12000071',
//     URI:'https://token.staging.artblocks.io/0x41cc069871054c1efb4aa40af12f673ea2b6a1fc/12000071',
//     URIType:'http',
//     ImageURL:'https://media-proxy-staging.artblocks.io/0x41cc069871054c1efb4aa40af12f673ea2b6a1fc/1200007111.png',
//     VideoURL:'https://generator-staging-goerli.artblocks.io/0x41cc069871054c1efb4aa40af12f673ea2b6a1fc/12000071',
//     Description:'e.',
//     Name:'THIS ART IS ILLEGAL! #71',
//     VectorState: 40,
//     VectorID: '442184147106664837',
//     Distance: 1.0636201,
//     IPFSImageURL: 'url',
//     ImageSnapshotID: 'ImageSnapshotID',
//     Loading: false,
//   },
//   {
//     ID:'1d74f859-7860-453e-a3cd-72b33f8600c2',
//     ChainType:'Ethereum',
//     ChainID:'5',
//     Contract:'0x41cc069871054C1EfB4Aa40aF12f673eA2b6a1fC',
//     TokenID:'12000071',
//     URI:'https://token.staging.artblocks.io/0x41cc069871054c1efb4aa40af12f673ea2b6a1fc/12000071',
//     URIType:'http',
//     ImageURL:'https://media-proxy-staging.artblocks.io/0x41cc069871054c1efb4aa40af12f673ea2b6a1fc/1200007.png',
//     VideoURL:'https://generator-staging-goerli.artblocks.io/0x41cc069871054c1efb4aa40af12f673ea2b6a1fc/12000071',
//     Description:'e.',
//     Name:'THIS ART IS ILLEGAL! #71',
//     VectorState: 40,
//     VectorID: '442184147106664837',
//     Distance: 1.0636201,
//     IPFSImageURL: 'url',
//     ImageSnapshotID: 'ImageSnapshotID',
//     Loading: false,
//   }
// ] as Array<NFTMeta>)

watch(nfts, () => {
  if(nfts.value?.length > 0) {
    statRetrieve(nfts.value)
  }
})
onMounted(() => {
  if(nfts.value?.length > 0) {
    statRetrieve(nfts.value)
  }
})

</script>
<style lang='sass' scoped>
.nft-container
  ::v-deep .q-splitter--vertical > .q-splitter__separator
    width: 0px
  ::v-deep .q-timeline
    margin-top: 0
  .nft-result
    margin-left: 15px
  
  .content-container h5
    margin: 20px 0 6px 0
  .content-container div
    height: 30px
    line-height: 30px           
</style>