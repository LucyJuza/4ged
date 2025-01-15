<template>
  <GenericList :list="elts">
  </GenericList>
</template>

<script setup>
import { computed } from 'vue';
import GenericList from './GenericList.vue';
const props = defineProps({
  list: [{
    "id": String,
    "name": String,
    "image": String,
    "genres": [
      {"name": String},
    ]
  }]
})
const elts = computed(() => props.list?.map(el => 
({
  title: el.name, 
  subtitle: el.genres ? el.genres.reduce( (acc,curr) => acc += ", " + curr.name,"" ).substring(2).substring(0,30) + "..." : "Pas de genres",
  image: el.image,
  link: `/games/${el.id}/details/`
})
)) 
</script>