<template>
  <FullContentCard>
    <div class="w-100 h-100 d-flex flex-column align-center justify-center">
      <Form
      @submit="() => {appStore.addPerson({
        name: formValues.name,
        image: formValues.image
      });router.back()}"
      @cancel="() => {router.back()}"
      >
        <FormAddPlayer
        v-model="formValues">
        </FormAddPlayer>
      </Form>
    </div>
  </FullContentCard>
</template>
<route lang="yaml">
meta:
  title: "Ajouter une personne"
</route>
<script setup>
import Form from '@/components/Forms/Form.vue';
import FormAddPlayer from '@/components/Forms/FormAddPlayer.vue';
import FullContentCard from '@/components/FullContentCard.vue';
import { useAppStore } from '@/stores/app';
import { useRouter } from 'vue-router';
const appStore = useAppStore()
const router = useRouter()
let formValues = { 
  name: undefined,
  image: undefined,
}
</script>
