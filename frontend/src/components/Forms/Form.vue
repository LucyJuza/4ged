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

<style>
@layer utilities {
  input[type="number"]::-webkit-inner-spin-button,
  input[type="number"]::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
  }

  input[type="number"] {
    appearance: textfield;
    -moz-appearance: textfield;
  }
}

.v-label.v-field-label--floating {
  background-color: rgb(var(--v-theme-secondaryContainer)) !important;
  border-radius: 4px !important;
  padding: 0 4px !important;
}

.v-input__prepend {
  display: none !important;
}

.v-field__outline {
  color: rgb(var(--v-theme-secondary)) !important;
}

.mdi-checkbox-marked {
  color: rgb(var(--v-theme-primary)) !important;
}
</style>