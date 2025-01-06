<template>
  <v-container fluid class="d-flex flex-row justify-start" v-for="item in navigationElements">
    <RouterLink :to="item.path" class="text-decoration-none w-100">
      <v-row class="pa-2 text-onSurface rounded-pill d-flex flex-row justify-start w-100 ga-3" 
      :class="{'bg-secondaryContainer': route.currentRoute.value.path == item.path}">
        <div>
          <v-icon v-if="route.currentRoute.value.path == item.path" size="32">
            {{ item.iconSelected }}
          </v-icon>
          <v-icon v-else size="32" color="onSurface" >
            {{ item.icon }}
          </v-icon>
        </div>
        <div class="d-flex flex-column justify-center">
          <span>{{ item.name }}</span>
        </div>
      </v-row>
    </RouterLink>
  </v-container>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { navigationElements } from './navigation';
import { computed } from 'vue';
const route = useRouter()
console.log(route.currentRoute.value.path)
</script>