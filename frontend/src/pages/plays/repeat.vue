<template>
  <FullContentCard>
    <div class="w-100 h-100 d-flex flex-column align-center justify-center overflow-auto">
      <Form
      @submit="() => {
      appStore.addPlay({
        gameId: appStore.getPlayById(formValues.selectedPlay).gameId,
        date: (new Date(formValues.date)).toISOString(),
        duration: Number(formValues.duration),
        location: formValues.location,
        participants: formValues.players,
        winners: formValues.winners
      });
      appStore.setSelectedPlayIdForRepetition(undefined)
      router.back()}"
      @cancel="() => {router.back()}"
      >
        <FormAddPlayRepeat
        v-model="formValues">
        </FormAddPlayRepeat>
      </Form>
    </div>
  </FullContentCard>
</template>
<route lang="yaml">
meta:
  title: "Répéter une partie"
</route>
<script setup>
import Form from '@/components/Forms/Form.vue';
import FormAddPlayRepeat from '@/components/Forms/FormAddPlayRepeat.vue';
import FullContentCard from '@/components/FullContentCard.vue';
import { useAppStore } from '@/stores/app';
import { useRouter } from 'vue-router';
const appStore = useAppStore()
const router = useRouter()
let formValues = {
  selectedPlay: undefined,
  date: "",
  location: "",
  duration: 0,
  players: [],
  winners: [],
}
</script>
