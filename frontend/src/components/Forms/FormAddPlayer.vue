<template>
  <v-text-field required v-model="form.name" label="Nom" placeholder="Nom de la personne" variant="outlined"
    density="comfortable" color="secondaryContainer" />
  <v-file-input v-model="form.image" accept="image/*" label="Image" prepend-icon="mdi-camera" show-size
    variant="outlined" required @update:model-value="handleImageChange" />
</template>

<script setup>
import { ref, watch } from "vue";

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  },
});

const emit = defineEmits(["update:modelValue"]);

const form = ref({
  name: "",
  image: null,
});

const previewImage = ref(null);
const imageBase64 = ref(null);

const handleImageChange = async (file) => {
  if (file) {
    try {
      imageBase64.value = await convertToBase64(file);
      if (previewImage.value) {
        URL.revokeObjectURL(previewImage.value);
      }
      previewImage.value = URL.createObjectURL(file);
    } catch (error) {
      showError('Erreur lors du traitement de l\'image');
    }
  } else {
    if (previewImage.value) {
      URL.revokeObjectURL(previewImage.value);
    }
    previewImage.value = null;
    imageBase64.value = null;
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


watch(
  form,
  (newValue) => {
    emit("update:modelValue", newValue);
  },
  { deep: true }
);

watch(
  () => props.modelValue,
  (newValue) => {
    if (Object.keys(newValue).length) {
      form.value = { ...newValue };
      if (newValue.image && newValue.image instanceof File) {
        handleImageChange(newValue.image);
      }
    }
  },
  { immediate: true }
);
</script>
