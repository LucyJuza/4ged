<template>
  <GenericList :list="elts">
  </GenericList>
</template>

<script setup>
import { computed } from 'vue';
import GenericList from './GenericList.vue';
import { useAppStore } from '@/stores/app';
const appStore = useAppStore()
const props = defineProps({
  list: [{
    "id": String,
    "gameId": String,
    "date": String,
    "location": String,
    "duration": Number,
    "participants": [String],
    "winners": [String]
  }],
  imagesEnabled: {
    type: Boolean,
    default: false
  }
})
const sorted = [...(props.list)].sort((a,b) => (new Date(b.date)) - (new Date(a.date)))
const elts = computed(() => sorted?.map(el => 
({
  title: `Partie du ${(new Date(el.date)).toLocaleDateString()}`, 
  subtitle: `Gagnant·e·s: ${el.winners.reduce((acc, curr) => acc += ", " + appStore.getPersonById(curr).name, "").substring(2)}`,
  image: props.imagesEnabled ? appStore.getGameById(el.gameId).image : '' ,
  link: `/plays/${el.id}/details`
})
)) 
</script>