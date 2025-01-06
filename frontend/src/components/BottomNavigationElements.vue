<template>
  <v-container fluid class="pa-0" v-for="item in navigationElements">
    <RouterLink :to="item.path" class="text-decoration-none">
      <v-col class="pa-2 text-onSurface">
        <v-row v-if="route.currentRoute.value.path == item.path"
          class="justify-center text-center ma-0 rounded-pill bg-secondaryContainer">
          <v-icon size="32">{{ item.iconSelected }}</v-icon>
        </v-row>
        <v-row v-else class="justify-center text-center ma-0">
          <v-icon size="32" color="onSurface">
            {{ item.icon }}
          </v-icon>
        </v-row>
        <v-row class="justify-center text-center ma-0">
          <span>{{ item.name }}</span>
        </v-row>
      </v-col>
    </RouterLink>
  </v-container>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { navigationElements } from './navigation';
const route = useRouter()
console.log(route.currentRoute.value.path)
</script>