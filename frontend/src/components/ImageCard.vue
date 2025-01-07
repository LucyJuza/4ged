<template>
  <v-card
    variant="text"
    class="mx-auto flex-grow-1 flex-shrink-0  rounded-xl w-100 h-100"
  >
    <v-img
      class="align-end w-100 h-100"
      cover
      gradient="to bottom left, rgba(255,255,255,.33), rgba(0,0,0,.7)"
      :src=props.src
    >
        <v-card-title class="text-white">{{ props.text }}</v-card-title>
        <v-card-text class="pb-2"><span class="text-subtitle-2 font-weight-medium text-white">{{ props.subtext }}</span></v-card-text>
    </v-img>
  </v-card>
</template>

<script setup>
  const props = defineProps(['text', 'src', 'subtext'])
</script>
