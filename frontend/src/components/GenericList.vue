<template>
  <v-container fluid class="w-100 pa-0">
    <v-col
      class="w-100 justify-stretch px-0"
      v-for="(item, index) in list"
      :key="index"
    >
      <v-card class="w-100" rounded="lg" height="125" 
      @click="() => {route.push(item.link)}">
        <div class="w-100 h-100 d-flex flex-row flex-nowrap">
          <div class="d-flex flex-column flex-grow-1 h-100 ml-5 justify-center">
            <span class="text-h6">{{ item.title }}</span>
            <span class="text-subtitle">{{ item.subtitle }}</span>
          </div>
          <div v-if="item.image" class="w-25 flex-shrink-0">
            <v-img
            class="w-100 h-100"
            :src="item.image"
            cover
            />
          </div>
        </div>
      </v-card>
    </v-col>
  </v-container>
</template>

<script setup>
import { useRouter } from 'vue-router';

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