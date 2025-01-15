<template>
  <FullContentCard>
    <v-container v-if="game" fluid class="pa-0 w-100 h-100 d-flex flex-column flex-nowrap align-center">
      <span class="text-h6 text-center">Informations</span>
      <div>
        <div>
          <span class="font-weight-bold">Nom:</span> <span>{{ game.name }}</span>
        </div>
        <div>
          <span class="font-weight-bold">Genres:</span> <span>{{ game.genres?.reduce( (acc,curr) => acc += ", " + curr.name,"" ).substring(2) }}</span>
        </div>
        <v-img
        class="mt-2"
        width="100"
        max-height="100"
        rounded="xl"
        :aspect-ratio="1/1"
        :src="game.image"
        lazy-src=""
        cover
        />
      </div>
      <span class="text-h6 mt-3">Historique</span>
      <div class="w-100 overflow-auto flex-shrink-0" max-height="300">
        <PlaysList style="max-width: 500px;" :list="history">
        </PlaysList>
      </div>
      <v-btn
        color="primary"
        @click="() => {
          appStore.setSelectedGameIdForPlayCreation(id);
          router.push('/plays/add')
        }"
      >
        <div class="d-flex flex-row justify-start align-center">
          <v-icon>mdi-plus</v-icon>
          <span>Nouvelle Partie</span>
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
import PlaysList from '@/components/PlaysList.vue';
import { useFuture } from '@/composables/future';
import { useAppStore } from '@/stores/app';
import { useRoute, useRouter } from 'vue-router';
const appStore = useAppStore()
const route = useRoute();
const router = useRouter()
const id = route.params.id;
const {data,err} = useFuture(appStore.getGameById(id))
const game = data
console.log("coucou")
console.log(game.value)
const history = appStore.getGameHistory(id)
</script>