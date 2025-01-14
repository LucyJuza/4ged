<template>
  <div class="h-100 pa-0 d-flex flex-column ga-4">
    <TwoCards firstCardTitle="Mes Jeux" secondCardTitle="Tous les jeux" 
    @scroll-first-card=""
    @scroll-second-card="scrollBhvrGames">
      <template #firstCardContent>
        <GameList :list="reactiveStore.filteredUserGames.value">

        </GameList>
      </template>
      <template #secondCardContent>
        <GameList :list="reactiveStore.filteredGames.value">

        </GameList>
      </template>
    </TwoCards>
    <v-container fluid class="bg-tertiaryContainer rounded-xl w-100 pa-2 d-flex flex-row justify-space-between flex-nowra ga-4">
          <v-text-field
            v-model="filter"
            label="Rechercher un jeu"
            prepend-icon="mdi-magnify"
            hide-details
            variant="outlined"
            @change="() => {appStore.filterGames(filter)}"
          ></v-text-field>
        <v-btn color="inversePrimary" 
        rounded="xl" size="60" variant="flat"
        @click="() => route.push('/games/add')">
          <v-icon size="32">mdi-plus</v-icon>
        </v-btn>
      </v-container>
  </div>
</template>
<route lang="yaml">
meta:
  title: "Mes Jeux"
</route>
<script setup>
import GameList from '@/components/GameList.vue';
import TwoCards from '@/components/TwoCards.vue';
import { useAppStore } from '@/stores/app';
import { storeToRefs } from 'pinia';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const route = useRouter()
const appStore = useAppStore()
const reactiveStore = storeToRefs(appStore)
const filter = ref('')
appStore.filterGames(filter.value)
let lastPullGamesDate = new Date()
const scrollBhvrGames = (e) =>{
  setTimeout(() => {
      if ( (new Date()) - lastPullGamesDate > 500
      && (Math.abs(e.target.scrollHeight - e.target.clientHeight - e.target.scrollTop) < 1)) {
      lastPullGamesDate = new Date()
      console.log("pulling next Games page")
      appStore.pushNextGamesPage()
    }
  },100)
}
</script>
