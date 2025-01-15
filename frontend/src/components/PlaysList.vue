<template>
  <GenericList :list="elts">
  </GenericList>
</template>

<script setup>
import { computed } from 'vue';
import GenericList from './GenericList.vue';
import { useAppStore } from '@/stores/app';
import { useFutureList } from '@/composables/futureList';
import { storeToRefs } from 'pinia';
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
let sorted = ref([])
sorted.value = props.list.sort((a,b) => (new Date(b.date)) - (new Date(a.date)) )
console.log(sorted)
let {data,error} = useFutureList(sorted.value?.map(async (play) => ({
  title: `Partie du ${(new Date(play.date)).toLocaleDateString()}`, 
  subtitle: `Gagnant·e·s: ${play.winners.reduce((acc, curr) => acc += ", " + appStore.getPersonById(curr).name, "").substring(2)}`,
  image: props.imagesEnabled ? (await appStore.getGameById(play.gameId)).image : '' ,
  link: `/plays/${play.id}/details`
})))
let elts = ref(data)
console.log(data)
</script>