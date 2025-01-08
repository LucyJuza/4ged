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
    "name": String,
    "image": String,
    "winrate": Number
  }],
  imagesEnabled: {
    type: Boolean,
    default: false
  }
})

const elts = computed(() => props.list?.map(el => 
({
  title: el.name, 
  subtitle: `winrate: ${el.winrate}%`,
  image: el.image,
  link: ''
})
)) 
</script>