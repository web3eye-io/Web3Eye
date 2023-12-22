<template>
  <div class='outer-bg'>
    <div class='outer-container'>
    <div>
      <div class='title row'>
        <q-item-label>{{ title }}</q-item-label>
        <q-space />
        <div class='column'>
          <q-space />
          <div class='row'>
            <div class='cursor-pointer lang'>
              {{ lang }}
              <q-menu>
                <q-list style="min-width: 100px">
                  <q-item clickable v-close-popup @click='lang = "EN"' dense>
                    <q-item-section>EN</q-item-section>
                  </q-item>
                  <q-item clickable v-close-popup @click='lang = "CN"' dense>
                    <q-item-section>CN</q-item-section>
                  </q-item>
                </q-list>
              </q-menu>
            </div>
            <!-- <q-btn icon='download' flat dense rounded /> -->
          </div>
        </div>
      </div>
      <div class='author row'>By
        <div class='author-name'>{{ author }}</div>
      </div>
      <div class='time-hint'>
        The article has {{ letters }} letters, read all will take about {{ timeNeed }} minutes.
      </div>
    </div>
    <div class='blog-body'>
      <slot name='EN' />
      <slot name='CN' />
    </div>
  </div>
  </div>
  
</template>

<script setup lang='ts'>
import { ref, toRef, watch } from 'vue'

interface Prop {
  title: string
  author: string
  downloadUrl: string
  letters: number
  timeNeed: number
}

const prop = defineProps<Prop>()
const title = toRef(prop, 'title')
const author = toRef(prop, 'author')

// TODO: use link for download url

const lang = ref('EN')

const emit = defineEmits<{(e: 'update:lang', v: string): void }>()
watch(lang, () => {
  emit('update:lang', lang.value)
})

</script>

<style scoped lang='sass'>
.title
  font-size: 64px
  color: $grey-8
  border-bottom: 4px solid $grey-4
  padding-bottom: 20px
  width: 100%

.content
  margin: 48px auto 0 auto

.lang
  font-size: 18px
  font-weight: bold
  color: blue
  width: 24px

.blog-body
  margin: 48px 0 48px 0

.author
  font-size: 18px
  color: $grey-8
  margin: 10px 0 10px 0

.author-name
  margin-left: 10px
  font-weight: bold
  color: blue

.time-hint
  color: $grey-8
</style>
