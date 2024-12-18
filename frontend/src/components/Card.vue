<template>
  <v-card
    class="mx-auto"
    max-width="400"
  >
    <v-img
      class="align-end text-white"
      height="250"
      :src=props.src
      cover
    >
      <v-card-title>{{ props.text }}</v-card-title>
      <v-card-subtitle class="text-white pb-2">{{ props.subtext }}</v-card-subtitle>
    </v-img>
  </v-card>
</template>

<script setup>
  const props = defineProps(['text', 'src', 'subtext'])
</script>
