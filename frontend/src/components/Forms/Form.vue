<template>
  <v-container>
    <v-form @submit.prevent="submitForm">
      <component 
        :is="currentFormComponent" 
        v-model="formData"
        @update:modelValue="updateFormData"
      />
      
      <v-row class="mt-4">
        <v-col cols="6">
          <v-btn
            block
            color="inversePrimary"
            variant="flat"
            @click="$emit('cancel')"
          >
          <v-icon icon="mdi-close"></v-icon>

            Annuler
          </v-btn>
        </v-col>
        <v-col cols="6">
          <v-btn
            block
            color="primary"
            type="submit"
            variant="flat"
          >
          <v-icon icon="mdi-plus"></v-icon>

            Ajouter
          </v-btn>
        </v-col>
      </v-row>
    </v-form>
  </v-container>
</template>

<script setup>
import { ref, computed, defineAsyncComponent } from 'vue'

const FormAddGame = defineAsyncComponent(() => import('./FormAddGame.vue'))
const FormAddPlay = defineAsyncComponent(() => import('./FormAddPlay.vue'))
const FormAddPlayer = defineAsyncComponent(() => import('./FormAddPlayer.vue'))

const props = defineProps({
  formType: {
    type: String,
    required: true,
    validator: (value) => ['game', 'play', 'player'].includes(value)
  }
})

const emit = defineEmits(['submit', 'cancel'])

const formData = ref({})

const currentFormComponent = computed(() => {
  const forms = {
    game: FormAddGame,
    play: FormAddPlay,
    player: FormAddPlayer
  }
  return forms[props.formType]
})

const updateFormData = (newData) => {
  formData.value = newData
}

const submitForm = () => {
  emit('submit', formData.value)
}
</script>