<template>
  <FullContentCard>
    <div class="w-100 d-flex flex-column align-center">
      <div>
        <div>
          <span class="font-weight-bold">Temps total de parties: </span>
          <span>~{{ Math.round(stats.playtime/60) }} heures</span>
        </div>
        <div>
          <span class="font-weight-bold">Plus de victoires: </span>
          <span>{{ appStore.getPersonById(stats.mostWins.person).name}} {{ stats.mostWins.won }}/{{ stats.mostWins.played }} ({{ Math.round(stats.mostWins.rate*100) }}%)</span>
        </div>
        <div>
          <span class="font-weight-bold">Plus haut taux de victoires: </span>
          <span>{{ appStore.getPersonById(stats.highestWinrate.person).name}} {{ Math.round(stats.highestWinrate.rate*100) }}%  ({{ stats.highestWinrate.won }}/{{ stats.highestWinrate.played }})</span>
        </div>
        <div>
          <span class="font-weight-bold">Jeux: </span>
          <span>{{ stats.games }}</span>
        </div>
        <div>
          <span class="font-weight-bold">Parties: </span>
          <span>{{ stats.plays }}</span>
        </div>
        <div>
          <span class="font-weight-bold">Joueur·euse·s: </span>
          <span>{{ stats.persons }}</span>
        </div>
        <v-divider class="border-opacity-100"></v-divider>
      </div>
      <div style="max-width: 300px;">
        <WinratePie/>
      </div>
    </div>
  </FullContentCard>
</template>
<route lang="yaml">
meta:
  title: "Statistiques"
</route>
<script setup>
import FullContentCard from '@/components/FullContentCard.vue';
import WinratePie from '@/components/WinratePie.vue';
import { useAppStore } from '@/stores/app';
const appStore = useAppStore()
const stats = appStore.getStatistics()
//
</script>
