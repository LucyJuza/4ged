<template>
  <Form 
  form-type="game" 
  @submit="handleSubmit" 
  @cancel="handleCancel"
  />

  <Form 
    form-type="play" 
    @submit="handleSubmit" 
    @cancel="handleCancel"
  />

  <Form 
    form-type="player" 
    @submit="handleSubmit" 
    @cancel="handleCancel"
  />
</template>

<script setup>
  //
</script>
