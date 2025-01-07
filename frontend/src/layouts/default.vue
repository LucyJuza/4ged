<template>
  <v-app :full-height="false">
    <v-container fluid class="overflow-hidden d-flex flex-column align-stretch pa-0 ma-0 w-100 height-screen">
      <Header :title="title"/>
      <div class="flex-grow-1 overflow-hidden  ma-2 pa-0">
        <v-main class="h-100">
          <router-view />
        </v-main>
      </div>
      <v-bottom-navigation v-if="smAndDown" :elevation="0" grow class="bg-surfaceVariant" height="80">
        <BottomNavigationElements/>
      </v-bottom-navigation>
    </v-container>
  </v-app>
</template>

<script setup>
import BottomNavigationElements from '@/components/BottomNavigationElements.vue';
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useDisplay } from 'vuetify';
const { smAndDown, mdAndUp } = useDisplay()
const route = useRouter()
const title = computed( () => route.currentRoute.value.meta.title ?? "No title" )
</script>
<style>
#app{
  height: 100vh;
}
.v-application{
  height: 100dvh;
}
.height-screen{
  height: 100dvh;
}
</style>