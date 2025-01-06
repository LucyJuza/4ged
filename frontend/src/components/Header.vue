<template>
  <v-container fluid class="d-flex flex-row">
    <v-icon v-if="mdAndUp" size="42" @click="drawer = !drawer">mdi-menu</v-icon>
    <v-spacer></v-spacer>
    <h1>{{ title }}</h1>
    <v-spacer></v-spacer>
  </v-container>
  <v-navigation-drawer v-if="mdAndUp"
    v-model="drawer"
    class="bg-surfaceVariant rounded"
    temporary
  >
    <NavigationDrawerElementsCopy/>
  </v-navigation-drawer>
  <v-bottom-navigation v-if="smAndDown" :elevation="0" grow class="bg-surfaceVariant" height="80">
    <BottomNavigationElements/>
  </v-bottom-navigation>
</template>
<script setup>
import { useDisplay } from 'vuetify';
import NavigationDrawerElementsCopy from './NavigationDrawerElements copy.vue';
import BottomNavigationElements from './BottomNavigationElements.vue';

 const { smAndDown, mdAndUp } = useDisplay()
defineProps({
  title:String
})
let drawer = ref(false)
</script>