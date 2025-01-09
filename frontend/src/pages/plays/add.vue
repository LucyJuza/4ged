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
        <FormAddPlay
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
import FormAddPlayer from '@/components/Forms/FormAddPlayer.vue';
import FullContentCard from '@/components/FullContentCard.vue';
import { useAppStore } from '@/stores/app';
import { useRouter } from 'vue-router';
const appStore = useAppStore()
const router = useRouter()
const formValues = { 
  name: undefined,
  image: undefined,
}
</script>
