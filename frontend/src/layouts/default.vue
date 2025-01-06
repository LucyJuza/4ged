<template>
  <v-app>
    <h1>{{ title }}</h1>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { useRouter } from 'vue-router';
const route = useRouter()
const title = ref(route.currentRoute.value.meta.title ?? "No title" )
</script>
