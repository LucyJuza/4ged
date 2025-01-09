<template>
  <div class="pa-0 ma-0 w-100 h-100 overflow-auto d-flex flex-column">
    <div style="height: 100px;" class="flex-shrink-0">
      <div class="d-flex flex-row align-center ga-3 flex-grow-0">
        <div :height="75" :width="75">
          <v-img rounded="circle" 
          :height="75"
          :width="75"
          :aspect-ratio="1/1"
          :src="appStore.userData.image"
          />
        </div>
        <h3>Bienvenue {{ appStore.userData.name }} !</h3>
      </div>
    </div>
    <div class="flex-grow-1 flex-shrink-0 overflow-auto" style="flex-basis: 0;">
      <div class="h-100 w-100">
        <v-container fluid class="w-100 h-100 pa-0 d-flex flex-column flex-md-wrap-reverse ga-4 overflow-auto">
          <HalfContentCard title="Statistiques" >
            <div class="h-100">
              <WinratePie/>
            </div>
          </HalfContentCard>
          <HalfContentCard title="Dernières parties">
            <div v-if="smAndDown" class="h-100 d-flex align-center">
              <HorizontalScroll>
                <div style="width:200px; height: 100px;" v-for="play in last4Plays">
                  <ImageCard
                  @click="() => router.push(`/plays/${play.id}/details`)"
                  :src="appStore.userData.games.find(g => g.id === play.gameId).image"
                  :text="appStore.userData.games.find(g => g.id === play.gameId).name"
                  :subtext="play.winners.includes(appStore.userData.personId) ? 'Gagnée' : 'Perdue'"/>
                </div>
              </HorizontalScroll>
            </div>
            <div v-else class="h-100 d-flex flex-row flex-wrap align-center justify-space-around ga-9">
              <div style="width:280px; height: 200px;" v-for="play in last4Plays">
                <ImageCard 
                  @click="() => router.push(`/plays/${play.id}/details`)"
                  :src="appStore.userData.games.find(g => g.id === play.gameId).image"
                  :text="appStore.userData.games.find(g => g.id === play.gameId).name"
                  :subtext="play.winners.includes(appStore.userData.personId) ? 'Gagnée' : 'Perdue'"/>
              </div>
            </div>
          </HalfContentCard>
        </v-container>
      </div>
    </div>
  </div>
  <div style="position: fixed; right: 20px;" :style="'bottom:' + (smAndDown ? '100px' :'20px')">
    <div class="d-flex flex-column align-end ga-2">
      <v-slide-x-transition>
        <div class="d-flex flex-column align-end ga-2" v-if="fabGroupOpenned">
          <v-btn
          rounded="lg"
          color="inversePrimary"
          height="56"
          elevation="8"
          @click="() => {router.push('/plays/add')}"
          >
            <v-icon size="32">
              mdi-gamepad-outline
            </v-icon>
            <span>
              Nouvelle partie
            </span>
          </v-btn>
          <v-btn
          rounded="lg"
          color="inversePrimary"
          height="56"
          elevation="8"
          @click="() => {router.push('/plays/repeat')}"
          >
            <v-icon size="32">
              mdi-repeat
            </v-icon>
            <span>
              Répéter partie
            </span>
          </v-btn>
        </div>
      </v-slide-x-transition>
      <v-fab-transition>
        <v-btn
        v-if="fabGroupOpenned"
        @click="() => fabGroupOpenned = false"
        rounded="lg"
        color="primary"
        size="56"
        elevation="200"
        >
          <v-icon size="32">
            mdi-close
          </v-icon>
        </v-btn>
        <v-btn
        v-else
        @click="() => fabGroupOpenned = true"
        rounded="lg"
        color="inversePrimary"
        size="56"
        elevation="8"
        >
          <v-icon size="32">
            mdi-plus
          </v-icon>
        </v-btn>
      </v-fab-transition>
    </div>
    
  </div>
</template>
<router lang="yaml">
meta:
  title: "Accueil"
</router>
<script setup>
import HalfContentCard from '@/components/HalfContentCard.vue';
import ImageCard from '@/components/ImageCard.vue';
import WinratePie from '@/components/WinratePie.vue';
import { useAppStore } from '@/stores/app';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useDisplay } from 'vuetify';
const { smAndDown } = useDisplay()
const router = useRouter()
const appStore = useAppStore()
const last4Plays = appStore.userData.plays.sort((a,b) => (new Date(b.date)) - (new Date(a.date)) ).slice(0,4)
const fabGroupOpenned = ref(false)
</script>
