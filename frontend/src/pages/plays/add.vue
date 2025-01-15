<template>
  <FullContentCard>
    <div class="w-100 h-100 d-flex flex-column align-center justify-center overflow-auto">
      <Form
      @submit="() => {
      appStore.addPlay({
        gameId: formValues.selectedGame,
        date: (new Date(formValues.date)).toISOString(),
        duration: Number(formValues.duration),
        location: formValues.location,
        participants: formValues.players,
        winners: formValues.winners
      });
      appStore.setSelectedGameIdForPlayCreation(undefined)
      router.back()}"
      @cancel="() => {router.back()}"
      >
        <FormAddPlay :key="formValues.selectedGame"
        v-model="formValues">
        </FormAddPlay>
      </Form>
    </div>
  </FullContentCard>
</template>
<route lang="yaml">
meta:
  title: "Ajouter une nouvelle partie"
</route>
<script setup>
import Form from '@/components/Forms/Form.vue';
import FormAddPlay from '@/components/Forms/FormAddPlay.vue';
import FullContentCard from '@/components/FullContentCard.vue';
import { useFuture } from '@/composables/future';
import { useAppStore } from '@/stores/app';
import { computed, reactive, watch } from 'vue';
import { useRouter } from 'vue-router';
const appStore = useAppStore()
const router = useRouter()
let formValues = {
  selectedGame: undefined,
  date: "",
  location: "",
  duration: undefined,
  players: [],
  winners: [],
}
</script>
