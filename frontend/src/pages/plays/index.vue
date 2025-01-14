<template>
  <FullContentCard>
    <div class="d-flex flex-column">
      <PlaysList :images-enabled="true"
      style="max-width: 500px;"
      :list="appStore.userData.plays"
      />
    </div>
  </FullContentCard>
  <div style="position: fixed; right: 20px;" :style="'bottom:' + (smAndDown ? '100px' :'20px')">
    <v-btn
    rounded="lg"
    color="inversePrimary"
    size="56"
    elevation="8"
    @click="() => {router.push('/plays/add')}"
    >
      <v-icon size="32">
        mdi-plus
      </v-icon>
    </v-btn>
  </div>
</template>
<route lang="yaml">
meta:
  title: "Mes Parties"
</route>
<script setup>
import FullContentCard from '@/components/FullContentCard.vue';
import PlaysList from '@/components/PlaysList.vue';
import { useAppStore } from '@/stores/app';
import { useRouter } from 'vue-router';
import { useDisplay } from 'vuetify';
const { smAndDown } = useDisplay()
const appStore = useAppStore()
const router = useRouter()
//
</script>
