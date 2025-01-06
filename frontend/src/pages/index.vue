<template>
  <p>Contenu Accueil</p>
</template>
<route lang="yaml">
meta:
  title: "Accueil"
</route>
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
