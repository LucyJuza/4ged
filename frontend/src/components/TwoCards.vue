<template>
    <v-container fluid class="h-100 w-100 pa-0 d-flex flex-column flex-wrap ga-4 overflow-auto">
      <HalfContentCard v-if="mdAndUp" :title="firstCardTitle">
        <slot name="firstCardContent"></slot>
      </HalfContentCard>
      <HalfContentCard v-if="mdAndUp" :title="secondCardTitle">
        <slot name="secondCardContent"></slot>
      </HalfContentCard>
      <div v-else class="h-100 w-100 d-flex flex-column ga-0">
        <div class="w-100 d-flex flex-row flex-nowrap justify-stretch text-center">
          <div class="flex-grow-1 rounded-t-lg text-h6 cursor-pointer" 
          :class="isFirstSelected ? 'bg-tertiaryContainer' : 'bg-tertiaryFixedDim'"
          @click="() => isFirstSelected=true">
            {{ firstCardTitle }}
          </div>
          <div class="flex-grow-1 rounded-t-lg text-h6 cursor-pointer" 
          :class="!isFirstSelected ? 'bg-tertiaryContainer' : 'bg-tertiaryFixedDim'"
          @click="() => isFirstSelected=false">
            {{ secondCardTitle }}
          </div>
        </div>
        <FullContentCard v-if="isFirstSelected" rounded-class="rounded-b-lg">
          <slot name="firstCardContent"></slot>
        </FullContentCard>
        <FullContentCard v-else rounded-class="rounded-b-lg">
          <slot name="secondCardContent" ></slot>
        </FullContentCard>
      </div>
    </v-container>
</template>

<script setup>
import { useDisplay } from 'vuetify';
import ContentCard from './ContentCard.vue';
import HalfContentCard from './HalfContentCard.vue';
import FullContentCard from './FullContentCard.vue';
import { ref } from 'vue';
const isFirstSelected = ref(true)
const { mdAndUp } = useDisplay()
defineProps({
  firstCardTitle: String,
  secondCardTitle: String,
})

</script>