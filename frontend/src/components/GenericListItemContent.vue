<template>
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
</template>
<script setup>
defineProps({
  item:{
    title: String,
    subtitle: String,
    link: String,
    image: String
  }
})
</script>