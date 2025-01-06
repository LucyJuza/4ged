<template>
  <p>Contenu détails du jeu id: {{ id }}</p>
</template>
<route lang="yaml">
meta:
  title: "Détails"
</route>
<script setup>
import { useRoute } from 'vue-router';

const route = useRoute();
const id = route.params.id;
</script>