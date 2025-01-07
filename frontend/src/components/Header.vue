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
</template>
<script setup>
import { useDisplay } from 'vuetify';
import NavigationDrawerElementsCopy from './NavigationDrawerElements copy.vue';

const { mdAndUp } = useDisplay()
defineProps({
  title:String
})
let drawer = ref(false)
</script>