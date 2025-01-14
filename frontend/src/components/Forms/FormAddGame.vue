<template>
  <v-text-field v-model="form.name" label="Nom" placeholder="Mon super jeu custom" variant="outlined"
    density="comfortable" required :rules="nameRules" />
  <v-file-input v-model="form.previewImage" accept="image/*" label="Image" prepend-icon="mdi-camera" show-size
    variant="outlined" required :rules="imageRules" @update:model-value="handleImageChange" />
  <v-autocomplete v-model="form.genres" :items="items" label="Genres" chips multiple placeholder="Ajouter des genres"
    variant="outlined" required :rules="genreRules">
    <template v-slot:chip="{ props, item }">
      <v-chip v-bind="props" class="bg-secondaryContainer">
        {{ item.raw }}
      </v-chip>
    </template>
  </v-autocomplete>
</template>

<script setup>
import { useAppStore } from '@/stores/app'
import { ref, watch } from 'vue'
const nameRules = [v => !!v || 'Le nom est requis']
const imageRules = [v => !!v || 'L\'image est requise']
const genreRules = [v => v?.length > 0 || 'Au moins un genre est requis']
const appStore = useAppStore()
const items = appStore.genres.map(g => g.name)
const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update:modelValue'])

const form = ref({
  name: '',
  image: null,
  previewImage: null,
  genres: []
})

const handleImageChange = async (file) => {
  if (file) {
    try {
      form.value.image = await convertToBase64(file);
      if (form.value.previewImage) {
        URL.revokeObjectURL(previewImage);
      }
      form.value.previewImage = URL.createObjectURL(file);
    } catch (error) {
      showError('Erreur lors du traitement de l\'image');
    }
  } else {
    if (form.value.previewImage) {
      URL.revokeObjectURL(form.value.previewImage);
    }
    form.value.previewImage = null;
    form.value.image = null;
  }
};

const convertToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
};

// Synchronise les changements avec le parent via v-model
watch(form, (newValue) => {
  emit('update:modelValue', newValue)
}, { deep: true })

// Initialise le formulaire avec les valeurs existantes si présentes
watch(() => props.modelValue, (newValue) => {
  if (Object.keys(newValue).length) {
    form.value = { ...newValue }
    if (newValue.image && newValue.image instanceof File) {
      handleImageChange(newValue.image)
    }
  }
}, { immediate: true })
</script>