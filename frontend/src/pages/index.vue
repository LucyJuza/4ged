<template>
  <v-container fluid class=" h-100 pa-0 d-flex flex-column flex-md-wrap ga-4">
    <ContentCard title="Statistiques">
      <div class="h-100">
        <WinratePie/>
      </div>
    </ContentCard>
    <ContentCard title="Dernières parties">
      <p>aled</p>
    </ContentCard>
  </v-container>
</template>
<route lang="yaml">
meta:
  title: "Accueil"
</route>
<script setup>
import ContentCard from '@/components/ContentCard.vue'
import WinratePie from '@/components/WinratePie.vue';
const handleSubmit = (formData) => {
  console.log('Form submitted:', formData)
}

const handleCancel = () => {
  console.log('Form cancelled')
}
</script>
