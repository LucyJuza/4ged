<template>
  <FullContentCard>
    <div class="w-100 h-100 d-flex flex-column align-center justify-center">
      <Form
      @submit="() => {appStore.addGame({
        name: formValues.name,
        image: formValues.image,
        genres: formValues.genres.map(g => ({name: g}))
      });router.back()}"
      @cancel="() => {router.back()}"
      >
        <FormAddGame
        v-model="formValues">
        </FormAddGame>
      </Form>
    </div>
  </FullContentCard>
</template>
<route lang="yaml">
meta:
  title: "Ajouter un jeu"
</route>
<script setup>
import Form from '@/components/Forms/Form.vue';
import FormAddGame from '@/components/Forms/FormAddGame.vue';
import FullContentCard from '@/components/FullContentCard.vue';
import { useAppStore } from '@/stores/app';
import { useRouter } from 'vue-router';
const appStore = useAppStore()
const router = useRouter()
let formValues = { 
  name: undefined,
  image: undefined,
  genres: []
}
</script>