<template>
  <v-app>
    <Header :title="title"/>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup>
import { computed } from 'vue';
import { useRouter } from 'vue-router';
const route = useRouter()
const title = computed( () => route.currentRoute.value.meta.title ?? "No title" )
</script>
