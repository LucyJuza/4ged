<template>
  <ContentCard :title="title" :width-percent="(smAndDown ? '100%' : '49%')" :rounded-class="roundedClass">
    <slot>
    </slot>
  </ContentCard>
</template>

<script setup>
import { useDisplay } from 'vuetify';
import ContentCard from './ContentCard.vue';
const { smAndDown } = useDisplay()
defineProps({
  title: String,
  roundedClass: String,
}) 
</script>