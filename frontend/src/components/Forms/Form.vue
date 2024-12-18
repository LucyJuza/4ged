<template>
  <v-container>
    <v-form @submit.prevent="submitForm">
      <slot
        :form-data="formData"
        @update:form-data="updateFormData"
      ></slot>

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
import { ref } from 'vue'

defineEmits(['submit', 'cancel'])

const formData = ref({})

const updateFormData = (newData) => {
  formData.value = newData
}

const submitForm = () => {
  emit('submit', formData.value)
}
</script>