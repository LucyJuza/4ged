<template>
  <v-container fluid class="w-100 pa-0">
    <v-col
      class="w-100 justify-stretch px-0"
      v-for="(item, index) in list"
      :key="index"
    >
      <v-card v-if="item.link" class="w-100" rounded="lg" height="125" 
      @click="() => {route.push(item.link)}">
        <GenericListItemContent :item="item"/>
      </v-card>
      <v-card v-else class="w-100" rounded="lg" height="125">
        <GenericListItemContent :item="item"/>
      </v-card>
    </v-col>
  </v-container>
</template>

<script setup>
import { useRouter } from 'vue-router';
import GenericListItemContent from './GenericListItemContent.vue';

const route = useRouter()
defineProps({
  list: [{
    title: String,
    subtitle: String,
    link: String,
    image: String
  }]
})
</script>