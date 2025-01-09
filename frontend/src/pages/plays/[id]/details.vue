<template>
  <FullContentCard>
    <v-container fluid class="pa-0 w-100 h-100 d-flex flex-column flex-nowrap align-center">
      <span class="text-h6 text-center">Informations</span>
      <div>
        <div>
          <span>Jeu: </span> <span>{{ appStore.getGameById(play.gameId).name }}</span>
        </div>
        <div>
          <span>Date de la partie: </span> <span>{{ (new Date(play.date)).toLocaleDateString() }}</span>
        </div>
        <div>
          <span>Lieu: </span> <span>{{ play.location }}</span>
        </div>
        <div>
          <span>Durée: </span> <span>{{ play.duration }} minutes</span>
        </div>
        <div>
          <span>Participant·e·s: </span> <span>{{ play.participants.reduce((acc,curr) => acc + ", " + appStore.getPersonById(curr).name,"").substring(2) }}</span>
        </div>
        <div>
          <span>Gagnant·e·s: </span> <span>{{ play.winners.reduce((acc,curr) => acc + ", " + appStore.getPersonById(curr).name,"").substring(2) }}</span>
        </div>
      </div>
      <v-btn
        color="primary"
        @click="() => {
          appStore.setSelectedPlayIdForRepetition(id);
          router.push('/plays/repeat')
        }"
      >
        <div class="d-flex flex-row justify-start align-center">
          <v-icon>mdi-refresh</v-icon>
          <span>Répéter la partie</span>
        </div>
      </v-btn>
    </v-container>
  </FullContentCard>
</template>
<route lang="yaml">
meta:
  title: "Détails"
</route>
<script setup>
import FullContentCard from '@/components/FullContentCard.vue';
import { useAppStore } from '@/stores/app';
import { useRoute, useRouter } from 'vue-router';
const appStore = useAppStore()
const route = useRoute();
const router = useRouter()
const id = route.params.id;
const play = appStore.getPlayById(id)
</script>