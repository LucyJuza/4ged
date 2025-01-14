<template>
  <FullContentCard>
    <div class="d-flex flex-column w-100">
        <PeopleList class="ma-auto" style="max-width: 500px;"
        :list="appStore.userData.persons"
        />
    </div>
  </FullContentCard>
  <div style="position: fixed; right: 20px;" :style="'bottom:' + (smAndDown ? '100px' :'20px')">
    <v-btn
    rounded="lg"
    color="inversePrimary"
    size="56"
    elevation="8"
    @click="() => {router.push('/people/add')}"
    >
      <v-icon size="32">
        mdi-plus
      </v-icon>
    </v-btn>
  </div>
</template>
<route lang="yaml">
meta:
  title: "Gens"
</route>
<script setup>
import FullContentCard from '@/components/FullContentCard.vue';
import PeopleList from '@/components/PeopleList.vue';
import { useAppStore } from '@/stores/app';
import { useRouter } from 'vue-router';
import { useDisplay } from 'vuetify';
const { smAndDown } = useDisplay()
const appStore = useAppStore()
const router = useRouter()
  //
</script>
