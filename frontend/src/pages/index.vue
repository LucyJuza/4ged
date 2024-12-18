<template>
  <BaseForm @submit="handleSubmit" @cancel="handleCancel">
    <template #default="{ formData, 'update:form-data': updateFormData }">
      <FormAddPlay
        :model-value="formData"
        @update:model-value="updateFormData"
      />
      <FormAddGame
        :model-value="formData"
        @update:model-value="updateFormData"
      />
      <FormAddPlayer
        :model-value="formData"
        @update:model-value="updateFormData"
      />
    </template>
  </BaseForm>
</template>

<script setup>
import BaseForm from '../components/Forms/Form.vue'
import FormAddPlay from '../components/Forms/FormAddPlay.vue'
import FormAddGame from '../components/Forms/FormAddGame.vue'
import FormAddPlayer from '../components/Forms/FormAddPlayer.vue'

const handleSubmit = (formData) => {
  console.log('Form submitted:', formData)
}

const handleCancel = () => {
  console.log('Form cancelled')
}
</script>
