<template>
  <v-container class="fill-height align-center">
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="6" lg="4">
        <v-form 
          ref="form"
          @submit.prevent="handleSubmit" 
          class="pa-0"
        >
          <div>
            <slot :form-data="formData" @update:form-data="updateFormData"></slot>
          </div>

          <div class="d-flex flex-row w-100 ga-2">
            <div class="flex-grow-1">
              <v-btn class="w-100 pa-0" color="inversePrimary" variant="flat" @click="$emit('cancel')">
                <v-icon icon="mdi-close"></v-icon>
                Annuler
              </v-btn>
            </div>
            <div class="flex-grow-1">
              <v-btn class="w-100 pa-0" color="primary" type="submit" variant="flat">
                <v-icon icon="mdi-plus"></v-icon>
                Ajouter
              </v-btn>
            </div>
          </div>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['submit', 'cancel'])
const form = ref(null)
const formData = ref({})

const updateFormData = (newData) => {
  formData.value = newData
}

const handleSubmit = async () => {
  const { valid } = await form.value.validate()
  
  if (valid) {
    emit('submit', formData.value)
  }
}
</script>

<style>
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