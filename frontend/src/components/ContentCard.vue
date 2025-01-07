<template>
  <v-card class="bg-tertiaryContainer d-flex flex-column h-100" :width="(smAndDown ? '100%' : '49%')" variant="tonal">
    <v-card-title primary-title class="text-center">
      {{ title }}
    </v-card-title>
    <div class="d-flex flex-column align-stretch justify-center pa-2 flex-grow-1 overflow-y-scroll">
      <div class="h-100 w-100">
        <slot>
        </slot>
      </div>
    </div>
  </v-card>
</template>
<script setup>
import { useDisplay } from 'vuetify';

const { smAndDown } = useDisplay()
defineProps({
  title: String,
})
</script>