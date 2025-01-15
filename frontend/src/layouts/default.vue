<template>
  <v-app :full-height="false">
    <v-container v-if="isAuthenticated" fluid class="overflow-hidden d-flex flex-column align-stretch pa-0 ma-0 w-100 height-screen">
      <Header :title="title"/>
      <div class="flex-grow-1 overflow-y-hidden  ma-2 pa-0">
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
import Header from '@/components/Header.vue';
import BottomNavigationElements from '@/components/BottomNavigationElements.vue';
import { useAuthentication } from '@/composables/authentication';
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useDisplay } from 'vuetify';
const { isAuthenticated } = useAuthentication()
const { smAndDown, mdAndUp } = useDisplay()
const router = useRouter()
const title = computed( () => router.currentRoute.value.meta.title ?? "No title" )
if(!isAuthenticated.value){
  router.push('/login')
}
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